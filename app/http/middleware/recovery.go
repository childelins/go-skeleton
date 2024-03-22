package middleware

import (
	"errors"
	"net"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/childelins/go-skeleton/pkg/log"
	"github.com/childelins/go-skeleton/pkg/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery 使用 zap.Error() 来记录 Panic 和 call stack
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 判断链接是否中断，客户端中断连接为正常行为，不需要记录堆栈信息
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 链接中断的情况
				if brokenPipe {
					// 获取用户的请求信息
					httpRequest, _ := httputil.DumpRequest(c.Request, false)

					log.Warn(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					// 链接已断开，无法写状态码
					return
				}

				// 如果不是链接中断，就开始记录堆栈信息
				log.Error("recovery from panic", zap.Any("error", err))

				// 返回 500 状态码
				resp.Error(c, "internal server error")
			}
		}()

		c.Next()
	}
}
