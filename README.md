# slogger: simple go logger

This is a small go package that provides two simple logging functions to aid in development


**install in project**

From project directory run:

```
go get github.com/jesse-rb/slogger-go 
```


**example usage**

```
package main

import (
	slogger "github.com/jesse-rb/slogger-go"
)

func main() {
    slogger.LogInfo("main", "Some info.", 2+4)
    slogger.LogError("main", "Some error", 4+0)
}
```

** prewview output **

[preview image A](previews/preview-a.png)
