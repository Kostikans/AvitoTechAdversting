package advertisingRepository

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lib/pq"

	"github.com/Kostikans/AvitoTechadvertising/internal/package/customError"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/responseCodes"

	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAdvertisingRepository_AddAdvertising(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("AddAdvertising", func(t *testing.T) {
		advertisingID := advertisingModel.AdvertisingID{AdvID: 3}
		rowsAdvertising := sqlmock.NewRows([]string{"advertising_id"}).AddRow(
			3)

		query := AddAdvertising

		advertisingAddTest := advertisingModel.AdvertisingAdd{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}

		mock.ExpectQuery(query).
			WithArgs(advertisingAddTest.Name, advertisingAddTest.Description, advertisingAddTest.Photos[0], pq.Array(advertisingAddTest.Photos[1:]), advertisingAddTest.Cost).
			WillReturnRows(rowsAdvertising)

		query = IncrementAdvertisingCount
		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows()

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		advertisingIDActual, err := rep.AddAdvertising(advertisingAddTest)
		assert.NoError(t, err)
		assert.Equal(t, advertisingID, advertisingIDActual)
	})

	t.Run("AddAdvertisingErr", func(t *testing.T) {

		query := AddAdvertising

		advertisingAddTest := advertisingModel.AdvertisingAdd{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}

		mock.ExpectQuery(query).
			WithArgs(advertisingAddTest).
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		_, err := rep.AddAdvertising(advertisingAddTest)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), responseCodes.ServerInternalError)
	})

	t.Run("AddAdvertisingErr2", func(t *testing.T) {
		query := AddAdvertising

		rowsAdvertising := sqlmock.NewRows([]string{"advertising_id"}).AddRow(
			3)

		advertisingAddTest := advertisingModel.AdvertisingAdd{Name: "Продам гараж", Description: "Очень хороший",
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Cost: 1432}

		mock.ExpectQuery(query).
			WithArgs(advertisingAddTest.Name, advertisingAddTest.Description, advertisingAddTest.Photos[0], pq.Array(advertisingAddTest.Photos[1:]), advertisingAddTest.Cost).
			WillReturnRows(rowsAdvertising)

		query = IncrementAdvertisingCount
		mock.ExpectQuery(query).
			WithArgs().
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		_, err := rep.AddAdvertising(advertisingAddTest)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), responseCodes.ServerInternalError)
	})
}

func TestAdvertisingRepository_GetAdvertising(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("GetAdvertising", func(t *testing.T) {
		advertisingID := advertisingModel.AdvertisingID{AdvID: 3}
		rowsAdvertising := sqlmock.NewRows([]string{"name", "cost", "photos[1]"}).AddRow(
			"Продам гараж", 1432, "sfd/test.jpeg")

		query := GetAdvertising

		advertisingAddTest := advertisingModel.Advertising{Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432}

		mock.ExpectQuery(query).
			WithArgs(advertisingID.AdvID).
			WillReturnRows(rowsAdvertising)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		advertisingActual, err := rep.GetAdvertising(advertisingID.AdvID, "")
		assert.NoError(t, err)
		assert.Equal(t, advertisingAddTest, advertisingActual)
	})
	t.Run("GetAdvertising1", func(t *testing.T) {
		advertisingID := advertisingModel.AdvertisingID{AdvID: 3}
		rowsAdvertising := sqlmock.NewRows([]string{"name", "cost", "photos[1]", "photos"}).AddRow(
			"Продам гараж", 1432, "sfd/test.jpeg", pq.Array([]string{"sfd/test.jpeg", "svcx/gd.jpeg"}))

		query := GetAdvertisingWithPhotos

		advertisingAddTest := advertisingModel.Advertising{Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432,
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}}

		mock.ExpectQuery(query).
			WithArgs(advertisingID.AdvID).
			WillReturnRows(rowsAdvertising)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		advertisingActual, err := rep.GetAdvertising(advertisingID.AdvID, "photos")
		assert.NoError(t, err)
		assert.Equal(t, advertisingAddTest, advertisingActual)
	})
	t.Run("GetAdvertising2", func(t *testing.T) {
		advertisingID := advertisingModel.AdvertisingID{AdvID: 3}
		rowsAdvertising := sqlmock.NewRows([]string{"name", "cost", "photos[1]", "description"}).AddRow(
			"Продам гараж", 1432, "sfd/test.jpeg", "Очень хороший")

		query := GetAdvertisingWithDescription

		advertisingAddTest := advertisingModel.Advertising{Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432, Description: "Очень хороший"}

		mock.ExpectQuery(query).
			WithArgs(advertisingID.AdvID).
			WillReturnRows(rowsAdvertising)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		advertisingActual, err := rep.GetAdvertising(advertisingID.AdvID, "description")
		assert.NoError(t, err)
		assert.Equal(t, advertisingAddTest, advertisingActual)
	})
	t.Run("GetAdvertising3", func(t *testing.T) {
		advertisingID := advertisingModel.AdvertisingID{AdvID: 3}
		rowsAdvertising := sqlmock.NewRows([]string{"name", "cost", "photos[1]", "photos", "description"}).AddRow(
			"Продам гараж", 1432, "sfd/test.jpeg", pq.Array([]string{"sfd/test.jpeg", "svcx/gd.jpeg"}), "Очень хороший")

		query := GetAdvertisingWithPhotosAndDescription

		advertisingAddTest := advertisingModel.Advertising{Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432,
			Photos: []string{"sfd/test.jpeg", "svcx/gd.jpeg"}, Description: "Очень хороший"}

		mock.ExpectQuery(query).
			WithArgs(advertisingID.AdvID).
			WillReturnRows(rowsAdvertising)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		advertisingActual, err := rep.GetAdvertising(advertisingID.AdvID, "description,photos")
		assert.NoError(t, err)
		assert.Equal(t, advertisingAddTest, advertisingActual)
	})
	t.Run("GetAdvertisingErr", func(t *testing.T) {
		advertisingID := advertisingModel.AdvertisingID{AdvID: 3}

		query := GetAdvertisingWithPhotosAndDescription

		mock.ExpectQuery(query).
			WithArgs(advertisingID.AdvID).
			WillReturnError(errors.New(""))

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		_, err := rep.GetAdvertising(advertisingID.AdvID, "description,photos")
		assert.Error(t, err)
		assert.Equal(t, responseCodes.ServerInternalError, customError.ParseCode(err))
	})
}

func TestAdvertisingRepository_ListAdvertising(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("ListAdvertising", func(t *testing.T) {
		rowsCount := sqlmock.NewRows([]string{"count"}).AddRow(
			11)

		rowsAdvertising := sqlmock.NewRows([]string{"name", "mainphoto", "cost"}).AddRow(
			"Продам гараж", "sfd/test.jpeg", 1432).AddRow(
			"Продам гараж", "sfd/test.jpeg", 1432)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)
		query := rep.GenerateQueryForGetAdvertisingList("cost", "true")

		advertisingListTest := advertisingModel.AdvertisingList{List: []advertisingModel.AdvertisingListItem{{
			Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432,
		}, {Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432}},
			Page: advertisingModel.Page{CurrentPage: 1, PerPage: 10, LastPage: 2}}

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsAdvertising)

		query = GetPageCount
		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsCount)

		advertisingListActual, err := rep.ListAdvertising("cost", "true", 1)
		assert.NoError(t, err)
		assert.Equal(t, advertisingListTest, advertisingListActual)
	})

	t.Run("ListAdvertising", func(t *testing.T) {
		rowsCount := sqlmock.NewRows([]string{"count"}).AddRow(
			11)

		rowsAdvertising := sqlmock.NewRows([]string{"name", "mainphoto", "cost"}).AddRow(
			"Продам гараж", "sfd/test.jpeg", 1432).AddRow(
			"Продам гараж", "sfd/test.jpeg", 1432)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)
		query := rep.GenerateQueryForGetAdvertisingList("created", "false")

		advertisingListTest := advertisingModel.AdvertisingList{List: []advertisingModel.AdvertisingListItem{{
			Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432,
		}, {Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432}},
			Page: advertisingModel.Page{CurrentPage: 1, PerPage: 10, LastPage: 2}}

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsAdvertising)

		query = GetPageCount
		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsCount)

		advertisingListActual, err := rep.ListAdvertising("created", "false", 1)
		assert.NoError(t, err)
		assert.Equal(t, advertisingListTest, advertisingListActual)
	})

	t.Run("ListAdvertising", func(t *testing.T) {
		rowsCount := sqlmock.NewRows([]string{"count"}).AddRow(
			11)

		rowsAdvertising := sqlmock.NewRows([]string{"name", "mainphoto", "cost"}).AddRow(
			"Продам гараж", "sfd/test.jpeg", 1432).AddRow(
			"Продам гараж", "sfd/test.jpeg", 1432)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)
		query := rep.GenerateQueryForGetAdvertisingList("created", "true")

		advertisingListTest := advertisingModel.AdvertisingList{List: []advertisingModel.AdvertisingListItem{{
			Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432,
		}, {Name: "Продам гараж", MainPhoto: "sfd/test.jpeg", Cost: 1432}},
			Page: advertisingModel.Page{CurrentPage: 1, PerPage: 10, LastPage: 2}}

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsAdvertising)

		query = GetPageCount
		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsCount)

		advertisingListActual, err := rep.ListAdvertising("created", "true", 1)
		assert.NoError(t, err)
		assert.Equal(t, advertisingListTest, advertisingListActual)
	})

	t.Run("ListAdvertisingErr", func(t *testing.T) {

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)
		query := rep.GenerateQueryForGetAdvertisingList("cost", "true")

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnError(errors.New(""))

		_, err := rep.ListAdvertising("cost", "true", 1)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), responseCodes.ServerInternalError)
	})

	t.Run("ListAdvertisingErr2", func(t *testing.T) {

		sqlxDb := sqlx.NewDb(db, "sqlmock")
		rowsAdvertising := sqlmock.NewRows([]string{"name", "mainphoto", "cost"}).AddRow(
			"Продам гараж", "sfd/test.jpeg", 1432).AddRow(
			"Продам гараж", "sfd/test.jpeg", 1432)

		rep := NewAdvertisingRepository(sqlxDb)
		query := rep.GenerateQueryForGetAdvertisingList("cost", "true")

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsAdvertising)

		query = GetPageCount
		mock.ExpectQuery(query).
			WithArgs().
			WillReturnError(errors.New(""))

		_, err := rep.ListAdvertising("cost", "true", 1)
		assert.Error(t, err)
		assert.Equal(t, customError.ParseCode(err), responseCodes.ServerInternalError)
	})
}

func TestAdvertisingRepository_CheckAdvertisingExist(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("CheckAdvertisingExist", func(t *testing.T) {
		rowsAdvertising := sqlmock.NewRows([]string{"exists"}).AddRow(
			true)

		query := fmt.Sprintf("SELECT exists (%s)", CheckAdvertisingExist)

		mock.ExpectQuery(query).
			WithArgs(3).
			WillReturnRows(rowsAdvertising)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		exist, err := rep.CheckAdvertisingExist(3)
		assert.NoError(t, err)
		assert.Equal(t, true, exist)
	})

	t.Run("CheckAdvertisingExistFalse", func(t *testing.T) {
		rowsAdvertising := sqlmock.NewRows([]string{"exists"}).AddRow(
			false)

		query := fmt.Sprintf("SELECT exists (%s)", CheckAdvertisingExist)

		mock.ExpectQuery(query).
			WithArgs(3).
			WillReturnRows(rowsAdvertising)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		exist, err := rep.CheckAdvertisingExist(3)
		assert.NoError(t, err)
		assert.Equal(t, false, exist)
	})

}
