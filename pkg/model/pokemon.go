package model

import "github.com/mtslzr/pokeapi-go/structs"

type LocalPokemon struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"`
	LocationAreaEncounters string `json:"location_area_encounters"`
}

func newLocalPokemon(id int, name string, experience, height int, location string) *LocalPokemon {
	return &LocalPokemon{
		ID:                     id,
		Name:                   name,
		BaseExperience:         experience,
		Height:                 height,
		LocationAreaEncounters: location,
	}
}

func PokemonDaoToPokemon(pokemonDao structs.Pokemon) *LocalPokemon {
	return newLocalPokemon(
		pokemonDao.ID,
		pokemonDao.Name,
		pokemonDao.BaseExperience,
		pokemonDao.Height,
		pokemonDao.LocationAreaEncounters,
	)
}
