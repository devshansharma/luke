/*
Copyright © 2025 Shantanu Sharma sharmashan0805@gmail.com
*/
package cmd

import (
	"os"

	"github.com/devshansharma/luke/internal/request"
	"github.com/spf13/cobra"
)

// requestCmd represents the request command
var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "To make a request",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := request.RequestConfig{
			URL: args[0],
		}

		writer := parseGlobalFlags(cmd, args)
		defer writer.Close()

		err := request.ParseFlags(cmd, &cfg)
		if err != nil {
			writer.Error(err.Error())
			os.Exit(1)
		}

		err = request.HandleRequest(cfg)
		if err != nil {
			writer.Error(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(requestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// requestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// requestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	requestCmd.Flags().StringP("method", "X", "GET", "Request method")
	requestCmd.Flags().StringP("data", "d", "", "Request body")
	requestCmd.Flags().StringSliceP("header", "H", []string{}, "Request headers")
	requestCmd.Flags().Int64P("timeout", "t", 60000, "Request timeout in ms")
}
