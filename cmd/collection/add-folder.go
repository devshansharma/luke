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

// addFolderCmd represents the list command
var addFolderCmd = &cobra.Command{
	Use:               "add-folder",
	Short:             "To add folder",
	Args:              cobra.ExactArgs(2),
	ValidArgsFunction: handlers.AddFolderCompletion,
	Long: `To add a folder

luke collection add-folder <collection_name> <folder_name>	

`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.ParseFlags(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		cfg := handlers.AddFolderConfig{
			Name:       args[0],
			FolderName: args[1],
		}

		err = utils.ValidateFlags(cmd, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = handlers.AddFolder(&cfg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	CollectionCmd.AddCommand(addFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
