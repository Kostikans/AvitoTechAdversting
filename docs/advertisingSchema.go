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

//swagger:parameters GetAdvertising
type GetAdvertisingParam struct {
	//in:path
	Body int `json:"id"`
}

//swagger:parameters ListAdvertising
type ListAdvertisingParam struct {
	//in:query
	Sort string `json:"sort"`
	//in:query
	Desc bool `json:"desc"`
}

type CursorWrap struct {
	HasNext string `json:"next"`
	HasPrev string `json:"prev"`
}

//swagger:response ListAdvertising
type ListAdvertisingResponse struct {
	//in:body
	Body ListAdvertisingResponseWrap
}

type ListAdvertisingResponseWrap struct {
	Data []GetAdvertisingWrapModel `json:"data,omitempty"`
	Code int                       `json:"code"`
}

type AddAdvertisingParamWrap struct {
	//max length 200
	Name string `json:"name" db:"name" validate:"max=200"`
	//max length 1000
	Description string `json:"description,omitempty" db:"description" validate:"max=1000"`
	//max length 3
	Photos []string `json:"photos,omitempty" db:"photos" validate:"max=3"`
	Cost   int      `json:"cost" `
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

//swagger:parameters GetAdvertising
type ListAGetAdvertisingParam struct {
	//in:query
	//description or photos, or both with ',' splitter; example(description,photos)
	Body string `json:"fields"`
}
