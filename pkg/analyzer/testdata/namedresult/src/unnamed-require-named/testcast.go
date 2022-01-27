package testcase

// This file is a copy of unnamed-require-unnamed/testcase.go with additional "want" comments

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

func aStr() string { return "a" } // want "should use named function result parameter"

func bStr() string { // want "should use named function result parameter"
	return "b"
}

var cStr = func() string { // want "should use named function result parameter"
	return "c"
}

func multiStrErr() (string, error) { // want "should use named function result parameter"
	return "", nil
}

func multiStrErrMultiline() (
	string, // want "should use named function result parameter"
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

	maybeSay := func(what string) bool { // want "should use named function result parameter"
		return false
	}

	maybeSay("maybe hello")

	var (
		maybeSay2 = func(what string) bool { // want "should use named function result parameter"
			return false
		}
	)

	maybeSay2("maybe hello")

	defer func() {
		deferFn := func() error { return nil } // want "should use named function result parameter"
		deferFn()
	}()
}

type myGlobalFunc func(arg1, arg2 interface{}) (int, error) // want "should use named function result parameter"

func randomname2() {
	type myAnonFunc func(arg1, arg2 interface{}) (int, error) // want "should use named function result parameter"

	myAnonFuncImpl := func(arg1, arg2 interface{}) (int, error) { // want "should use named function result parameter"
		return 0, nil
	}

	_, _ = myAnonFuncImpl(3.1415, 42)
}

type obj struct{}

func (o *obj) hello1() error { return nil } // want "should use named function result parameter"

func hello2() error { return nil } // want "should use named function result parameter"
