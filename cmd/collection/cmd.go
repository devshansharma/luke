/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package collection

import (
	"fmt"
	"os"

	"github.com/devshansharma/luke/internal/handlers"
	"github.com/devshansharma/luke/pkg/utils"
	"github.com/spf13/cobra"
)

// CollectionCmd represents the collection command
var CollectionCmd = &cobra.Command{
	Use:               "collection",
	Short:             "To see list of folders and items in a collection",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: handlers.AddFolderCompletion,
	Long: `To see list of folders and items
luke collection <collection-name>

To see list of items inside a folder:
luke collection <collection-name> --folder <folder-name>
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.ParseFlags(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		cfg := handlers.CollectionDetailsConfig{
			Name: args[0],
		}

		folder, _ := cmd.Flags().GetString("folder")
		cfg.FolderName = folder

		err = utils.ValidateFlags(cmd, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = handlers.CollectionDetails(&cfg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// collection/collectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// collection/collectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	CollectionCmd.Flags().StringP("folder", "f", "", "list items in a folder")
}
