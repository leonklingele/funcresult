package analyzer

import (
	"fmt"
	"go/ast"

	"github.com/leonklingele/funcresult/pkg/analyzer/namedresult"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const (
	Name = "funcresult"
	Doc  = `function result parameter analyzer: require named / unnamed function result parameters`
)

func New() *analysis.Analyzer {
	return &analysis.Analyzer{ //nolint:exhaustivestruct // we do not need all fields
		Name:     Name,
		Doc:      Doc,
		Flags:    Flags(),
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(p *analysis.Pass) (interface{}, error) {
	flagLookupBool := func(name string) bool {
		return p.Analyzer.Flags.Lookup(name).Value.String() == "true"
	}

	c := &Config{
		NamedResultConfig: &namedresult.Config{
			RequireNamed:   flagLookupBool(FlagNameRequireNamed),
			RequireUnnamed: flagLookupBool(FlagNameRequireUnnamed),
		},
	}

	return nil, pass(c, p)
}

func pass(c *Config, p *analysis.Pass) error {
	for _, f := range p.Files {
		if err := filepass(c, p, f); err != nil {
			return err
		}
	}

	return nil
}

func filepass(c *Config, p *analysis.Pass, f *ast.File) error {
	if err := namedresult.Filepass(c.NamedResultConfig, p, f); err != nil {
		return fmt.Errorf("failed to funcresult.Filepass: %w", err)
	}

	return nil
}
