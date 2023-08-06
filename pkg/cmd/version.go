package cmd

import (
	"fmt"
	"log"

	"github.com/BabouZ17/pokemon-cli/pkg/version"
	"github.com/spf13/cobra"
)

const VERSION_PATH = "./VERSION"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		version, err := version.ReadVersion(VERSION_PATH)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(version)
	},
}
