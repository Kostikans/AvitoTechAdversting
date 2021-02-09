package advertising

import advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"

type Repository interface {
	AddAdvertising(advertising advertisingModel.Advertising) (int, error)
	GetAdvertising(advertisingID int) (advertisingModel.Advertising, error)
	ListAdvertising(field string, desc bool) (advertisingModel.AdvertisingList, error)
}
