package locationareas

import (
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreas(url string) (data []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

func GetPokemonFromLocationArea(area string) (data []byte, err error) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)

	res, err := http.Get(fullUrl)
	if err != nil {
		return
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
