package model

import "github.com/mtslzr/pokeapi-go/structs"

type Ability struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IsMainSeries bool   `json:"is_main_series"`
}

func newAbility(id int, name string, is_main_series bool) *Ability {
	return &Ability{
		ID:           id,
		Name:         name,
		IsMainSeries: is_main_series,
	}
}

func AbilityDTOtoAbility(abilityDTO structs.Ability) *Ability {
	return newAbility(
		abilityDTO.ID,
		abilityDTO.Name,
		abilityDTO.IsMainSeries,
	)
}
