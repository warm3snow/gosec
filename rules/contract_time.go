package rules

import (
	"fmt"
	"github.com/securego/gosec/v2"
	"go/ast"
)

type useSystemTime struct {
	gosec.MetaData
	timeCalls gosec.CallList
}

func (u *useSystemTime) ID() string {
	return u.MetaData.ID
}

func (u *useSystemTime) Match(node ast.Node, ctx *gosec.Context) (*gosec.Issue, error) {
	aliasMap := GetAliasMap(ctx)
	switch n := node.(type) {
	case *ast.CallExpr:
		if u.timeCalls.ContainsPkgCallExpr(n, ctx, false) != nil {
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

// NewUsesGoroutine detects the use of goroutines
func NewUsesSystemTime(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	timeCalls := gosec.NewCallList()
	timeCalls.AddAll("time", "Date", "Now", "Since", "Unix", "UnixMicro", "UnixMilli")

	return &useSystemTime{
		MetaData: gosec.MetaData{
			ID:         id,
			Severity:   gosec.High,
			Confidence: gosec.Medium,
			What:       "Use of system time",
		},
		timeCalls: timeCalls,
	}, []ast.Node{(*ast.CallExpr)(nil)}
}
