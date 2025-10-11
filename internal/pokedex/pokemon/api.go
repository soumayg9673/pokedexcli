package pokemon

import (
	"fmt"
	"io"
	"net/http"
)

func GetPokemon(name string) (data []byte, err error) {
	fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	res, err := http.Get(fullUrl)
	if err != nil {
		return
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
