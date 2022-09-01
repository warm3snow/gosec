package _map

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

//map和sync map遍历都为无序
func TestMap(t *testing.T) {
	intMap := make(map[int]int)
	//for i := 0; i < 10; i++ {
	//	intMap[i] = i
	//}

	for _, v := range intMap {
		fmt.Printf("%d, ", v)
	}
}

//func TestSyncMap(t *testing.T) {
//	var syncMap sync.Map
//	//for i := 0; i < 10; i++ {
//	//	syncMap.Store(i, i)
//	//}
//
//	syncMap.Range(func(key, value any) bool {
//		//fmt.Printf("%d, ", value)
//		return true
//	})
//}

func TestGoAst(t *testing.T) {
	// src is the input for which we want to print the AST.
	src := `
package main
func main() {
	intMap := make(map[int]int)
	for _, v := range intMap {
		fmt.Printf("%d, ", v)
	}
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, _ := parser.ParseFile(fset, "", src, 0)

	// Print the AST.
	ast.Print(fset, f)
}
