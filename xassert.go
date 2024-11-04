//go:build removeasserts

package assert

func Assert(_ bool)                        {}
func Nil(_ any)                            {}
func ErrIs(_ error, _ ...error)            {}
func ErrIsNot(_ error, _ ...error)         {}
func XAssert(_ func() bool)                {}
func XNil(_ func() any)                    {}
func XErrIs(_ func() error, _ ...error)    {}
func XErrIsNot(_ func() error, _ ...error) {}
func Unreachable()                               {}
