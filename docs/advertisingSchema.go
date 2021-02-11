package docs

import advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"

//swagger:response AddAdvertising
type AddAdvertisingResponse struct {
	//in:body
	Body struct {
		Data advertisingModel.AdvertisingID `json:"data,omitempty"`
		Code int                            `json:"code"`
	}
}

//swagger:parameters AddAdvertising
type AddAdvertisingParam struct {
	//in:body
	Body advertisingModel.AdvertisingAdd
}

//swagger:parameters GetAdvertising
type GetAdvertisingParam struct {
	//in:path
	ID int `json:"id"`
	// description or photos, or both with ',' splitter; example(description,photos)
	//in:query
	Fields string `json:"fields"`
}

//swagger:parameters ListAdvertising
type ListAdvertisingParam struct {
	//"cost" or "created"
	//in:query
	Sort string `json:"sort"`
	//in:query
	Desc bool `json:"desc"`
	// by default 1
	//in:query
	Page int `json:"page"`
}

//swagger:response ListAdvertising
type ListAdvertisingResponse struct {
	//in: Body
	Body struct {
		Data advertisingModel.AdvertisingList `json:"data,omitempty"`
		Code int                              `json:"code"`
	}
}

type GetAdvertising struct {
	Name string `json:"name" db:"name" validate:"max=200"`
	//optional
	Description string `json:"description,omitempty" db:"description" validate:"max=1000"`
	MainPhoto   string `json:"mainPhoto" `
	//optional
	Photos []string `json:"photos,omitempty" db:"photos" validate:"max=3"`
	Cost   int      `json:"cost" `
}

//swagger:response GetAdvertising
type GetAdvertisingResponse struct {
	//in: Body
	Body struct {
		Data GetAdvertising `json:"data,omitempty"`
		Code int            `json:"code"`
	}
}
