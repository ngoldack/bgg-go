package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bgg-go",
	Short: "cli to interact with boardgamegeek",
	Long:  `bgg-go is a go cli & library for boardgamegeek.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
