/*
Copyright © 2025 Shantanu Sharma sharmashan0805@gmail.com
*/
package environment

import (
	"fmt"
	"os"

	"github.com/devshansharma/luke/internal/env"
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

		cfg := env.Config{
			Name: name,
		}

		cfg.Description, _ = cmd.Flags().GetString("description")

		err := env.Handler(cfg)
		if err != nil {
			writer.Error(err.Error())
			os.Exit(1)
		}

		writer.Response(fmt.Sprintf("Environment created successfully %q", name))
	},
}

func init() {
	EnvironmentCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("description", "d", "", "Description for the environment")
}
