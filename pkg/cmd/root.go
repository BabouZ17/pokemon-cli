package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pokemon-cli",
	Short: "pokemon-cli is a simple cli to query the pokeapi",
	Long:  `More info on https://pokeapi.co/`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	getPokemonCmd.Flags().StringVarP(&name, "name", "n", "", "Pokemon name or id")
	getPokemonCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(getPokemonCmd)

	getPokemonsCmd.Flags().IntVarP(&from, "from", "f", 0, "Start from (0 to 905)")
	getPokemonsCmd.Flags().IntVarP(&to, "to", "t", 20, "Go to (0 to 905)")
	rootCmd.AddCommand(getPokemonsCmd)

	getPokemonMovesCmd.Flags().StringVarP(&name, "name", "n", "", "Pokemon name or id")
	getPokemonMovesCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(getPokemonMovesCmd)

	getPokemonAbilitiesCmd.Flags().StringVarP(&name, "name", "n", "", "Pokemon name or id")
	getPokemonAbilitiesCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(getPokemonAbilitiesCmd)

	rootCmd.AddCommand(versionCmd)
}
