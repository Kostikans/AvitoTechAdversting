package advertisingUsecase

import (
	"github.com/Kostikans/AvitoTechadvertising/internal/app/advertising"
	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/customError"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/responseCodes"
	"github.com/go-playground/validator/v10"
)

type AdvertisingUsecase struct {
	AdvertisingRepo advertising.Repository
	validation      *validator.Validate
}

func NewAdvertisingUsecase(AdvertisingRepo advertising.Repository, validator *validator.Validate) *AdvertisingUsecase {
	return &AdvertisingUsecase{AdvertisingRepo: AdvertisingRepo, validation: validator}
}
func (AdvUsecase *AdvertisingUsecase) AddAdvertising(advertising advertisingModel.Advertising) (int, error) {
	err := AdvUsecase.validation.Struct(advertising)
	if err != nil {
		return 0, customError.NewCustomError(err, responseCodes.BadRequest, 1)
	}
	return AdvUsecase.AdvertisingRepo.AddAdvertising(advertising)
}
func (AdvUsecase *AdvertisingUsecase) GetAdvertising(advertisingID int, fields string) (advertisingModel.Advertising, error) {
	return AdvUsecase.AdvertisingRepo.GetAdvertising(advertisingID, fields)
}
func (AdvUsecase *AdvertisingUsecase) ListAdvertising(field string, desc string) (advertisingModel.AdvertisingList, error) {
	return AdvUsecase.AdvertisingRepo.ListAdvertising(field, desc)
}
