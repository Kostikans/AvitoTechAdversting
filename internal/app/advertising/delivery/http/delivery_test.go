package advertisingDelivery

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mitchellh/mapstructure"

	"github.com/Kostikans/AvitoTechadvertising/internal/package/responseCodes"

	advertising_mock "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/mocks"

	"github.com/mailru/easyjson"

	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"

	"github.com/Kostikans/AvitoTechadvertising/internal/package/logger"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/responses"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdvertisingDelivery_AddAdvertisingHandler(t *testing.T) {
	t.Run("AddAdvertising", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testAdvertisingAdd := advertisingModel.AdvertisingAdd{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}
		advertisingIDTest := advertisingModel.AdvertisingID{AdvID: 3}
		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		body, err := easyjson.Marshal(testAdvertisingAdd)
		assert.NoError(t, err)

		mockAdvertisingUseCase.EXPECT().
			AddAdvertising(testAdvertisingAdd).
			Return(advertisingIDTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/advertising", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := AdvertisingDelivery{
			AdvertisingUsecase: mockAdvertisingUseCase,
			logger:             logger.NewLogger(os.Stdout),
		}

		handler.AddAdvertisingHandler(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		var advertisingIDActual advertisingModel.AdvertisingID
		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &advertisingIDActual)
		assert.Equal(t, advertisingIDTest, advertisingIDActual)
		assert.Equal(t, responseCodes.CreateCode, response.Code)
	})
	t.Run("AddAdvertisingBadRequest", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testAdvertisingAdd := advertisingModel.AdvertisingAdd{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}
		advertisingIDTest := advertisingModel.AdvertisingID{AdvID: 3}
		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		body, err := easyjson.Marshal(testAdvertisingAdd)
		assert.NoError(t, err)

		mockAdvertisingUseCase.EXPECT().
			AddAdvertising(testAdvertisingAdd).
			Return(advertisingIDTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/advertising", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := AdvertisingDelivery{
			AdvertisingUsecase: mockAdvertisingUseCase,
			logger:             logger.NewLogger(os.Stdout),
		}

		handler.AddAdvertisingHandler(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		var advertisingIDActual advertisingModel.AdvertisingID
		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &advertisingIDActual)
		assert.Equal(t, advertisingIDTest, advertisingIDActual)
		assert.Equal(t, responseCodes.CreateCode, response.Code)
	})
}
