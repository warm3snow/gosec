package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
)

type useGoroutine struct {
	gosec.MetaData
}

func (u *useGoroutine) ID() string {
	return u.MetaData.ID
}

func (u *useGoroutine) Match(node ast.Node, ctx *gosec.Context) (*gosec.Issue, error) {
	switch n := node.(type) {
	case *ast.GoStmt:
		return gosec.NewIssue(ctx, n, u.ID(), u.What, u.Severity, u.Confidence), nil
	}

	return nil, nil
}

// NewUsesGoroutine detects the use of goroutines
func NewUsesGoroutine(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	return &useGoroutine{
		MetaData: gosec.MetaData{
			ID:         id,
			Severity:   gosec.High,
			Confidence: gosec.High,
			What:       "Use of goroutine",
		},
	}, []ast.Node{(*ast.GoStmt)(nil)}
}
