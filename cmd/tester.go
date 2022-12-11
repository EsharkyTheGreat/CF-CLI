package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func cleanup() {
  command := exec.Command("rm","a.out")
  command.Run()
}

func compileProg() {
  command := exec.Command("g++","solve.cpp") 
  command.Run()
}

func RunProg() {
  command := exec.Command("./a.out")
  var o bytes.Buffer
  output := io.Writer(&o)
  buf,_ := os.ReadFile("input.txt")
  inp := bytes.Split(buf,seperator)
  i := bytes.NewReader(inp[0])
  command.Stdin = i;
  command.Stdout = output 
  command.Run()
  fmt.Println("Output")
  fmt.Println(o.String())
}

var tester = &cobra.Command{
	Use:   "test",
	Short: "test the program against inputs",
	Run: func(cmd *cobra.Command, args []string) {
    compileProg()
    RunProg()
    cleanup()
	},
}

func init() {
	rootCmd.AddCommand(tester)
}
