package locationareas

type LocationAreaList struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []LocationAreaObj `json:"results"`
}

type LocationAreaObj struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
