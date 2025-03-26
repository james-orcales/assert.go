//go:build !removeasserts

/*
Package assert provides condition evaluation utilities that terminate the program using os.Exit(1),
preventing panic recovery, along with Zig-style stacktrace output. Assertions can be disabled with
the `removeasserts` build tag. Dot import is the intended import style for this package.
*/

package assert

import (
	"errors"
	"fmt"
	"os"
)

// Unreachable crashes with an exit code of 1 and writes a stacktrace.
func Unreachable() {
	PanicAlways("reached unreachable code")
}

// Signal that a condition is sometimes true or false. This will never crash the program.
func Maybe(ok bool) {
	Assert(ok || !ok)
}

// Assert crashes if cond is false. If you need Assert(item == nil), use [Nil](item) instead.
// Do not defer assertions. There's no way to get the original line number of a deferred
// function leading to confusing stacktraces.
func Assert(cond bool) {
	if !cond {
		Unreachable() // assertion failure
	}
}

// AssertNil crashes if x is NOT nil and prints the non-null object.
// Prefer this over [Assert](x == nil) for readability.
func AssertNil(x any) {
	if x != nil {
		fmt.Fprintf(os.Stderr, "%v\n", x)
		Unreachable() // assertion failure
	}
}

// AssertErrIs crashes if actual is NOT one of the specified targets.
// Must provide at least one target. All targets must not be nil.
func AssertErrIs(actual error, targets ...error) {
	Assert(len(targets) > 0)

	for _, t := range targets {
		Assert(t != nil)
		if errors.Is(actual, t) {
			return
		}
	}
	fmt.Fprintf(os.Stderr, "%v\n", actual)
	Unreachable() // assertion failure
}

// AssertErrIsNot crashes if actual is one of the specified targets.
// Must provide at least one target. All targets must not be nil.
func AssertErrIsNot(actual error, targets ...error) {
	Assert(len(targets) > 0)

	for _, t := range targets {
		Assert(t != nil)
		if errors.Is(actual, t) {
			fmt.Fprintf(os.Stderr, "%v\n", actual)
			Unreachable() // assertion failure
		}
	}
}

/*
XAssert evaluates fn and crashes if it returns false.
It is designed for use cases where you want to perform expensive validations that can be disabled
in production builds using the `removeasserts` build tag.

	expensiveFn := func() bool { ... }
	// expensiveFn is still evaluated but boolean check is a noop under removeasserts
	Assert(expensiveFn())
	// expensiveFn itself will be a noop under removeasserts
	XAssert(expensiveFn)
*/
func XAssert(fn func() bool) {
	if !fn() {
		Unreachable() // assertion failure
	}
}

func XAssertNil(fn func() any) {
	x := fn()
	if x != nil {
		fmt.Fprintf(os.Stderr, "%v\n", x)
		Unreachable() // assertion failure
	}
}

func XAssertErrIs(fn func() error, targets ...error) {
	Assert(len(targets) > 0)

	actual := fn()
	for _, t := range targets {
		if errors.Is(actual, t) {
			return
		}
	}
	fmt.Fprintf(os.Stderr, "%v\n", actual)
	Unreachable() // assertion failure
}

func XAssertErrIsNot(fn func() error, targets ...error) {
	Assert(len(targets) > 0)

	actual := fn()
	for _, t := range targets {
		if errors.Is(actual, t) {
			fmt.Fprintf(os.Stderr, "%v\n", actual)
			Unreachable() // assertion failure
		}
	}
}
