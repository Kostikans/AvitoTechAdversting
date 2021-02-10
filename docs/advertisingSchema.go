package docs

//swagger:response AddAdvertising
type AddAdvertisingDoc struct {
	//in:body
	Body AddAdvertisingWrap
}

type AddAdvertisingWrap struct {
	Data int `json:"data,omitempty"`
	Code int `json:"code"`
}

//swagger:parameters AddAdvertising
type AddAdvertisingParam struct {
	//in:body
	Body AddAdvertisingParamWrap
}

type AddAdvertisingParamWrap struct {
	Name        string   `json:"name" db:"name" validate:"max=200"`
	Description string   `json:"description,omitempty" db:"description" validate:"max=1000"`
	Photos      []string `json:"photos,omitempty" db:"photos" validate:"max=3"`
	Cost        int      `json:"cost" `
}

type GetAdvertisingWrapModel struct {
	Name        string   `json:"name" db:"name" validate:"max=200"`
	Description string   `json:"description,omitempty" db:"description" validate:"max=1000"`
	MainPhoto   string   `json:"mainPhoto" `
	Photos      []string `json:"photos,omitempty" db:"photos" validate:"max=3"`
	Cost        int      `json:"cost" `
}

type GetAdvertisingWrap struct {
	Data GetAdvertisingWrapModel `json:"data,omitempty"`
	Code int                     `json:"code"`
}

//swagger:response GetAdvertising
type GetAdvertising struct {
	//in: body
	Body GetAdvertisingWrap
}
