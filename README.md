# slogger: simple go logger

This is a small go package that provides a few simple logging functions to aid in development.
This package allows you to register a new logger very similar to how you would a standard go `log.Logger`,
with the added ability to pass paremeters to configure a log color (using ANSI color codes) and some log prefixese.

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
    // Declare some loggers
    infoLogger := slogger.New(os.Stdout, slogger.ANSIBlue, "info", log.Lshortfile+log.Ldate);
    errorLogger := slogger.New(os.Stderr, slogger.ANSIRed, "error", log.Lshortfile+log.Ldate);

    // Log some things
    infoLogger.Log("main", "Something worth noting happened", 2+4)
    errorLogger.Log("main", "Some horrible error happened", []int{3, 5, 7})
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

