package model

import "github.com/mtslzr/pokeapi-go/structs"

type Move struct {
	ID           int         `json:"id"`
	Name         string      `json:"pound"`
	Accuracy     int         `json:"accuracy"`
	EffectChance interface{} `json:"effect_chance"`
	PowerPoints  int         `json:"power_points"`
	Priority     int         `json:"priority"`
	Power        int         `json:"power"`
}

func newMove(id int, name string, accuracy int, effectChance interface{}, powerPoints, priority, power int) *Move {
	return &Move{
		ID:           id,
		Name:         name,
		Accuracy:     accuracy,
		EffectChance: effectChance,
		PowerPoints:  powerPoints,
		Priority:     priority,
		Power:        power,
	}
}

func MoveDTOtoMove(move structs.Move) *Move {
	return newMove(
		move.ID,
		move.Name,
		move.Accuracy,
		move.EffectChance,
		move.Pp,
		move.Priority,
		move.Power,
	)
}
