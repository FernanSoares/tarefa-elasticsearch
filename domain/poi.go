package domain

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type POI struct {
	ID            string   `json:"id"`
	OSMID         string   `json:"osm_id"`
	Type          string   `json:"type"`
	CategoryName  string   `json:"category_name"`
	CategoryValue string   `json:"category_value"`
	Name          string   `json:"name"`
	Location      Location `json:"location"`
}
