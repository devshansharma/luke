/*
Copyright © 2025 Shantanu Sharma sharmashan0805@gmail.com
*/
package collection

import (
	"fmt"
	"os"

	"github.com/devshansharma/luke/internal/collection"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		writer := parseGlobalFlags(cmd, args)
		defer writer.Close()

		c := collection.Collection{
			Name: name,
		}

		c.Description, _ = cmd.Flags().GetString("description")

		err := collection.AddHandler(c)
		if err != nil {
			writer.Error(err.Error())
			os.Exit(1)
		}

		writer.Response(fmt.Sprintf("Collection created successfully %q", name))
	},
}

func init() {
	CollectionCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	addCmd.Flags().StringP("description", "d", "", "Description for the collection")
	// addCmd.MarkFlagRequired("description")
}
