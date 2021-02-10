//go:generate easyjson -all model.go
package advertisingModel

// easyjson:json
type Advertising struct {
	Name        string   `json:"name" db:"name" validate:"max=200"`
	Description string   `json:"description,omitempty" db:"description" validate:"max=1000"`
	MainPhoto   string   `json:"mainPhoto" db:"mainphoto"`
	Photos      []string `json:"photos,omitempty" db:"photos" validate:"max=3"`
	Cost        int      `json:"cost" db:"cost"`
	Created     string   `json:"created,omitempty" db:"created"`
}

// easyjson:json
type AdvertisingID struct {
	AdvID int `json:"advertisingID" db:"advertising_id" mapstructure:"advertisingID"`
}

// easyjson:json
type AdvertisingAdd struct {
	Name        string   `json:"name" db:"name" validate:"max=200"`
	Description string   `json:"description,omitempty" db:"description" validate:"max=1000"`
	Photos      []string `json:"photos,omitempty" db:"photos" validate:"max=3"`
	Cost        int      `json:"cost" db:"cost"`
}

// easyjson:json
type AdvertisingList struct {
	List []Advertising `json:"list"`
}

// easyjson:json
type Cursor struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
