package advertisingRepository

import (
	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/customError"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/responseCodes"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AdvertisingRepository struct {
	db *sqlx.DB
}

func NewAdvertisingRepository(db *sqlx.DB) *AdvertisingRepository {
	return &AdvertisingRepository{db: db}
}
func (advRepo *AdvertisingRepository) AddAdvertising(advertising advertisingModel.Advertising) (int, error) {
	var advertisingID int
	err := advRepo.db.QueryRow(AddAdvertising, advertising.Name, advertising.Description, pq.Array(advertising.Photos), advertising.Cost).Scan(&advertisingID)
	if err != nil {
		return advertisingID, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}
	return advertisingID, nil
}
func (advRepo *AdvertisingRepository) GetAdvertising(advertisingID int) (advertisingModel.Advertising, error) {
	var advertising advertisingModel.Advertising
	err := advRepo.db.QueryRow(GetAdvertising, advertisingID).Scan(&advertising.Name, advertising.Cost, advertising.Photos)
	if err != nil {
		return advertising, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}
	return advertising, nil
}
func (advRepo *AdvertisingRepository) ListAdvertising(field string, desc bool) (advertisingModel.AdvertisingList, error) {
	var advertisingList advertisingModel.AdvertisingList
	var advertisings []advertisingModel.Advertising
	err := advRepo.db.Select(&advertisings, listAdvertising)
	if err != nil {
		return advertisingList, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}
	advertisingList.List = advertisings
	return advertisingList, nil
}
