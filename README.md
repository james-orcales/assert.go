```go
package main

import (
	"errors"

	"github.com/amesaine/assert.go"
)

func main() {
	assert.Assert('a' < 'b')
	foo()
}

func foo() {
	bar()
}

func bar() {
	err := baz()
	assert.Nil(err)
}

func baz() error {
	return errors.New("where am i?")
}

/*
where am i?
/home/amesaine/personal/git/scratch/go/main.go:20
	assert.Nil(err)

/home/amesaine/personal/git/scratch/go/main.go:15
	bar()

/home/amesaine/personal/git/scratch/go/main.go:11
	foo()

/usr/local/go/src/runtime/proc.go:272
	fn()

exit status 1
*/
```
