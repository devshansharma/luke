/*
Copyright © 2025 Shantanu Sharma sharmashan0805@gmail.com
*/
package cmd

import (
	"os"

	"github.com/devshansharma/luke/cmd/collection"
	"github.com/devshansharma/luke/cmd/environment"
	"github.com/devshansharma/luke/pkg/writer"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "luke",
	Short: "A fast, open-source CLI alternative to Postman for automating API requests.",
	Long: `This CLI tool is designed to be an efficient and scriptable alternative to Postman, 
enabling developers to automate API testing and request workflows directly from the command line. 

It supports features like environment variables, request chaining, response validation, 
and structured output formats. 

Built with automation in mind, this open-source project aims to simplify API testing 
for developers, DevOps engineers, and testers. 🚀`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(collection.CollectionCmd)
	rootCmd.AddCommand(environment.EnvironmentCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.luke.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().BoolP("silent", "s", false, "silent output")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringP("output", "o", "", "output file")
	rootCmd.PersistentFlags().StringP("log-file", "l", "", "log file")
}

func parseGlobalFlags(cmd *cobra.Command, args []string) *writer.OutputWriter {
	cmd.ParseFlags(args)

	output, _ := cmd.Flags().GetString("output")
	logFile, _ := cmd.Flags().GetString("log-file")
	silent, _ := cmd.Flags().GetBool("silent")
	verbose, _ := cmd.Flags().GetBool("verbose")

	return writer.InitOutputWriter(output, logFile, silent, verbose)
}
