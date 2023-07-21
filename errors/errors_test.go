package errors

import (
	"fmt"
	"reflect"
	"regexp"
	"syscall"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestStackTrace(t *testing.T) {
	const testMsg = "test error"
	er := New(testMsg)
	assert := assert.New(t)

	assert.Equal(er.GetMessage(), testMsg)
	assert.NotContains(er.GetStack(), "ovya.fr/olbase/errors/errors.go")
	assert.Contains(er.GetStack(), "testing.tRunner")

	for i, r := range er.GetStack() {
		if !(unicode.IsSpace(r) || unicode.IsPrint(r)) {
			t.Errorf("stack trace has an unexpected rune at index %v (%q)", i, r)
			break
		}
	}
}

func TestWrappedError(t *testing.T) {
	const (
		innerMsg  = "I am inner error"
		middleMsg = "I am the middle error"
		outerMsg  = "I am the mighty outer error"
	)

	inner := fmt.Errorf(innerMsg)
	middle := Wrap(inner, middleMsg)
	outer := Wrap(middle, outerMsg)
	errorStr := outer.Error()
	assert := assert.New(t)

	assert.Contains(errorStr, innerMsg)
	assert.Contains(errorStr, middleMsg)
	assert.Contains(errorStr, outerMsg)
}

func TestStackAddrs(t *testing.T) {
	pat := regexp.MustCompile("^0x[a-h0-9]+( 0x[a-h0-9]+)*$")
	er := New("big trouble")

	if !pat.MatchString(er.StackAddrs()) {
		t.Errorf("StackAddrs didn't match `%s`: %q", pat, er.StackAddrs())
	}
}

// ---------------------------------------
// minimal example + test for custom error
type databaseError struct {
	Error
	code int
}

// "constructor" for creating error needed to store return value of
// StackTrace() to get the code.
func newDatabaseError(msg string, code int) databaseError {
	return databaseError{Error: New(msg), code: code}
}

func TestCustomError(t *testing.T) {
	assert := assert.New(t)

	dbMsg := "database error 1205 (lock wait time exceeded)"
	outerMsg := "outer msg"

	dbError := newDatabaseError(dbMsg, 1205)
	outerError := Wrap(dbError.Error, outerMsg)

	errorStr := outerError.Error()

	assert.Contains(errorStr, dbMsg)
	assert.Contains(errorStr, outerMsg)
	assert.Contains(errorStr, "errors.newDatabaseError")
}

type customErr struct {
}

func (ce *customErr) Error() string { return "testing error" }

type customNestedErr struct {
	Err interface{}
}

func (cne *customNestedErr) Error() string { return "nested testing error" }

func TestRootError(t *testing.T) {
	assert := assert.New(t)
	err := RootError(nil)
	assert.Nil(err)

	var ce *customErr
	err = RootError(ce)
	assert.Equal(err, ce)

	ce = &customErr{}
	err = RootError(ce)
	assert.Equal(err, ce)

	cne := &customNestedErr{}
	err = RootError(cne)
	assert.Equal(err, cne)

	cne = &customNestedErr{reflect.ValueOf(ce).Pointer()}
	err = RootError(cne)
	assert.Equal(err, cne)

	cne = &customNestedErr{ce}
	err = RootError(cne)
	assert.Equal(err, ce)

	err = RootError(syscall.ECONNREFUSED)
	assert.Equal(err, syscall.ECONNREFUSED)
}

// Benchmarks creation of new errors.
// Current expected range is ~0.1-0.2ms to create errors from 100 go routines
// simultaneously. This is fairly close to just spinning up go routines
// and putting stuff on channels and doing some very simple work, thus
// error creation should be cheap enough for all most all use cases.
func BenchmarkNew(b *testing.B) {
	a := func() error {
		b := func() error {
			c := func() error {
				return New("Hello world, grab me a stack trace!")
			}
			return c()
		}
		return b()
	}
	nRoutines := 100
	errChan := make(chan error, nRoutines)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < nRoutines; k++ {
			go func() {
				err := a()
				errChan <- err
			}()
		}
		for k := 0; k < nRoutines; k++ {
			<-errChan
		}
	}
}
