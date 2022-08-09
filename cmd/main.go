package main

import (
	"encoding/json"
	"fmt"

	"github.com/BabouZ17/pokemon-cli/pkg/model"
	"github.com/mtslzr/pokeapi-go"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func main() {
	pokemonDao, _ := pokeapi.Pokemon("pikachu")
	pokemon := model.PokemonDaoToPokemon(pokemonDao)
	formated, _ := PrettyStruct(pokemon)
	fmt.Print(formated)
}
