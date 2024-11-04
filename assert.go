//go:build !removeasserts

/*
Package assert provides condition evaluation utilities that terminate the program
using os.Exit(1), preventing panic recovery, along with Zig-style stacktrace output.
Assertions can be disabled with the `removeasserts` build tag.
*/

package assert

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Unreachable crashes the program when evaluated with an exit code of 1 and writes a stacktrace.
func Unreachable() {
	callers := make([]uintptr, 50)

	// Stacktrace starts at the caller of this function
	const callerOfThisFunc = 2

	count := runtime.Callers(callerOfThisFunc, callers)
	frames := runtime.CallersFrames(callers[0:count])

	for {
		frame, ok := frames.Next()
		if !ok {
			break
		}

		fn := frame.Function
		if frame.File != "" && frame.Line > 0 {
			f, err := os.Open(frame.File)
			if err == nil {
				defer f.Close()

				sc := bufio.NewScanner(f)

				for range frame.Line {
					_ = sc.Scan()
				}
				fn = strings.TrimSpace(sc.Text())
			}
		}

		fmt.Printf(
			"%v:%v\n\t%v\n\n",
			frame.File,
			frame.Line,
			fn,
		)
	}
	os.Exit(1)
}

// Signal that a condition is sometimes true or false. This will never crash the program.
func Maybe(ok bool) {
	Assert(ok || !ok)
}

// Signal that a scope is unimplemented and crash the progam.
func Unimplemented(msg string) {
	fmt.Printf("unimplemented: %s", msg)
	Unreachable()
}

// Assert crashes if cond is false. If you need Assert(item == nil), use [Nil](item) instead.
func Assert(cond bool) {
	if !cond {
		Unreachable() // assertion failure
	}
}

// Nil crashes if x is NOT nil. Prefer this over [Assert](x == nil) for readability.
func Nil(x any) {
	if x != nil {
		fmt.Printf("%v\n", x)
		Unreachable() // assertion failure
	}
}

// ErrIs crashes if actual is NOT one of the specified targets.
// Must provide at least one target. All targets must not be nil.
func ErrIs(actual error, targets ...error) {
	Assert(len(targets) > 0)

	confirmed := false
	for _, t := range targets {
		Assert(t != nil)
		if errors.Is(actual, t) && !confirmed {
			confirmed = true
		}
	}
	fmt.Printf("%v\n", actual)
	Unreachable() // assertion failure
}

// ErrIsNot crashes if actual is one of the specified targets.
// Must provide at least one target. All targets must not be nil.
func ErrIsNot(actual error, targets ...error) {
	Assert(len(targets) > 0)

	for _, t := range targets {
		if errors.Is(actual, t) {
			fmt.Printf("%v\n", actual)
			Unreachable() // assertion failure
		}
	}
}

/*
XAssert evaluates fn and crashes if it returns false.
It is designed for use cases where you want to perform expensive validations that can be disabled
in production builds using the `removeasserts` build tag.
*/
func XAssert(fn func() bool) {
	if !fn() {
		Unreachable() // assertion failure
	}
}

func XNil(fn func() any) {
	x := fn()
	if x != nil {
		fmt.Printf("%v\n", x)
		Unreachable() // assertion failure
	}
}

func XErrIs(fn func() error, targets ...error) {
	Assert(len(targets) > 0)

	actual := fn()
	for _, t := range targets {
		if errors.Is(actual, t) {
			return
		}
	}
	fmt.Printf("%v\n", actual)
	Unreachable() // assertion failure
}

func XErrIsNot(fn func() error, targets ...error) {
	Assert(len(targets) > 0)

	actual := fn()
	for _, t := range targets {
		if errors.Is(actual, t) {
			fmt.Printf("%v\n", actual)
			Unreachable() // assertion failure
		}
	}
}
