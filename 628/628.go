// UVa 628 - Passwords

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	out  *os.File
	dict []string
	rule string
)

func dfs(level int, tokens []string) {
	if level == len(rule) {
		fmt.Fprintln(out, strings.Join(tokens, ""))
		return
	}
	switch rule[level] {
	case '#':
		for i := range dict {
			dfs(level+1, append(tokens, dict[i]))
		}
	default:
		for i := '0'; i <= '9'; i++ {
			dfs(level+1, append(tokens, string(i)))
		}
	}
}

func main() {
	in, _ := os.Open("628.in")
	defer in.Close()
	out, _ = os.Create("628.out")
	defer out.Close()

	var nd, nr int
	for {
		if _, err := fmt.Fscanf(in, "%d", &nd); err != nil {
			break
		}
		dict = make([]string, nd)
		for i := range dict {
			fmt.Fscanf(in, "%s", &dict[i])
		}
		fmt.Fscanf(in, "%d", &nr)
		fmt.Fprintln(out, "--")
		for nr > 0 {
			fmt.Fscanf(in, "%s", &rule)
			dfs(0, nil)
			nr--
		}
	}
}
