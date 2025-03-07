/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package collection

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CollectionCmd represents the collection command
var CollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("collection called")
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
}
