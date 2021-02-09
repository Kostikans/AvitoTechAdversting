package advertisingModel

type Adverting struct {
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Photos      []string `json:"photos" db:"photos"`
	Cost        int      `json:"cost" db:"cost"`
}

type Cursor struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
