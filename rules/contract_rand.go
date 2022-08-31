package rules

import (
	"fmt"
	"go/ast"

	"github.com/securego/gosec/v2"
)

type useRand struct {
	gosec.MetaData
	randCalls gosec.CallList
}

func (u *useRand) ID() string {
	return u.MetaData.ID
}

//func containsRandCall(node ast.Node, ctx *gosec.Context, list gosec.CallList) bool {
//	if list.ContainsPkgCallExpr(node, ctx, false) != nil {
//		return true
//	}
//	return false
//}

func (u *useRand) Match(node ast.Node, ctx *gosec.Context) (*gosec.Issue, error) {
	aliasMap := GetAliasMap(ctx)
	switch n := node.(type) {
	case *ast.CallExpr:
		if u.randCalls.ContainsPkgCallExpr(n, ctx, false) != nil {
			selector, ident, err := gosec.GetCallInfo(n, ctx)
			if err != nil {
				return nil, nil
			}
			var what = u.What
			alias, ok := aliasMap[selector]
			if ok {
				what = fmt.Sprintf(u.What, alias)
			}
			_ = ident
			return gosec.NewIssue(ctx, n, u.ID(), what, u.Severity, u.Confidence), nil
		}
	}
	return nil, nil
}

func GetAliasMap(ctx *gosec.Context) map[string]string {
	aliasMap := make(map[string]string, len(ctx.Imports.Aliased))
	for k, v := range ctx.Imports.Aliased {
		aliasMap[v] = k
	}
	return aliasMap
}

// NewUsesRandFunc detects the use of random number generator
func NewUsesRandFunc(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	randCalls := gosec.NewCallList()
	randCalls.AddAll("math/rand", "New", "Read", "Float32", "Float64", "Int", "Int31",
		"Int31n", "Int63", "Int63n", "Intn", "NormalFloat64", "Uint32", "Uint64")

	randCalls.AddAll("crypto/rand", "Read", "Int", "Prime")
	return &useRand{
		MetaData: gosec.MetaData{
			ID:         id,
			Severity:   gosec.High,
			Confidence: gosec.High,
			What:       "Use of random number generator (%s)",
		},
		randCalls: randCalls,
	}, []ast.Node{(*ast.CallExpr)(nil)}
}
