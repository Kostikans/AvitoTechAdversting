package advertisingDelivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	"github.com/Kostikans/AvitoTechadvertising/internal/package/customError"

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

		req, err := http.NewRequest("POST", "/api/v1/advertising", bytes.NewBuffer(body))
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
		testAdvertisingAdd := "fdsfsd"
		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		body, err := json.Marshal(testAdvertisingAdd)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/api/v1/advertising", bytes.NewBuffer(body))
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

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		assert.Equal(t, responseCodes.BadRequest, response.Code)
	})
	t.Run("AddAdvertisingServerError", func(t *testing.T) {
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
			Return(advertisingIDTest, customError.NewCustomError(errors.New(""), responseCodes.ServerInternalError, 1))

		req, err := http.NewRequest("POST", "/api/v1/advertising", bytes.NewBuffer(body))
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

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		assert.Equal(t, responseCodes.ServerInternalError, response.Code)
	})
}

func TestAdvertisingDelivery_GetAdvertisingHandler(t *testing.T) {
	t.Run("GetAdvertising", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testAdvertisingGet := advertisingModel.Advertising{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}
		advertisingIDTest := advertisingModel.AdvertisingID{AdvID: 3}
		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		mockAdvertisingUseCase.EXPECT().
			GetAdvertising(advertisingIDTest.AdvID, "photos").
			Return(testAdvertisingGet, nil)

		req, err := http.NewRequest("GET", "/api/v1/advertising/3?fields=photos", nil)
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.Itoa(advertisingIDTest.AdvID),
		})
		rec := httptest.NewRecorder()
		handler := AdvertisingDelivery{
			AdvertisingUsecase: mockAdvertisingUseCase,
			logger:             logger.NewLogger(os.Stdout),
		}

		handler.GetAdvertisingHandler(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		var advertisingActual advertisingModel.Advertising
		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &advertisingActual)
		assert.Equal(t, testAdvertisingGet, advertisingActual)
		assert.Equal(t, responseCodes.OkCode, response.Code)
	})
	t.Run("GetAdvertisingNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testAdvertisingGet := advertisingModel.Advertising{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}
		advertisingIDTest := advertisingModel.AdvertisingID{AdvID: 3}
		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		body, err := easyjson.Marshal(testAdvertisingGet)
		assert.NoError(t, err)

		mockAdvertisingUseCase.EXPECT().
			GetAdvertising(advertisingIDTest.AdvID, "photos").
			Return(testAdvertisingGet, customError.NewCustomError(errors.New(""), responseCodes.NotFound, 1))

		req, err := http.NewRequest("GET", "/api/v1/advertising/3?fields=photos", bytes.NewBuffer(body))
		assert.NoError(t, err)

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.Itoa(advertisingIDTest.AdvID),
		})

		rec := httptest.NewRecorder()
		handler := AdvertisingDelivery{
			AdvertisingUsecase: mockAdvertisingUseCase,
			logger:             logger.NewLogger(os.Stdout),
		}

		handler.GetAdvertisingHandler(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		assert.Equal(t, responseCodes.NotFound, response.Code)
	})
}

func TestAdvertisingDelivery_ListAdvertisingHandler(t *testing.T) {
	t.Run("ListAdvertising", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		advertisingListTest := advertisingModel.AdvertisingList{List: []advertisingModel.AdvertisingListItem{{
			Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432,
		}, {Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432}},
			Page: advertisingModel.Page{CurrentPage: 1, PerPage: 10, LastPage: 2},
			Links: advertisingModel.Links{Last: "/api/v1/advertising?desc=true&page=2&sort=created",
				NextUrl: "/api/v1/advertising?desc=true&page=2&sort=created"}}

		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		mockAdvertisingUseCase.EXPECT().
			ListAdvertising("created", "true", 1).
			Return(advertisingListTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/advertising?sort=created&desc=true&page=1", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := AdvertisingDelivery{
			AdvertisingUsecase: mockAdvertisingUseCase,
			logger:             logger.NewLogger(os.Stdout),
		}

		handler.ListAdvertisingHandler(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		var advertisingActual advertisingModel.AdvertisingList
		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &advertisingActual)
		assert.Equal(t, advertisingListTest, advertisingActual)
		assert.Equal(t, responseCodes.OkCode, response.Code)
	})
	t.Run("ListAdvertising1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		advertisingListTest := advertisingModel.AdvertisingList{List: []advertisingModel.AdvertisingListItem{{
			Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432,
		}, {Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432}},
			Page: advertisingModel.Page{CurrentPage: 1, PerPage: 10, LastPage: 2},
			Links: advertisingModel.Links{Last: "/api/v1/advertising?desc=true&page=2&sort=created",
				NextUrl: "/api/v1/advertising?desc=true&page=2&sort=created"}}

		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		mockAdvertisingUseCase.EXPECT().
			ListAdvertising("created", "true", 1).
			Return(advertisingListTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/advertising?sort=created&desc=true", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := AdvertisingDelivery{
			AdvertisingUsecase: mockAdvertisingUseCase,
			logger:             logger.NewLogger(os.Stdout),
		}

		handler.ListAdvertisingHandler(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		var advertisingActual advertisingModel.AdvertisingList
		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &advertisingActual)
		assert.Equal(t, advertisingListTest, advertisingActual)
		assert.Equal(t, responseCodes.OkCode, response.Code)
	})
	t.Run("ListAdvertising2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		advertisingListTest := advertisingModel.AdvertisingList{List: []advertisingModel.AdvertisingListItem{{
			Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432,
		}, {Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432}},
			Page: advertisingModel.Page{CurrentPage: 3, PerPage: 10, LastPage: 4},
			Links: advertisingModel.Links{Last: "/api/v1/advertising?desc=true&page=4&sort=created",
				PrevUrl: "/api/v1/advertising?desc=true&page=2&sort=created",
				NextUrl: "/api/v1/advertising?desc=true&page=4&sort=created"}}

		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		mockAdvertisingUseCase.EXPECT().
			ListAdvertising("created", "true", 2).
			Return(advertisingListTest, nil)

		req, err := http.NewRequest("GET", "/api/v1/advertising?sort=created&desc=true&page=2", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := AdvertisingDelivery{
			AdvertisingUsecase: mockAdvertisingUseCase,
			logger:             logger.NewLogger(os.Stdout),
		}

		handler.ListAdvertisingHandler(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		var advertisingActual advertisingModel.AdvertisingList
		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		err = mapstructure.Decode(response.Data.(map[string]interface{}), &advertisingActual)
		assert.Equal(t, advertisingListTest, advertisingActual)
		assert.Equal(t, responseCodes.OkCode, response.Code)
	})
	t.Run("ListAdvertisingError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		advertisingListTest := advertisingModel.AdvertisingList{List: []advertisingModel.AdvertisingListItem{{
			Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432,
		}, {Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432}},
			Page: advertisingModel.Page{CurrentPage: 1, PerPage: 10, LastPage: 2}}

		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		mockAdvertisingUseCase.EXPECT().
			ListAdvertising("created", "true", 1).
			Return(advertisingListTest, customError.NewCustomError(errors.New(""), responseCodes.ServerInternalError, 1))

		req, err := http.NewRequest("GET", "/api/v1/advertising?sort=created&desc=true&page=1", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := AdvertisingDelivery{
			AdvertisingUsecase: mockAdvertisingUseCase,
			logger:             logger.NewLogger(os.Stdout),
		}

		handler.ListAdvertisingHandler(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		assert.Equal(t, responseCodes.ServerInternalError, response.Code)
	})
	t.Run("ListAdvertisingError2", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockAdvertisingUseCase := advertising_mock.NewMockUsecase(ctrl)

		req, err := http.NewRequest("GET", "/api/v1/advertising?sort=created&desc=true&page=-1", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		handler := AdvertisingDelivery{
			AdvertisingUsecase: mockAdvertisingUseCase,
			logger:             logger.NewLogger(os.Stdout),
		}

		handler.ListAdvertisingHandler(rec, req)
		resp := rec.Result()

		bodyTest, _ := ioutil.ReadAll(resp.Body)
		response := responses.HttpResponse{}

		err = json.Unmarshal(bodyTest, &response)
		assert.NoError(t, err)
		assert.Equal(t, responseCodes.BadRequest, response.Code)
	})
}
