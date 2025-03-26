package assert

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// PanicAlways cannot be recovered from, unlike the builtin panic().
// PanicAlways crashes regardless of build tag removeasserts.
func PanicAlways(msg string) {
	fmt.Fprintln(os.Stderr, "panic: "+msg+"\n")
	fatalStackTrace()
}

// Signal that a scope is unimplemented and crash the progam.
// Unimplemented crashes regardless of build tag removeasserts.
func Unimplemented() {
	PanicAlways("reached unimplemented code")
}

func fatalStackTrace() {
	callers := make([]uintptr, 50)

	// Stacktrace starts at the caller of this function
	const fatalStackTraceCallerFrame = 2

	count := runtime.Callers(fatalStackTraceCallerFrame, callers)
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
		fmt.Fprintf(
			os.Stderr,
			"%v:%v\n\t%v\n\n",
			frame.File,
			frame.Line,
			fn,
		)
	}
	os.Exit(1)
}
