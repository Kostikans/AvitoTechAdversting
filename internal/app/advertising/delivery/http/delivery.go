package advertisingDelivery

import (
	"net/http"

	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
	"github.com/mailru/easyjson"

	"github.com/Kostikans/AvitoTechadvertising/internal/app/advertising"
	"github.com/gorilla/mux"
)

type AdvertisingDelivery struct {
	AdvertisingUsecase advertising.Usecase
}

func NewAdvertisingDelivery(r *mux.Router, AdvertisingUsecase advertising.Usecase) *AdvertisingDelivery {
	delivery := &AdvertisingDelivery{AdvertisingUsecase: AdvertisingUsecase}
	r.HandleFunc("api/v1/advertising", delivery.AddAdvertisingHandler).Methods("POST")
	r.HandleFunc("api/v1/advertising/{id:[0-9]+}", delivery.GetAdvertisingHandler).Methods("POST")
	r.HandleFunc("api/v1/advertising", delivery.AddAdvertisingHandler).Methods("GET")
	return delivery
}

func (AdvDelivery *AdvertisingDelivery) AddAdvertisingHandler(w http.ResponseWriter, r *http.Request) {
	var advertisingIn advertisingModel.Advertising
	err := easyjson.UnmarshalFromReader(r.Body, &advertisingIn)
	if err != nil {

	}

}

func (AdvDelivery *AdvertisingDelivery) GetAdvertisingHandler(w http.ResponseWriter, r *http.Request) {

}

func (AdvDelivery *AdvertisingDelivery) ListAdvertisingHandler(w http.ResponseWriter, r *http.Request) {
}
