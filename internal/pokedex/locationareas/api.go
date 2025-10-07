package locationareas

import (
	"encoding/json"
	"net/http"
)

func GetLocationAreas(url string) (loc LocationAreaList, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if errDec := decoder.Decode(&loc); errDec != nil {
		return
	}
	return
}
