package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
)

type useGlobalVars struct {
	gosec.MetaData
}

func (u *useGlobalVars) ID() string {
	return u.MetaData.ID
}

func (u *useGlobalVars) Match(node ast.Node, ctx *gosec.Context) (*gosec.Issue, error) {
	globalVarNames := getGlobalVarNames(ctx)
	switch n := node.(type) {
	case *ast.ValueSpec:
		spec, _ := node.(*ast.ValueSpec)
		for _, name := range spec.Names {
			if _, ok := globalVarNames[name.Name]; ok {
				return gosec.NewIssue(ctx, n, u.ID(), u.What, u.Severity, u.Confidence), nil
			}
		}
	}

	return nil, nil
}

func getGlobalVarNames(ctx *gosec.Context) map[string]struct{} {
	varNamesSet := make(map[string]struct{})
	if ctx.Root != nil && ctx.Root.Decls != nil {
		for _, decl := range ctx.Root.Decls {
			if genDecl, ok1 := decl.(*ast.GenDecl); ok1 {
				for _, spec := range genDecl.Specs {
					if valueSpec, ok2 := spec.(*ast.ValueSpec); ok2 {
						for _, name := range valueSpec.Names {
							varNamesSet[name.Name] = struct{}{}
						}
					}
				}
			}
		}
	}
	return varNamesSet
}

// NewUsesGoroutine detects the use of goroutines
func NewUsesGlobalVars(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	return &useGlobalVars{
		MetaData: gosec.MetaData{
			ID:         id,
			Severity:   gosec.High,
			Confidence: gosec.Medium,
			What:       "Use of global variables",
		},
	}, []ast.Node{(*ast.ValueSpec)(nil), (*ast.GenDecl)(nil)}
}
