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
type AdvertisingList struct {
	List   []Advertising `json:"list"`
	Cursor Cursor        `json:"cursor"`
}

// easyjson:json
type Cursor struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
