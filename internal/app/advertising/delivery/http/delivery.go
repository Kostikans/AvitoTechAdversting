package advertisingDelivery

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Kostikans/AvitoTechadvertising/internal/package/responseCodes"

	"github.com/Kostikans/AvitoTechadvertising/internal/package/abstractResponse"

	"github.com/Kostikans/AvitoTechadvertising/internal/package/customError"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/logger"

	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
	"github.com/mailru/easyjson"

	"github.com/Kostikans/AvitoTechadvertising/internal/app/advertising"
	"github.com/gorilla/mux"
)

type AdvertisingDelivery struct {
	AdvertisingUsecase advertising.Usecase
	logger             *logger.CustomLogger
}

func NewAdvertisingDelivery(r *mux.Router, AdvertisingUsecase advertising.Usecase, customLogger *logger.CustomLogger) *AdvertisingDelivery {
	delivery := &AdvertisingDelivery{AdvertisingUsecase: AdvertisingUsecase, logger: customLogger}
	r.HandleFunc("/api/v1/advertising", delivery.AddAdvertisingHandler).Methods("POST")
	r.HandleFunc("/api/v1/advertising/{id:[0-9]+}", delivery.GetAdvertisingHandler).Methods("GET")
	r.HandleFunc("/api/v1/advertising", delivery.ListAdvertisingHandler).Methods("GET")
	return delivery
}

// swagger:route POST /api/v1/advertising advertising AddAdvertising
// responses:
//  200:
//  201: AddAdvertising
//  400: badrequest
func (AdvDelivery *AdvertisingDelivery) AddAdvertisingHandler(w http.ResponseWriter, r *http.Request) {
	var advertisingIn advertisingModel.AdvertisingAdd
	err := easyjson.UnmarshalFromReader(r.Body, &advertisingIn)
	if err != nil {
		customError.PostError(w, r, AdvDelivery.logger, err, responseCodes.BadRequest)
		return
	}
	advertisingID, err := AdvDelivery.AdvertisingUsecase.AddAdvertising(advertisingIn)
	if err != nil {
		customError.PostError(w, r, AdvDelivery.logger, err, nil)
		return
	}
	abstractResponse.SendDataResponse(w, responseCodes.CreateCode, advertisingID)

}

// swagger:route GET /api/v1/advertising/{id} advertising GetAdvertising
// responses:
//  200: GetAdvertising
//  400: badrequest
//  404: notfound
func (AdvDelivery *AdvertisingDelivery) GetAdvertisingHandler(w http.ResponseWriter, r *http.Request) {
	advertisingIDVar := mux.Vars(r)["id"]
	advertisingID, _ := strconv.Atoi(advertisingIDVar)
	fields := r.FormValue("fields")

	advertising, err := AdvDelivery.AdvertisingUsecase.GetAdvertising(advertisingID, fields)
	if err != nil {
		customError.PostError(w, r, AdvDelivery.logger, err, nil)
		return
	}
	abstractResponse.SendDataResponse(w, responseCodes.OkCode, advertising)
}

// swagger:route GET /api/v1/advertising advertising ListAdvertising
// responses:
//  200: ListAdvertising
//  400: badrequest
func (AdvDelivery *AdvertisingDelivery) ListAdvertisingHandler(w http.ResponseWriter, r *http.Request) {
	sort := r.FormValue("sort")
	desc := r.FormValue("desc")

	pageVar := r.FormValue("page")
	page, err := strconv.Atoi(pageVar)
	if err != nil {
		page = 1
	}
	if page < 1 {
		customError.PostError(w, r, AdvDelivery.logger, errors.New("page must be positive number"), responseCodes.BadRequest)
		return
	}

	advertisingList, err := AdvDelivery.AdvertisingUsecase.ListAdvertising(sort, desc, page)
	if err != nil {
		customError.PostError(w, r, AdvDelivery.logger, err, nil)
		return
	}

	advertisingList.Links = AdvDelivery.makeLinksHttpForListAdvertising(sort, desc, advertisingList.Page)
	abstractResponse.SendDataResponse(w, responseCodes.OkCode, advertisingList)
}

func (AdvDelivery *AdvertisingDelivery) makeLinksHttpForListAdvertising(sort string, desc string, page advertisingModel.Page) advertisingModel.Links {
	url := url.URL{
		Path: "/api/v1/advertising",
	}
	links := advertisingModel.Links{}
	query := url.Query()

	query.Set("page", strconv.Itoa(page.LastPage))
	query.Set("sort", sort)
	query.Set("desc", desc)
	url.RawQuery = query.Encode()
	links.Last = url.Path + "?" + url.RawQuery

	if page.CurrentPage < page.LastPage {
		query.Set("page", strconv.Itoa(page.CurrentPage+1))
		url.RawQuery = query.Encode()
		links.NextUrl = url.Path + "?" + url.RawQuery
	}

	if page.CurrentPage > 1 {
		query.Set("page", strconv.Itoa(page.CurrentPage-1))
		url.RawQuery = query.Encode()
		links.PrevUrl = url.Path + "?" + url.RawQuery
	}

	return links
}
