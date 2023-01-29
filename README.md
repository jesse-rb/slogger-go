# slogger: simple go logger

This is a small go package that provides two simple logging functions to aid in development

[pkg.go.dev docs page](https://pkg.go.dev/github.com/jesse-rb/slogger-go)


**install in project**

Using go modules, from your project directory you can run:

```
go get github.com/jesse-rb/slogger-go 
```


**example usage**

```
package main

import (
	slogger "github.com/jesse-rb/slogger-go"
    log
    os
)

func main() {
    // Init new logger
    var localPrefix string = "example e.g. <package name>"
    var flags int = log.Ldate+log.Lshortfile

    var l = slogger.New(os.Stdout, slogger.ANSIGreen, localPrefix, flags)
    
    // Log to output
    var tag string = "example e.g. <function name>"
    var msg string = "Testing a new logger"
    var data []int = []int{3, 5, 7}

    l.Log(tag, msg, data)
}
```

**preview logs**

![preview image A](previews/preview-a.png)


## References

Referenced [3 bit 8 colours from Haoyi's Programming Blog](https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html#8-colors)
as most terminals should supprot this
```
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
```

