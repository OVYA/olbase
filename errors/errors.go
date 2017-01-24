package errors

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

// Error interface exposes additional information about the error.
type Error interface {
	// This returns the error message without the stack trace.
	GetMessage() string

	// This returns the wrapped error.  This returns nil if this does not wrap
	// another error.
	GetInner() error

	// Implements the built-in error interface.
	Error() string

	// Returns stack addresses as a string that can be supplied to
	// a helper tool to get the actual stack trace. This function doesn't result
	// in resolving full stack frames thus is a lot more efficient.
	StackAddrs() string

	// Returns stack frames.
	StackFrames() []StackFrame

	// Returns string representation of stack frames.
	// Stack frame formatting looks generally something like this:
	// dropbox/rpc.(*clientV4).Do
	//   /srv/server/go/src/dropbox/rpc/client.go:87 +0xbf9
	// dropbox/exclog.Report
	//   /srv/server/go/src/dropbox/exclog/client.go:129 +0x9e5
	// main.main
	//   /home/cdo/tmp/report_exception.go:13 +0x84
	// It is discouraged to parse stack frames using string parsing since it can change at any time.
	// Use StackFrames() function instead to get actual stack frame metadata.
	GetStack() string
}

// StackFrame Represents a single stack frame.
type StackFrame struct {
	PC         uintptr
	Func       *runtime.Func
	FuncName   string
	File       string
	LineNumber int
}

// Standard struct for general types of errors.
//
// For an example of custom error type, look at databaseError/newDatabaseError
// in errors_test.go.
type baseError struct {
	msg   string
	inner error

	stack       []uintptr
	framesOnce  sync.Once
	stackFrames []StackFrame
}

func getInternalMsgFromError(err error) string {
	return "\n Partial error msg...\nInternal lib error is " + err.Error()
}

// GetMessage returns the error string without stack trace information.
func GetMessage(err interface{}) string {
	switch e := err.(type) {
	case Error:
		errMsg, err := extractFullErrorMessage(e, false)

		if err != nil {
			errMsg += getInternalMsgFromError(err)
		}

		return errMsg
	case runtime.Error:
		return runtime.Error(e).Error()
	case error:
		return e.Error()
	default:
		return "Passed a non-error to GetMessage"
	}
}

// Error returns a string with all available error information, including inner
// errors that are wrapped by this errors.
func (e *baseError) Error() string {
	errMsg, err := extractFullErrorMessage(e, true)
	if err != nil {
		errMsg += getInternalMsgFromError(err)
	}

	return errMsg
}

// GetMessage implements Error interface.
func (e *baseError) GetMessage() string {
	return e.msg
}

// GetInner implements Error interface.
func (e *baseError) GetInner() error {
	return e.inner
}

// StackAddrs implements Error interface.
func (e *baseError) StackAddrs() string {
	buf := bytes.NewBuffer(make([]byte, 0, len(e.stack)*8))
	for _, pc := range e.stack {
		fmt.Fprintf(buf, "0x%x ", pc)
	}
	bufBytes := buf.Bytes()
	return string(bufBytes[:len(bufBytes)-1])
}

// StackFrames implements Error interface.
func (e *baseError) StackFrames() []StackFrame {
	e.framesOnce.Do(func() {
		e.stackFrames = make([]StackFrame, len(e.stack))
		for i, pc := range e.stack {
			frame := &e.stackFrames[i]
			frame.PC = pc
			frame.Func = runtime.FuncForPC(pc)
			if frame.Func != nil {
				frame.FuncName = frame.Func.Name()
				frame.File, frame.LineNumber = frame.Func.FileLine(frame.PC - 1)
			}
		}
	})
	return e.stackFrames
}

// GetStack implements Error interface.
func (e *baseError) GetStack() string {
	stackFrames := e.StackFrames()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	var err error

	for _, frame := range stackFrames {
		_, err = buf.WriteString(frame.FuncName)
		if err != nil {
			fmt.Fprintf(buf, getInternalMsgFromError(err)+"\n")
		}
		_, err = buf.WriteString("\n")
		if err != nil {
			fmt.Fprintf(buf, getInternalMsgFromError(err)+"\n")
		}

		fmt.Fprintf(buf, "\t%s:%d +0x%x\n",
			frame.File, frame.LineNumber, frame.PC)
	}

	return buf.String()
}

// New returns a new baseError initialized with the given message and
// the current stack trace.
func New(msg string) Error {
	return new(nil, msg)
}

// Newf is the same as New, but with fmt.Printf-style parameters.
func Newf(format string, args ...interface{}) Error {
	return new(nil, fmt.Sprintf(format, args...))
}

// Wrap wraps another error in a new baseError.
func Wrap(err error, msg string) Error {
	return new(err, msg)
}

// Wrapf is the same as Wrap, but with fmt.Printf-style parameters.
func Wrapf(err error, format string, args ...interface{}) Error {
	return new(err, fmt.Sprintf(format, args...))
}

// Internal helper function to create new baseError objects,
// note that if there is more than one level of redirection to call this function,
// stack frame information will include that level too.
func new(err error, msg string) *baseError {
	stack := make([]uintptr, 200)
	stackLength := runtime.Callers(3, stack)
	return &baseError{
		msg:   msg,
		stack: stack[:stackLength],
		inner: err,
	}
}

// Constructs full error message for a given Error by traversing
// all of its inner errors. If includeStack is True it will also include
// stack trace from deepest Error in the chain.
func extractFullErrorMessage(e Error, includeStack bool) (string, error) {
	var ok bool
	var lastDbxErr Error
	errMsg := bytes.NewBuffer(make([]byte, 0, 1024))
	var _ int
	var err error

	dbxErr := e
	for {
		lastDbxErr = dbxErr
		_, err = errMsg.WriteString(dbxErr.GetMessage())
		if err != nil {
			goto returnError
		}

		innerErr := dbxErr.GetInner()
		if innerErr == nil {
			break
		}

		dbxErr, ok = innerErr.(Error)
		if !ok {
			// We have reached the end and traveresed all inner errors.
			// Add last message and exit loop.
			_, err = errMsg.WriteString(innerErr.Error())

			if err != nil {
				goto returnError
			}

			break
		}

		_, err = errMsg.WriteString("\n")

		if err != nil {
			goto returnError
		}
	}

	if includeStack {
		_, err = errMsg.WriteString("\nORIGINAL STACK TRACE:\n")
		if err != nil {
			goto returnError
		}

		_, err = errMsg.WriteString(lastDbxErr.GetStack())
		if err != nil {
			goto returnError
		}
	}

	return errMsg.String(), nil

returnError:
	return errMsg.String(), err
}

// unwrapError returns a wrapped error or nil if there is none.
func unwrapError(ierr error) (nerr error) {
	// Internal errors have a well defined bit of context.
	if dbxErr, ok := ierr.(Error); ok {
		return dbxErr.GetInner()
	}

	// At this point, if anything goes wrong, just return nil.
	defer func() {
		if x := recover(); x != nil {
			nerr = nil
		}
	}()

	// Go system errors have a convention but paradoxically no
	// interface.  All of these panic on error.
	errV := reflect.ValueOf(ierr).Elem()
	errV = errV.FieldByName("Err")
	return errV.Interface().(error)
}

// RootError keep peeling away layers or context until a primitive
// error is revealed.
func RootError(ierr error) (nerr error) {
	nerr = ierr
	for i := 0; i < 20; i++ {
		terr := unwrapError(nerr)
		if terr == nil {
			return nerr
		}
		nerr = terr
	}
	return fmt.Errorf("too many iterations: %T", nerr)
}

// IsError performs a deep check, unwrapping errors as much as
// possilbe and
// comparing the string version of the error.
func IsError(err, errConst error) bool {
	if err == errConst {
		return true
	}
	// Must rely on string equivalence, otherwise a value is not equal
	// to its pointer value.
	rootErrStr := ""
	rootErr := RootError(err)
	if rootErr != nil {
		rootErrStr = rootErr.Error()
	}
	errConstStr := ""
	if errConst != nil {
		errConstStr = errConst.Error()
	}
	return rootErrStr == errConstStr
}
