package cmd

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

const PROBLEMSET_URL string = "https://codeforces.com/problemset/problem"

func createProblemDir(problemname string) {
  curr_dir ,err := os.Getwd()
  if err != nil {
    fmt.Println("Error Getting Current Working Directory")
    return
  }
  fmt.Println(curr_dir)
  x := strings.Split(problemname,"/")
  contest_num,problem_num := x[0],x[1]
  if err := os.Mkdir(fmt.Sprint(curr_dir,"/",contest_num),os.ModePerm); err != nil {
    fmt.Println("Directory already exists")
  }
  if err := os.Mkdir(fmt.Sprint(curr_dir,"/",contest_num,"/",problem_num),os.ModePerm); err != nil {
    fmt.Println("Error Making Directory")
    return
  }
  if err := os.Chdir(fmt.Sprint(curr_dir,"/",contest_num,"/",problem_num)); err != nil {
    fmt.Println("Error Changing Directory")
    return
  }
}

func writeSamples(input [][]byte,output [][]byte) {
  for i:=1 ; i <= len(input) ; i++ {
    os.WriteFile(fmt.Sprintf("input%d.txt",i),input[i-1],0644)
    os.WriteFile(fmt.Sprintf("output%d.txt",i),output[i-1],0644)
  }
}

func ParseProblem(problem string) {
  url := fmt.Sprintf("%s/%s",PROBLEMSET_URL,problem)     
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println("Error Fetching Problem");
  }
  defer resp.Body.Close()
  body_bytes,err := io.ReadAll(resp.Body)
  body := body_bytes
  inp_regex := regexp.MustCompile(`class="input"[\s\S]*?<pre>([\s\S]*?)</pre>`)
  out_regex := regexp.MustCompile(`class="output"[\s\S]*?<pre>([\s\S]*?)</pre>`)
  inp_match := inp_regex.FindAllSubmatch(body,-1);
  out_match := out_regex.FindAllSubmatch(body,-1);
  fmt.Println(inp_match)
  fmt.Println(out_match)
  newline := regexp.MustCompile(`<[\s/br]+?>`)
	filter := func(src []byte) []byte {
		src = newline.ReplaceAll(src, []byte("\n"))
		s := html.UnescapeString(string(src))
		return []byte(strings.TrimSpace(s) + "\n")
	}
  var input [][]byte;
  var output [][]byte;
	for i := 0; i < len(inp_match); i++ {
		input = append(input, filter(inp_match[i][1]))
		output = append(output, filter(out_match[i][1]))
	}
  for i:=0; i < len(input); i++ {
    fmt.Println(string(input[i]))
  }
  for i:=0; i < len(output); i++ {
    fmt.Println(string(output[i]))
  }
  createProblemDir(problem)
  writeSamples(input,output)
}

var parser = &cobra.Command{
	Use:   "parse",
	Short: "parse given problem and get input and output",
  Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
    println(len(args))
    println(args)
	  ParseProblem(args[0])
	},
}

func init() {
	rootCmd.AddCommand(parser)
}
