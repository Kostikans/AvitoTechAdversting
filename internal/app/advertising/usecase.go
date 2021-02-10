package advertising

import advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"

type Usecase interface {
	AddAdvertising(advertising advertisingModel.Advertising) (int, error)
	GetAdvertising(advertisingID int, fields string) (advertisingModel.Advertising, error)
	ListAdvertising(field string, desc string) (advertisingModel.AdvertisingList, error)
}
