package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/ngoldack/bgg-go/bgg"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		query, err := cmd.Flags().GetString("query")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		c := bgg.New()
		results, err := c.Search(query, nil)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("ID", "Name", "Published")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for i, result := range results {
			fmt.Printf("%d: %s - %s\n", i, result.ID, result.Name.Value)
			tbl.AddRow(result.ID, result.Name.Value, result.Yearpublished.Value)
		}

		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	searchCmd.Flags().StringP("query", "q", "", "The query you want to search for")
}
