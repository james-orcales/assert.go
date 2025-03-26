//go:build removeasserts && !customasserts

package assert

func Unreachable()                               {}
func Maybe()                                     {}
func Unimplemented(_ string)                     {}
func Assert(_ bool)                              {}
func AssertNil(_ any)                            {}
func AssertErrIs(_ error, _ ...error)            {}
func AssertErrIsNot(_ error, _ ...error)         {}
func XAssert(_ func() bool)                      {}
func XAssertNil(_ func() any)                    {}
func XAssertErrIs(_ func() error, _ ...error)    {}
func XAssertErrIsNot(_ func() error, _ ...error) {}
