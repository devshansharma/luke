/*
Copyright © 2025 Shantanu Sharma sharmashan0805@gmail.com
*/
package collection

import (
	"os"
	"strconv"

	"github.com/devshansharma/luke/internal/collection"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		writer := parseGlobalFlags(cmd, args)
		defer writer.Close()

		id := args[0]

		idNum, err := strconv.Atoi(id)
		if err != nil {
			writer.Error(err.Error())
			os.Exit(1)
		}

		err = collection.DeleteHandler(idNum)
		if err != nil {
			writer.Error(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	CollectionCmd.AddCommand(deleteCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
