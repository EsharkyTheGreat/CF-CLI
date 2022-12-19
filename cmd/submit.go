package cmd

import (
	"github.com/spf13/cobra"
)


func submitSolution() {
   
}

var submit = &cobra.Command{
	Use:   "submit",
	Short: "submit solution cpp file",
	Run: func(cmd *cobra.Command, args []string) {
    submitSolution()  
	},
}

func init() {
	rootCmd.AddCommand(submit)
}
