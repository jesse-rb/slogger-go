package logger

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
    var color string = ANSIRed;
    var localPrefix string = "slogger_test"
    var flags int = 0
    // Init new logger
    var l interface{} = New(os.Stdout, color, localPrefix, flags)

    if l, ok := l.(*Logger); ok {
        t.Log("l IS of type *Logger")

        if l.localPrefix != localPrefix+"xd" {
            t.Errorf("localPrefix: %s\ndoes not match expected: %s\n", l.localPrefix, localPrefix)
        }
    } else {
        t.Fatalf("l is NOT of type *Logger")
    }
}

func TestLog(t *testing.T) {
    // Init new logger
    var out bytes.Buffer
    var l = New(&out, ANSIGreen, "slogger_test", 0)

    // Log to output
    var tag string = "testNew"
    var msg string = "Testing a new logger"
    var data []int = []int{3, 5, 7}
    l.Log(tag, msg, data)

    // Expected string
    var expected string = fmt.Sprintf("%s%s\n", formatPrefix(l.color, l.localPrefix), formatLog(tag, msg, data))

    // Test that output is equal to expected
    var outString = out.String()
    t.Log(outString)
    if (outString != expected) {
        t.Fatalf("out:\n%s\ndoes not match expected:\n%s", outString, expected)
    }
}

