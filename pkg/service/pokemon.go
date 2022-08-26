package service

import (
	"log"

	"github.com/BabouZ17/pokemon-cli/pkg/model"
	"github.com/mtslzr/pokeapi-go"
)

type Service interface {
	GetPokemon(name string) model.Pokemon
	GetPokemons(from, to int) []model.Pokemon
	GetMoves(name string) []model.Move
	GetAbilities(name string) []model.Ability
}

type PokemonService struct{}

func (ps PokemonService) GetPokemon(name string) model.Pokemon {
	pokemonDto, err := pokeapi.Pokemon(name)
	if err != nil {
		log.Fatal(err)
	}
	pokemon := model.PokemonDtoToPokemon(pokemonDto)
	return *pokemon
}

func (ps PokemonService) GetPokemons(from, to int) []model.Pokemon {
	pokemonsDto, err := pokeapi.Resource("pokemon", from, to)
	if err != nil {
		log.Fatal(err)
	}

	var pokemons []model.Pokemon
	for _, pokemonDto := range pokemonsDto.Results {
		result, err := pokeapi.Pokemon(pokemonDto.Name)
		if err != nil {
			log.Fatal(err)
		}
		pokemons = append(pokemons, *model.PokemonDtoToPokemon(result))
	}
	return pokemons
}

func (ps PokemonService) GetMoves(name string) []model.Move {
	pokemonDto, err := pokeapi.Pokemon(name)
	if err != nil {
		log.Fatal(err)
	}

	var moves []model.Move
	for _, move := range pokemonDto.Moves {
		result, err := pokeapi.Move(move.Move.Name)
		if err != nil {
			log.Fatal(err)
		}
		moves = append(moves, *model.MoveDTOtoMove(result))
	}
	return moves
}

func (ps PokemonService) GetAbilities(name string) []model.Ability {
	pokemonDto, err := pokeapi.Pokemon(name)
	if err != nil {
		log.Fatal(err)
	}

	var abilities []model.Ability
	for _, ability := range pokemonDto.Abilities {
		result, err := pokeapi.Ability(ability.Ability.Name)
		if err != nil {
			log.Fatal(err)
		}
		abilities = append(abilities, *model.AbilityDTOtoAbility(result))
	}
	return abilities
}
