package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
)

type usesForMap struct {
	gosec.MetaData
}

func (u *usesForMap) ID() string {
	return u.MetaData.ID
}

func (u *usesForMap) Match(node ast.Node, ctx *gosec.Context) (*gosec.Issue, error) {
	//switch n := node.(type) {
	//case *ast.RangeStmt:
	//	if expr, ok := node.(*ast.RangeStmt); ok {
	//		fmt.Println(expr)
	//	}
	//	//expr.Fun.(*ast.SelectorExpr).X.(*ast.Ident).Name   //syncMap => sync.map
	//	//expr.Fun.(*ast.SelectorExpr).Sel.(*ast.Ident).Name //Range
	//case *ast.GenDecl:
	//	if decl, ok := node.(*ast.GenDecl); ok {
	//		fmt.Println(decl)
	//	}
	//
	//case *ast.ForStmt:
	//	return gosec.NewIssue(ctx, n, u.ID(), u.What, u.Severity, u.Confidence), nil
	//}

	return nil, nil
}

// NewUsesForMap detects the use of for range map
func NewUsesForMap(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	return &usesForMap{
		MetaData: gosec.MetaData{
			ID:         id,
			Severity:   gosec.High,
			Confidence: gosec.Medium,
			What:       "Use of for-range-map",
		},
	}, []ast.Node{(*ast.RangeStmt)(nil), (*ast.ForStmt)(nil)}
}
