package logger

import (
	"fmt"
	"io"
	"log"
)

// 3 bit color ansi color codes for colored logging output
// srouce: https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html#8-colors
const (
    ANSIReset    = "\u001b[0m"
    ANSIRed      = "\u001b[31m"
    ANSIBlue     = "\u001b[34m"
    ANSIBlack    = "\u001b[30m"
    ANSIGreen    = "\u001b[32m"
    ANSIYellow   = "\u001b[33m"
    ANSIMagenta  = "\u001b[35m"
    ANSICyan     = "\u001b[36m"
    ANSIWhite    = "\u001b[37m"
)

// Contains a standard golang log.logger for this package to exectute it's functions
type Logger struct {
    l *log.Logger
    color string
    localPrefix string
}

// Init logger global prefix
var GlobalPrefix = "project"

// Return formatted logger prefix
func formatPrefix(color string, localPrefix string) string {
    var prefix string = fmt.Sprintf("%s -> %s%s -> ", GlobalPrefix, color, localPrefix)
    return prefix
}

// Return formatted log
func formatLog(tag string, msg string, data interface{}) string {
    var log string = fmt.Sprintf("%s -> %s -> data:\n%d%s", tag, msg, data, ANSIReset)
    return log
}

// Internal log function that exported functions can call
func logGeneral(log *log.Logger, tag string, msg string, data interface{}) {
    log.Println(formatLog(tag, msg, data));
}

// Create a new logger similr to how you would create a default go log.Logger with log.New()
func New(out io.Writer, color string, localPrefix string, flag int) *Logger {
    var l *log.Logger = log.New(out, formatPrefix(color, localPrefix), flag)
    return &Logger{ l: l, color: color, localPrefix: localPrefix }
}

// Log something, you can provide:
// a tag e.g. function name,
// a msg e.g. info about something that happened,
// a data interface, e.g. []int{3, 5, 6} OR 4+2
func (l *Logger) Log(tag string, msg string, data interface{}) {
    logGeneral(l.l, tag, msg, data)
}

// Set local prefix of logger
func (l *Logger) SetLocalPrefix(localPrefix string) {
    l.localPrefix = localPrefix
    l.l.SetPrefix(formatPrefix(l.color, localPrefix))
}

// Set color of logger
func (l *Logger) SetColor(color string) {
    l.l.SetPrefix(formatPrefix(color, l.localPrefix))
}

