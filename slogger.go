package logger

import (
	"fmt"
	"io"
	"log"
)

// 3 bit color ansi color codes for colored logging output
// srouce: https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html#8-colors
const (
	ANSIReset   = "\u001b[0m"
	ANSIRed     = "\u001b[31m"
	ANSIBlue    = "\u001b[34m"
	ANSIBlack   = "\u001b[30m"
	ANSIGreen   = "\u001b[32m"
	ANSIYellow  = "\u001b[33m"
	ANSIMagenta = "\u001b[35m"
	ANSICyan    = "\u001b[36m"
	ANSIWhite   = "\u001b[37m"
)

// Contains a standard golang log.logger for this package to exectute it's functions
type Logger struct {
	l           *log.Logger
	color       string
	localPrefix string
	calldepth   int
}

// Init logger global prefix
var GlobalPrefix = "project"

// Return formatted logger prefix
func formatPrefix(color string, localPrefix string) string {
	prefix := ""

	if GlobalPrefix != "" {
		prefix = fmt.Sprintf("%s -> ", GlobalPrefix)
	}
	if color != "" {
		prefix = fmt.Sprintf("%s%s", prefix, color)
	}
	if localPrefix != "" {
		prefix = fmt.Sprintf("%s%s -> ", prefix, localPrefix)
	}

	return prefix
}

// Return formatted log
func formatLog(tag string, msg string, data any, shouldANSIReset bool) string {
	log := msg

	if tag != "" {
		log = fmt.Sprintf("%s -> %s", tag, log)
	}
	if data != nil {
		log = fmt.Sprintf("%s -> data:\n%#v", log, data)
	}
	if shouldANSIReset {
		log = fmt.Sprintf("%s%s", log, ANSIReset)
	}

	return log
}

// Create a new logger similr to how you would create a default go log.Logger with log.New()
func New(out io.Writer, color string, localPrefix string, flag int) *Logger {
	formattedPrefix := formatPrefix(color, localPrefix)
	var l *log.Logger = log.New(out, formattedPrefix, flag)
	return &Logger{l: l, color: color, localPrefix: localPrefix, calldepth: 1}
}

// Log something, you can provide:
// a tag e.g. function name,
// a msg e.g. info about something that happened,
// a data interface, e.g. []int{3, 5, 6} OR 4+2
func (l *Logger) Log(tag string, msg string, data any) {
	shouldANSIReset := l.color != ""
	formattedLog := formatLog(tag, msg, data, shouldANSIReset)
	l.l.Output(l.calldepth+1, formattedLog)
}

// Set local prefix of logger
func (l *Logger) SetLocalPrefix(localPrefix string) {
	l.localPrefix = localPrefix
	formattedPrefix := formatPrefix(l.color, localPrefix)
	l.l.SetPrefix(formattedPrefix)
}

// Set color of logger
func (l *Logger) SetColor(color string) {
	l.color = color
	formattedPrefix := formatPrefix(color, l.localPrefix)
	l.l.SetPrefix(formattedPrefix)
}

// Set the calldepth of the logger, this allows log flags such as log.Lshortfile or log.Llongfile to report
// the file/line of caller functions from further up the all stack if desired, otherwise a sensible default is used
// for behavior similar to go's log.Logger
func (l *Logger) SetCalldepth(calldepth int) {
	l.calldepth = calldepth
}
