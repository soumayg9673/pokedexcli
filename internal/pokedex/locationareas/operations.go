package locationareas

import (
	"encoding/json"
	"fmt"
)

func (l LocationAreaList) PrintLocationAreaResultsName() {
	for _, v := range l.Results {
		fmt.Println(v.Name)
	}
}

func GetLocationAreasData(data []byte) (loc LocationAreaList, err error) {
	if errJson := json.Unmarshal(data, &loc); errJson != nil {
		return
	}
	return
}

func (l LocationArea) PrintPokemonsFromLocationAreaResult(area string) {
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", area)
	for _, v := range l.PokemonEncounters {
		fmt.Printf("- %s\n", v.Pokemon.Name)
	}
}

func GetPokemonsFromLocationAreaData(data []byte) (loc LocationArea, err error) {
	if errJson := json.Unmarshal(data, &loc); errJson != nil {
		return
	}
	return
}
