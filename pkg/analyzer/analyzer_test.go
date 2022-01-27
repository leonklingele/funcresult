package analyzer_test

import (
	"flag"
	"path/filepath"
	"testing"

	"github.com/leonklingele/funcresult/pkg/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

// TODO(leon): Add fuzzing

func TestNamedResult(t *testing.T) {
	t.Parallel()

	fixtures := []struct {
		name  string
		flags flag.FlagSet
	}{
		{
			name: "unnamed-require-named",
			flags: flags().
				withRequireNamed().
				build(),
		},
		{
			name: "unnamed-require-unnamed",
			flags: flags().
				withRequireUnnamed().
				build(),
		},

		{
			name: "named-require-named",
			flags: flags().
				withRequireNamed().
				build(),
		},
		{
			name: "named-require-unnamed",
			flags: flags().
				withRequireUnnamed().
				build(),
		},
	}

	for _, f := range fixtures {
		f := f

		t.Run(f.name, func(t *testing.T) {
			t.Parallel()

			a := analyzer.New()
			a.Flags = f.flags

			testdata := filepath.Join(analysistest.TestData(), "namedresult")
			_ = analysistest.Run(t, testdata, a, f.name)
		})
	}
}

type flagger struct {
	fs *flag.FlagSet
}

func (f *flagger) withRequireNamed() *flagger {
	if err := f.fs.Lookup(analyzer.FlagNameRequireNamed).Value.Set("true"); err != nil {
		panic(err)
	}

	return f
}

func (f *flagger) withRequireUnnamed() *flagger {
	if err := f.fs.Lookup(analyzer.FlagNameRequireUnnamed).Value.Set("true"); err != nil {
		panic(err)
	}

	return f
}

func (f *flagger) build() flag.FlagSet {
	return *f.fs
}

func flags() *flagger {
	fs := analyzer.Flags()

	return &flagger{
		fs: &fs,
	}
}
