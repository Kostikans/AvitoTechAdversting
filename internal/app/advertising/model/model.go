//go:generate easyjson -all model.go
package advertisingModel

// easyjson:json
type Advertising struct {
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Photos      []string `json:"photos" db:"photos"`
	Cost        int      `json:"cost" db:"cost"`
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
