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

// addItemCmd represents the list command
var addItemCmd = &cobra.Command{
	Use:               "add-item",
	Short:             "To add item to a collection or a folder",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: handlers.AddFolderCompletion,
	Long: `To add a folder

luke collection add-item <collection_name> --name <item_name>

luke collection add-item <collection_name> --folder <folder_name> --name <item_name>

`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.ParseFlags(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		cfg := handlers.AddItemConfig{
			Name: args[0],
		}

		err = utils.ValidateFlags(cmd, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		folder, _ := cmd.Flags().GetString("folder")
		name, _ := cmd.Flags().GetString("name")

		cfg.FolderName = folder
		cfg.ItemName = name

		err = handlers.AddItem(&cfg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	CollectionCmd.AddCommand(addItemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addItemCmd.Flags().StringP("name", "n", "", "item name")
	addItemCmd.Flags().StringP("folder", "f", "", "folder name")
}
