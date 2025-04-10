package logger

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	var color string = ANSIRed
	var localPrefix string = "slogger_test"
	var flags int = 0
	// Init new logger
	var l interface{} = New(os.Stdout, color, localPrefix, flags)

	if l, ok := l.(*Logger); ok {
		t.Log("l IS of type *Logger")

		if l.localPrefix != localPrefix {
			t.Errorf("localPrefix: %s\ndoes not match expected: %s\n", l.localPrefix, localPrefix)
		}
	} else {
		t.Fatalf("l is NOT of type *Logger")
	}
}

func TestLog(t *testing.T) {
	// Init new logger
	var out bytes.Buffer
	var localPrefix string = "slogger_test"
	var color string = ANSIGreen
	var flags int = 0
	l := New(&out, ANSIGreen, localPrefix, flags)

	// Log to output
	var tag string = "testNew"
	var msg string = "Testing a new logger"
	var data []int = []int{3, 5, 7}

	testLogOutput(t, l, &out, localPrefix, color, flags, tag, msg, data)
}

func TestSetLocalPrefix(t *testing.T) {
	// Init new logger
	var out bytes.Buffer
	var color string = ANSIGreen
	var localPrefix string = "slogger_test"
	var flags int = 0
	l := New(&out, color, localPrefix, flags)
	var newLocalPrefix string = "[TestLocalPrefix]"

	testLogOutput(t, l, &out, localPrefix, color, flags, "[test1]", "Will this pass?", (true && false || true))

	// Set new local prefix
	l.SetLocalPrefix(newLocalPrefix)

	// Test that new local prefix is saved
	if l.localPrefix != newLocalPrefix {
		t.Fatalf("localPrefix: %s\ndoes not match expected: %s\n", l.localPrefix, newLocalPrefix)
	}

	// Test that new local prefix is being logged
	testLogOutput(t, l, &out, newLocalPrefix, color, flags, "[test2]", "How about this?", true)
}

func TestSetColor(t *testing.T) {
	// Init new logger
	var out bytes.Buffer
	var color string = ANSIGreen
	var localPrefix string = "slogger_test"
	var flags int = 0
	l := New(&out, color, localPrefix, flags)
	var newColor string = ANSIYellow

	testLogOutput(t, l, &out, localPrefix, color, flags, "[before]", "Before updating color", 1)

	l.SetColor(newColor)

	if l.color != newColor {
		t.Fatalf("color: %s\ndoes not match expected: %s\n", l.color, newColor)
	}

	testLogOutput(t, l, &out, localPrefix, newColor, flags, "[after]", "After updating color", 2)
}

// General test log output function used in other tests
// Does require log output location to be a bytes buffer so that we can get the output in order to test it
func testLogOutput(t *testing.T, l *Logger, out *bytes.Buffer, localPrefix string, color string, flags int, tag string, msg string, data interface{}) {
	// Log to output
	l.Log(tag, msg, data)

	// Expected string
	var expected string = fmt.Sprintf("%s%s\n", formatPrefix(color, localPrefix), formatLog(tag, msg, data))

	// Test that output is equal to expected
	outString := out.String()
	t.Log(outString)
	if outString != expected {
		t.Fatalf("out:\n%s\ndoes not match expected:\n%s", outString, expected)
	}

	out.Reset()
}
