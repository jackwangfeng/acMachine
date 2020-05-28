# acMachine
package main

import (
	"fmt"

	acMachine "github.com/jackwangfeng/acMachine"
)

func main() {
	m := acMachine.NewAcMachine()
	m.AddPatten("abc")
	m.AddPatten("cde")
	m.Build()
	results, pos := m.Match("abcdefabcdef")
	cLen := len(results)
	for i := 0; i < cLen; i++ {
		fmt.Printf("%d:%s\n", pos[i], results[i])
	}
}
