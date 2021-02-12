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
	// max length: 200
	Name string `json:"name" db:"name" validate:"max=200"`
	// max length: 1000
	Description string `json:"description,omitempty" db:"description" validate:"max=1000"`
	// max length: 3
	Photos []string `json:"photos,omitempty" db:"photos" validate:"max=3"`
	Cost   int      `json:"cost" db:"cost"`
}

// easyjson:json
type AdvertisingListItem struct {
	Name      string `json:"name" db:"name" `
	MainPhoto string `json:"mainPhoto" db:"mainphoto"`
	Cost      int    `json:"cost" db:"cost"`
}

// easyjson:json
type AdvertisingList struct {
	List  []AdvertisingListItem `json:"list"`
	Page  Page                  `json:"page" mapstructure:"page"`
	Links Links                 `json:"links" mapstructure:"links"`
}

// easyjson:json
type Links struct {
	NextUrl string `json:"next" mapstructure:"next"`
	PrevUrl string `json:"prev" mapstructure:"prev"`
	Last    string `json:"last" mapstructure:"last"`
}

// easyjson:json
type Page struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	LastPage    int `json:"lastPage"`
}
