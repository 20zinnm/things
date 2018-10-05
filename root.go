package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "things",
	Short: "Runs various programs that control or utilise physical things.",
	Long: `Developed for the Information Engineering class at St. Mark's School of Texas.`,
}

// execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// verbose logging flag
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose logging (useful for debugging)")
}