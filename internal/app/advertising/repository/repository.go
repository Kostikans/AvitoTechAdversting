package advertisingRepository

import (
	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
	"github.com/jmoiron/sqlx"
)

type AdvertisingRepository struct {
	db *sqlx.DB
}

func NewAdvertisingRepository(db *sqlx.DB) *AdvertisingRepository {
	return &AdvertisingRepository{db: db}
}
func (AdvRepo *AdvertisingRepository) AddAdvertising(advertising advertisingModel.Advertising) int {

}
func (AdvRepo *AdvertisingRepository) GetAdvertising(advertisingID int) advertisingModel.Advertising {

}
func (AdvRepo *AdvertisingRepository) ListAdvertising(field string, desc bool) advertisingModel.AdvertisingList {

}
