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
}

func TestAdvertisingRepository_GetAdvertising(t *testing.T) {

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

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)

		advertisingIDActual, err := rep.AddAdvertising(advertisingAddTest)
		assert.NoError(t, err)
		assert.Equal(t, advertisingID, advertisingIDActual)
	})
}

func TestAdvertisingRepository_ListAdvertising(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("ListAdvertising", func(t *testing.T) {
		rowsAdvertising := sqlmock.NewRows([]string{"name", "description", "mainphoto", "cost"}).AddRow(
			"Продам гараж", "Очень хороший", "sfd/test.jpeg", 1432).AddRow(
			"Продам гараж", "Очень хороший", "sfd/test.jpeg", 1432)

		sqlxDb := sqlx.NewDb(db, "sqlmock")

		rep := NewAdvertisingRepository(sqlxDb)
		query := rep.GenerateQueryForGetAdvertisingList("cost", "true")

		advertisingListTest := advertisingModel.AdvertisingList{List: []advertisingModel.Advertising{{
			Name: "Продам гараж", Description: "Очень хороший",
			MainPhoto: "sfd/test.jpeg", Cost: 1432,
		}, {Name: "Продам гараж", Description: "Очень хороший",
			MainPhoto: "sfd/test.jpeg", Cost: 1432}}}

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(rowsAdvertising)

		advertisingListActual, err := rep.ListAdvertising("cost", "true")
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

		_, err := rep.ListAdvertising("cost", "true")
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
