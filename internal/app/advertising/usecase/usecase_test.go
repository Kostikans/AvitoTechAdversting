package advertisingUsecase

import (
	"errors"
	"testing"

	"github.com/Kostikans/AvitoTechadvertising/internal/package/customError"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/responseCodes"

	advertising_mock "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/mocks"
	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdvertisingUsecase_AddAdvertising(t *testing.T) {
	t.Run("AddAdvertisingPass", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		advertisingAdd := advertisingModel.AdvertisingAdd{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}
		retID := 1
		mockAdvertisingRepository := advertising_mock.NewMockRepository(ctrl)

		mockAdvertisingRepository.EXPECT().
			AddAdvertising(advertisingAdd).
			Return(retID, nil)

		bookingUsecase := NewAdvertisingUsecase(mockAdvertisingRepository, validator.New())

		advID, err := bookingUsecase.AddAdvertising(advertisingAdd)
		assert.NoError(t, err)
		assert.Equal(t, retID, advID)
	})
	t.Run("AddAdvertisingValidationErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		advertisingAdd := advertisingModel.AdvertisingAdd{Name: string(make([]byte, 300)), Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}
		mockAdvertisingRepository := advertising_mock.NewMockRepository(ctrl)

		bookingUsecase := NewAdvertisingUsecase(mockAdvertisingRepository, validator.New())

		_, err := bookingUsecase.AddAdvertising(advertisingAdd)
		assert.Error(t, err)
		assert.Equal(t, responseCodes.BadRequest, customError.ParseCode(err))
	})
	t.Run("AddAdvertisingValidationErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		advertisingAdd := advertisingModel.AdvertisingAdd{Name: "Продам гараж", Description: string(make([]byte, 1200)),
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}
		mockAdvertisingRepository := advertising_mock.NewMockRepository(ctrl)

		bookingUsecase := NewAdvertisingUsecase(mockAdvertisingRepository, validator.New())

		_, err := bookingUsecase.AddAdvertising(advertisingAdd)
		assert.Error(t, err)
		assert.Equal(t, responseCodes.BadRequest, customError.ParseCode(err))
	})
	t.Run("AddAdvertisingValidationErr3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		advertisingAdd := advertisingModel.AdvertisingAdd{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "sfd/test.jpeg", "sfd/test.jpeg", "sfd/test.jpeg"}, Cost: 1432}
		mockAdvertisingRepository := advertising_mock.NewMockRepository(ctrl)

		bookingUsecase := NewAdvertisingUsecase(mockAdvertisingRepository, validator.New())

		_, err := bookingUsecase.AddAdvertising(advertisingAdd)
		assert.Error(t, err)
		assert.Equal(t, responseCodes.BadRequest, customError.ParseCode(err))
	})

	t.Run("AddAdvertisingPassErr", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		advertisingAdd := advertisingModel.AdvertisingAdd{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}
		retID := 1
		mockAdvertisingRepository := advertising_mock.NewMockRepository(ctrl)

		mockAdvertisingRepository.EXPECT().
			AddAdvertising(advertisingAdd).
			Return(retID, customError.NewCustomError(errors.New("fsd"), responseCodes.ServerInternalError, 1))

		bookingUsecase := NewAdvertisingUsecase(mockAdvertisingRepository, validator.New())

		_, err := bookingUsecase.AddAdvertising(advertisingAdd)
		assert.Error(t, err)
		assert.Equal(t, responseCodes.ServerInternalError, customError.ParseCode(err))
	})
}

func TestAdvertisingUsecase_GetAdvertising(t *testing.T) {
	t.Run("GetAdvertisingPass", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAdvertisingRepository := advertising_mock.NewMockRepository(ctrl)

		advertisingID := 1
		fields := "photos"

		advertisingTest := advertisingModel.Advertising{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}

		mockAdvertisingRepository.EXPECT().
			GetAdvertising(advertisingID, fields).
			Return(advertisingTest, nil)

		mockAdvertisingRepository.EXPECT().
			CheckAdvertisingExist(advertisingID).
			Return(true, nil)

		bookingUsecase := NewAdvertisingUsecase(mockAdvertisingRepository, validator.New())

		advertising, err := bookingUsecase.GetAdvertising(advertisingID, fields)
		assert.NoError(t, err)
		assert.Equal(t, advertisingTest, advertising)
	})

	t.Run("GetAdvertisingErr1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAdvertisingRepository := advertising_mock.NewMockRepository(ctrl)

		advertisingID := 1
		fields := "photos"

		mockAdvertisingRepository.EXPECT().
			CheckAdvertisingExist(advertisingID).
			Return(false, nil)

		bookingUsecase := NewAdvertisingUsecase(mockAdvertisingRepository, validator.New())

		_, err := bookingUsecase.GetAdvertising(advertisingID, fields)
		assert.Error(t, err)
		assert.Equal(t, responseCodes.NotFound, customError.ParseCode(err))
	})

	t.Run("GetAdvertisingErr2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAdvertisingRepository := advertising_mock.NewMockRepository(ctrl)

		advertisingID := 1
		fields := "photos"

		advertisingTest := advertisingModel.Advertising{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}

		mockAdvertisingRepository.EXPECT().
			GetAdvertising(advertisingID, fields).
			Return(advertisingTest, customError.NewCustomError(errors.New("fs"), responseCodes.ServerInternalError, 1))

		mockAdvertisingRepository.EXPECT().
			CheckAdvertisingExist(advertisingID).
			Return(true, nil)

		bookingUsecase := NewAdvertisingUsecase(mockAdvertisingRepository, validator.New())

		_, err := bookingUsecase.GetAdvertising(advertisingID, fields)
		assert.Error(t, err)
		assert.Equal(t, responseCodes.ServerInternalError, customError.ParseCode(err))
	})
}
