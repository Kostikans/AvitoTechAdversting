package advertisingRepository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Kostikans/AvitoTechadvertising/configs"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"

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
func (advRepo *AdvertisingRepository) AddAdvertising(advertising advertisingModel.AdvertisingAdd) (advertisingModel.AdvertisingID, error) {
	var advertisingID advertisingModel.AdvertisingID
	err := advRepo.db.QueryRow(AddAdvertising, advertising.Name, advertising.Description, advertising.Photos[0], pq.Array(advertising.Photos), advertising.Cost).Scan(&advertisingID.AdvID)
	if err != nil {
		return advertisingID, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}
	_, err = advRepo.db.Query(IncrementAdvertisingCount)
	if err != nil {
		return advertisingID, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}
	return advertisingID, nil
}
func (advRepo *AdvertisingRepository) GetAdvertising(advertisingID int, fields string) (advertisingModel.Advertising, error) {
	var advertising advertisingModel.Advertising
	containDescription := strings.Contains(fields, "description")
	containPhotos := strings.Contains(fields, "photos")

	var err error
	if containDescription && containPhotos {
		err = advRepo.db.QueryRow(GetAdvertisingWithPhotosAndDescription, advertisingID).
			Scan(&advertising.Name, &advertising.Cost, &advertising.MainPhoto, pq.Array(&advertising.Photos), &advertising.Description)
	} else if containDescription {
		err = advRepo.db.QueryRow(GetAdvertisingWithDescription, advertisingID).Scan(&advertising.Name, &advertising.Cost, &advertising.MainPhoto, &advertising.Description)
	} else if containPhotos {
		err = advRepo.db.QueryRow(GetAdvertisingWithPhotos, advertisingID).Scan(&advertising.Name, &advertising.Cost, &advertising.MainPhoto, pq.Array(&advertising.Photos))
	} else {
		err = advRepo.db.QueryRow(GetAdvertising, advertisingID).Scan(&advertising.Name, &advertising.Cost, &advertising.MainPhoto)
	}

	if err != nil {
		return advertising, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}

	return advertising, nil
}

func (advRepo *AdvertisingRepository) ListAdvertising(sort string, desc string, page int) (advertisingModel.AdvertisingList, error) {
	var advertisingList advertisingModel.AdvertisingList
	var advertising []advertisingModel.AdvertisingListItem

	elementsPerPage := configs.ElementsPerPage
	offset := (page - 1) * elementsPerPage

	err := advRepo.db.Select(&advertising, advRepo.GenerateQueryForGetAdvertisingList(sort, desc), offset)
	if err != nil {
		return advertisingList, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}

	advertisingList.List = advertising
	advertisingList.Page.CurrentPage = page
	advertisingList.Page.PerPage = elementsPerPage

	var elementsCount int
	err = advRepo.db.QueryRow(GetPageCount).Scan(&elementsCount)
	if err != nil {
		return advertisingList, customError.NewCustomError(err, responseCodes.ServerInternalError, 1)
	}

	advertisingList.Page.LastPage = (elementsCount-1)/elementsPerPage + 1
	return advertisingList, nil
}

func (advRepo *AdvertisingRepository) GenerateQueryForGetAdvertisingList(sort string, desc string) string {
	query := "SELECT name,mainPhoto,cost from advertising "

	if sort == "cost" {
		query += "order by cost"
	} else if sort == "created" {
		query += "order by created"
	}

	if desc == "true" {
		query += " DESC"
	} else if desc == "false" {
		query += " ASC"
	}

	query += fmt.Sprintf(" OFFSET $1 LIMIT %d", configs.ElementsPerPage)

	return query
}

func (advRepo *AdvertisingRepository) CheckAdvertisingExist(advertisingID int) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT exists (%s)", CheckAdvertisingExist)
	err := advRepo.db.QueryRow(query, advertisingID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return exists, customError.NewCustomError(err, serverError.ServerInternalError, 1)
	}
	return exists, nil
}
