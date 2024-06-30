package frame

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// Recovery
// @Description: 捕获 panic，并且将堆栈信息打印在日志中，向用户返回 Internal Server Error。
// @return HandlerFunc
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("[Recovery] %s", err)
				log.Panicf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}

// trace
// @Description: 获取触发 panic 的堆栈信息
// @param message
// @return string
func trace(message string) string {
	var pcs [32]uintptr
	// 第 0 个 Caller 是 Callers 本身
	// 第 1 个是上一层 trace
	// 第 2 个是再上一层的 defer func。
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
