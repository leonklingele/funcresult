package testcase

// This file is a copy of named-require-named/testcase.go with additional "want" comments

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

func aStr() (msg string) { return "a" } // want "should use unnamed function result parameter"

func bStr() (msg string) { // want "should use unnamed function result parameter"
	msg = "b"
	return
}

var cStr = func() (msg string) { // want "should use unnamed function result parameter"
	msg = "c"
	return msg
}

func multiStrErr() (msg string, err error) { // want "should use unnamed function result parameter"
	msg = ""
	err = nil
	return
}

func multiStrErrMultiline() (
	msg string, // want "should use unnamed function result parameter"
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

	maybeSay := func(what string) (shouldSay bool) { // want "should use unnamed function result parameter"
		return false
	}

	maybeSay("maybe hello")

	var (
		maybeSay2 = func(what string) (shouldSay bool) { // want "should use unnamed function result parameter"
			return false
		}
	)

	maybeSay2("maybe hello")

	defer func() {
		deferFn := func() (err error) { return nil } // want "should use unnamed function result parameter"
		deferFn()
	}()
}

type myGlobalFunc func(arg1, arg2 interface{}) (num int, err error) // want "should use unnamed function result parameter"

func randomname2() {
	type myAnonFunc func(arg1, arg2 interface{}) (num int, err error) // want "should use unnamed function result parameter"

	myAnonFuncImpl := func(arg1, arg2 interface{}) (num int, err error) { // want "should use unnamed function result parameter"
		return 0, nil
	}

	_, _ = myAnonFuncImpl(3.1415, 42)
}

type obj struct{}

func (o *obj) hello1() (err error) { return nil } // want "should use unnamed function result parameter"

func hello2() (err error) { return nil } // want "should use unnamed function result parameter"
