package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cf <subcommand>",
	Short: "Short Description",
	Long:  "CodeForces CLI for automating Codeforces Setup and Submission",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You just entered the main command")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
