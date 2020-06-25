package storage

type Datum struct {
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Datetime string  `json:"datetime"`
}
