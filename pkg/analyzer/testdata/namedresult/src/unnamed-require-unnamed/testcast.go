package testcase

// This file is a copy of unnamed-require-named/testcase.go without any "want" comments

import (
	"fmt"
)

func a() { fmt.Println(aStr()) }

func b() {
	fmt.Println(bStr())
}

var c = func() {
	fmt.Println(cStr())
}

var (
	d = func() {
		fmt.Println("d")
	}
)

func aStr() string { return "a" }

func bStr() string {
	return "b"
}

var cStr = func() string {
	return "c"
}

func multiStrErr() (string, error) {
	return "", nil
}

func multiStrErrMultiline() (
	string,
	error,
	interface{},
) {
	return "", nil, 42
}

func randomname1() {
	say := func(what string) {
		fmt.Println(what)
	}

	say("hello")

	maybeSay := func(what string) bool {
		return false
	}

	maybeSay("maybe hello")

	var (
		maybeSay2 = func(what string) bool {
			return false
		}
	)

	maybeSay2("maybe hello")

	defer func() {
		deferFn := func() error { return nil }
		deferFn()
	}()
}

type myGlobalFunc func(arg1, arg2 interface{}) (int, error)

func randomname2() {
	type myAnonFunc func(arg1, arg2 interface{}) (int, error)

	myAnonFuncImpl := func(arg1, arg2 interface{}) (int, error) {
		return 0, nil
	}

	_, _ = myAnonFuncImpl(3.1415, 42)
}

type obj struct{}

func (o *obj) hello1() error { return nil }

func hello2() error { return nil }
