package pokemon

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
)

func GetPokemonData(data []byte) (Pokemon, error) {
	var pok Pokemon
	if err := json.Unmarshal(data, &pok); err != nil {
		return pok, err
	}
	return pok, nil
}

func CatchPokemon(baseExp int) bool {
	r := rand.IntN(baseExp * 2)
	fmt.Println(r)
	topLimit := int((float64(baseExp) * 0.2) + float64(baseExp))
	if r >= baseExp && r <= topLimit {
		return true
	}
	return false
}

func (p Pokemon) InspectPokemon() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %v\n", p.Height)
	fmt.Printf("Weight: %v\n", p.Weight)
	fmt.Println("Stats:")
	for _, s := range p.Stats {
		fmt.Printf("\t-%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("\t-%s\n", t.Type.Name)
	}
}
