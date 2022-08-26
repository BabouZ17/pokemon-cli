package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/BabouZ17/pokemon-cli/pkg/prettify"
	"github.com/BabouZ17/pokemon-cli/pkg/service"
	"github.com/spf13/cobra"
)

const (
	TOTAL_POKEMONS = 905
	MAX_PAGINATION = 25
)

var (
	name           string
	from, to       int
	pokemonService = service.PokemonService{}
	getPokemonCmd  = &cobra.Command{
		Use:   "pokemon",
		Short: "Get the information of a pokemon",
		Long:  "Get the information of a pokemon",
		Run: func(cmd *cobra.Command, args []string) {
			pokemon := pokemonService.GetPokemon(name)
			fmt.Print(prettify.PrettifyStruct(pokemon))
		},
	}

	getPokemonsCmd = &cobra.Command{
		Use:   "pokemons",
		Short: "Get the first n pokemons information",
		Long:  "Get the first n pokemons information",
		Args: func(cmd *cobra.Command, args []string) error {
			if to < from || to > TOTAL_POKEMONS || from > TOTAL_POKEMONS {
				return errors.New("invalid from or to args values received")
			}

			if to-from > MAX_PAGINATION {
				msg := "invalid gap between to and from (max is " + strconv.Itoa(MAX_PAGINATION) + ")"
				return errors.New(msg)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			pokemons := pokemonService.GetPokemons(from, to)
			fmt.Print(prettify.PrettifyStruct(pokemons))
		},
	}

	getPokemonAbilitiesCmd = &cobra.Command{
		Use:   "abilities",
		Short: "Get the abilities of a pokemon",
		Long:  "Get the abilities of a pokemon",
		Run: func(cmd *cobra.Command, args []string) {
			abilities := pokemonService.GetAbilities(name)
			fmt.Print(prettify.PrettifyStruct(abilities))
		},
	}

	getPokemonMovesCmd = &cobra.Command{
		Use:   "moves",
		Short: "Get the moves of a pokemon",
		Long:  "Get the moves of a pokemon",
		Run: func(cmd *cobra.Command, args []string) {
			moves := pokemonService.GetMoves(name)
			fmt.Print(prettify.PrettifyStruct(moves))
		},
	}
)
