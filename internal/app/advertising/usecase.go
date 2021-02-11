//go:generate mockgen -source usecase.go -destination mocks/advertising_usecase_mock.go -package advertising_mock
package advertising

import advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"

type Usecase interface {
	AddAdvertising(advertising advertisingModel.AdvertisingAdd) (advertisingModel.AdvertisingID, error)
	GetAdvertising(advertisingID int, fields string) (advertisingModel.Advertising, error)
	ListAdvertising(field string, desc string, page int) (advertisingModel.AdvertisingList, error)
	CheckAdvertisingExist(AdvertisingID int) (bool, error)
}
