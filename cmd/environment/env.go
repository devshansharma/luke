/*
Copyright © 2025 Shantanu Sharma sharmashan0805@gmail.com
*/
package environment

import (
	"fmt"

	"github.com/devshansharma/luke/pkg/writer"
	"github.com/spf13/cobra"
)

// EnvironmentCmd represents the collection command
var EnvironmentCmd = &cobra.Command{
	Use:   "env",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("env called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func parseGlobalFlags(cmd *cobra.Command, args []string) *writer.OutputWriter {
	cmd.ParseFlags(args)

	output, _ := cmd.Flags().GetString("output")
	logFile, _ := cmd.Flags().GetString("log-file")
	silent, _ := cmd.Flags().GetBool("silent")
	verbose, _ := cmd.Flags().GetBool("verbose")

	return writer.InitOutputWriter(output, logFile, silent, verbose)
}
