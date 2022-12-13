package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var templateString string = `#include <bits/stdc++.h>
#include <stdlib.h>

#define ll long long
using namespace std;
void solve() {
  
}
int main() {
  ll t;
  cin >> t;
  while(t--) {
    solve();
  }
}
`

func createTemplateFile() {
	f, err := os.Create("solve.cpp")
	if err != nil {
		fmt.Println("Could'nt create solve.cpp")
    return
	}
  defer f.Close()
  f.WriteString(templateString)
}

var gen = &cobra.Command{
	Use:   "gen",
	Short: "gen template cpp file",
	Run: func(cmd *cobra.Command, args []string) {
    createTemplateFile()
	},
}

func init() {
	rootCmd.AddCommand(gen)
}
