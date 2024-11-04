```go
package main

import (
	"errors"

	. "github.com/amesaine/assert.go"
)

func main() {
	Assert('a' < 'b')
	foo()
}

func foo() {
	bar()
}

func bar() {
	err := baz()
	AssertNil(err)
}

func baz() error {
	return errors.New("where am i?")
}

/*
where am i?
/home/amesaine/personal/git/scratch/go/main.go:20
	AssertNil(err)

/home/amesaine/personal/git/scratch/go/main.go:15
	bar()

/home/amesaine/personal/git/scratch/go/main.go:11
	foo()

/usr/local/go/src/runtime/proc.go:272
	fn()

exit status 1
*/
```
