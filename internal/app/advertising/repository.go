//go:generate mockgen -source repository.go -destination mocks/advertising_repository_mock.go -package advertising_mock
package advertising

import advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"

type Repository interface {
	AddAdvertising(advertising advertisingModel.AdvertisingAdd) (int, error)
	GetAdvertising(advertisingID int, fields string) (advertisingModel.Advertising, error)
	ListAdvertising(field string, desc string) (advertisingModel.AdvertisingList, error)
	CheckAdvertisingExist(advertisingID int) (bool, error)
}
