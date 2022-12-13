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
  buf,_ := os.ReadFile("input.txt")
  out_bf,_ := os.ReadFile("output.txt")
  outx := bytes.Split(out_bf,seperator)
  inp := bytes.Split(buf,seperator)
  for i := 0;i < len(inp)-1;i++ {
    command := exec.Command("./a.out")
    println("Iteration ",i);
    print(string(inp[i]))
    var o bytes.Buffer
    output := io.Writer(&o)
    ib := bytes.NewReader(inp[i])
    command.Stdin = ib;
    command.Stdout = output 
    command.Run()
    fmt.Println("Output")
    fmt.Print(o.String())
    fmt.Println("Correct Output")
    fmt.Print(string(outx[i]));
    fmt.Print(string(seperator))
  }
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
