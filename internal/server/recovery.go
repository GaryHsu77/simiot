package server

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	goerror "github.com/go-errors/errors"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch v := err.(type) {
				case HttpError:
					c.JSON(v.Code(), gin.H{"error": gin.H{"code": v.Code(), "message": v.Error()}})
				case string:
					c.JSON(http.StatusExpectationFailed, gin.H{"error": gin.H{"code": 100, "message": v}})
				default:
					httprequest, _ := httputil.DumpRequest(c.Request, false)
					goErr := goerror.Wrap(err, 3)
					reset := string([]byte{27, 91, 48, 109})
					log.Printf("panic recovered:\n\n%s%s\n\n%s%s", httprequest, goErr.Error(), goErr.Stack(), reset)
					c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
				}
				c.Abort()
			}
		}()
		c.Next() // execute all the handlers
	}
}
