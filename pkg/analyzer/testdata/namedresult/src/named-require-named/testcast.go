package testcase

// This file is a copy of named-require-unnamed/testcase.go without any "want" comments

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

func aStr() (msg string) { return "a" }

func bStr() (msg string) {
	msg = "b"
	return
}

var cStr = func() (msg string) {
	msg = "c"
	return msg
}

func multiStrErr() (msg string, err error) {
	msg = ""
	err = nil
	return
}

func multiStrErrMultiline() (
	msg string,
	err error,
	num interface{},
) {
	msg = ""
	err = nil
	num = 42
	return msg, err, num
}

func randomname1() {
	say := func(what string) {
		fmt.Println(what)
	}

	say("hello")

	maybeSay := func(what string) (shouldSay bool) {
		return false
	}

	maybeSay("maybe hello")

	var (
		maybeSay2 = func(what string) (shouldSay bool) {
			return false
		}
	)

	maybeSay2("maybe hello")

	defer func() {
		deferFn := func() (err error) { return nil }
		deferFn()
	}()
}

type myGlobalFunc func(arg1, arg2 interface{}) (num int, err error)

func randomname2() {
	type myAnonFunc func(arg1, arg2 interface{}) (num int, err error)

	myAnonFuncImpl := func(arg1, arg2 interface{}) (num int, err error) {
		return 0, nil
	}

	_, _ = myAnonFuncImpl(3.1415, 42)
}

type obj struct{}

func (o *obj) hello1() (err error) { return nil }

func hello2() (err error) { return nil }
