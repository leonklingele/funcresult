package namedresult

import (
	"errors"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// https://go.dev/ref/spec#Function_types

type Func struct {
	Type *ast.FuncType
}

type funcVisitor struct {
	funcs []*Func
}

func (fv *funcVisitor) Visit(n ast.Node) ast.Visitor {
	switch t := n.(type) {
	case *ast.FuncType:
		fv.funcs = append(fv.funcs, &Func{
			Type: t,
		})
	}

	return fv
}

func (fv *funcVisitor) Funcs() []*Func {
	return fv.funcs
}

func Filepass(c *Config, p *analysis.Pass, f *ast.File) error {
	if c.RequireNamed && c.RequireUnnamed {
		// Can't use both at the same time
		return errors.New("configuration conflict: can't require both named and unnamed function result parameters")
	}

	fv := &funcVisitor{}
	ast.Walk(fv, f)
	funcs := fv.Funcs()

	for _, f := range funcs {
		checkFunc(c, p, f)
	}

	return nil
}

func checkFunc(c *Config, p *analysis.Pass, f *Func) {
	results := f.Type.Results
	if results == nil {
		// Function doesn't have any results, skip it.
		return
	}

	list := results.List
	if len(list) == 0 {
		return
	}

	// Only need to check the first list item.
	// The Go compiler ensures either all or none parameters are named / unnamed.
	firstitem := list[0]

	if c.RequireNamed && firstitem.Names == nil {
		msg := "should use named function result parameter"

		report := analysis.Diagnostic{ //nolint:exhaustivestruct // we do not need all fields
			Pos:     firstitem.Pos(),
			End:     firstitem.End(),
			Message: msg,
			// TODO(leon): Suggest fix
		}

		if len(list) > 1 {
			report.Related = toRelated(list[1:])
		}

		p.Report(report)
	}

	if c.RequireUnnamed && firstitem.Names != nil {
		msg := "should use unnamed function result parameter"

		report := analysis.Diagnostic{ //nolint:exhaustivestruct // we do not need all fields
			Pos:     firstitem.Pos(),
			End:     firstitem.End(),
			Message: msg,
			// TODO(leon): Suggest fix
		}

		if len(list) > 1 {
			report.Related = toRelated(list[1:])
		}

		p.Report(report)
	}
}

func toRelated(list []*ast.Field) []analysis.RelatedInformation {
	related := make([]analysis.RelatedInformation, 0, len(list))
	for _, item := range list {
		related = append(related, analysis.RelatedInformation{
			Pos:     item.Pos(),
			End:     item.End(),
			Message: "found here",
		})
	}

	return related
}
