package analyzer

import (
	"flag"
)

const (
	FlagNameRequireNamed   = "require-named"
	FlagNameRequireUnnamed = "require-unnamed"
)

func Flags() flag.FlagSet {
	fs := flag.NewFlagSet(Name, flag.ExitOnError)

	fs.Bool(FlagNameRequireNamed, false, "require the use of named function result parameters only")
	fs.Bool(FlagNameRequireUnnamed, false, "require the use of unnamed function result parameters only")

	return *fs
}
