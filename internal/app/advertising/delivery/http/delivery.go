package advertisingDelivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"

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
	r.HandleFunc("api/v1/advertising", delivery.AddAdvertisingHandler).Methods("POST")
	r.HandleFunc("api/v1/advertising/{id:[0-9]+}", delivery.GetAdvertisingHandler).Methods("GET")
	r.HandleFunc("api/v1/advertising", delivery.AddAdvertisingHandler).Methods("GET")
	return delivery
}

func (AdvDelivery *AdvertisingDelivery) AddAdvertisingHandler(w http.ResponseWriter, r *http.Request) {

	var advertisingIn advertisingModel.Advertising
	err := easyjson.UnmarshalFromReader(r.Body, &advertisingIn)
	if err != nil {
		customError.PostError(w, r, AdvDelivery.logger, err, nil)
	}
	advertisingID, err := AdvDelivery.AdvertisingUsecase.AddAdvertising(advertisingIn)
	if err != nil {
		customError.PostError(w, r, AdvDelivery.logger, err, nil)
	}
	responses.SendDataResponse(w, advertisingID)

}

func (AdvDelivery *AdvertisingDelivery) GetAdvertisingHandler(w http.ResponseWriter, r *http.Request) {

}

func (AdvDelivery *AdvertisingDelivery) ListAdvertisingHandler(w http.ResponseWriter, r *http.Request) {
}
