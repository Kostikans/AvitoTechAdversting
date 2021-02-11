//go:generate mockgen -source repository.go -destination mocks/advertising_repository_mock.go -package advertising_mock
package advertising

import advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"

type Repository interface {
	AddAdvertising(advertising advertisingModel.AdvertisingAdd) (advertisingModel.AdvertisingID, error)
	GetAdvertising(advertisingID int, fields string) (advertisingModel.Advertising, error)
	ListAdvertising(sort string, desc string, page int) (advertisingModel.AdvertisingList, error)
	CheckAdvertisingExist(advertisingID int) (bool, error)
	GenerateQueryForGetAdvertisingList(sort string, desc string) string
}
