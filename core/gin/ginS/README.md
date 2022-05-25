# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/go-colon/colon/core/gin"
	"github.com/go-colon/colon/framework/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```