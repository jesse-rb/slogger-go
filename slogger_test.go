package logger

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
    var out bytes.Buffer

    // Init new logger
    var localPrefix string = "slogger_test"
    var flags int = 0
    var l = New(&out, ANSIGreen, localPrefix, flags)
    
    // Log to output
    var tag string = "testNew"
    var msg string = "Testing a new logger"
    var data []int = []int{3, 5, 7}
    l.Log(tag, msg, data)

    // Expected string
    var expected string = fmt.Sprintf("%s%s\n", formatPrefix(ANSIGreen, localPrefix), formatLog(tag, msg, data))

    // Test that output is equal to expected
    var outString = out.String()
    t.Log(outString)
    if (outString != expected) {
        t.Fatalf("out:\n%s\ndoes not match:\n%s", outString, expected)
    }
}

