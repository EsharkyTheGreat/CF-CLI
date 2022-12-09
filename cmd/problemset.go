package cmd

import (
	"github.com/EsharkyTheGreat/cf-cli/util"

	"github.com/spf13/cobra"
)

var problemset = &cobra.Command{
	Use:   "problemset",
	Short: "Fetches Problemset from Codeforces",
	Run: func(cmd *cobra.Command, args []string) {
		util.GetProblemset()
	},
}

func init() {
	rootCmd.AddCommand(problemset)
}
