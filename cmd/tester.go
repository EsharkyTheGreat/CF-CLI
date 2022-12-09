package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func compileProg() {
  command := exec.Command("g++","solve.cpp") 
  command.Run()
}

func RunProg() {
  command := exec.Command("./a.out")
  f,_ := os.Open("input1.txt")
  var o bytes.Buffer
  output := io.Writer(&o)
  command.Stdin = f;
  command.Stdout = output 
  command.Run()
  fmt.Println(o.String())
}

var tester = &cobra.Command{
	Use:   "test",
	Short: "test the program against inputs",
	Run: func(cmd *cobra.Command, args []string) {
    compileProg()
    RunProg()
	},
}

func init() {
	rootCmd.AddCommand(tester)
}
