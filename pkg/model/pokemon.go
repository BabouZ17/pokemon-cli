package model

import "github.com/mtslzr/pokeapi-go/structs"

type Pokemon struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"`
	Weight                 int    `json:"weight"`
	Order                  int    `json:"order"`
	LocationAreaEncounters string `json:"location_area_encounters"`
}

func newPokemon(id int, name string, experience, height, weight, order int, location string) *Pokemon {
	return &Pokemon{
		ID:                     id,
		Name:                   name,
		BaseExperience:         experience,
		Height:                 height,
		Weight:                 weight,
		Order:                  order,
		LocationAreaEncounters: location,
	}
}

func PokemonDtoToPokemon(pokemonDTO structs.Pokemon) *Pokemon {
	return newPokemon(
		pokemonDTO.ID,
		pokemonDTO.Name,
		pokemonDTO.BaseExperience,
		pokemonDTO.Height,
		pokemonDTO.Weight,
		pokemonDTO.Order,
		pokemonDTO.LocationAreaEncounters,
	)
}
