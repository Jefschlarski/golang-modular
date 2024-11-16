package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jefschlarski/golang-modular/pkg/security"
	"github.com/gin-gonic/gin"
)

// Logger middleware
func Logger(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		next(ctx)
		fmt.Printf("%s %s %s %s %s %d %s\n", start.Format("02/Jan/2006:15:04:05 -0700"), ctx.Request.Method, ctx.Request.URL.Path, ctx.Request.Proto, ctx.ClientIP(), ctx.Writer.Status(), time.Since(start))
	}
}

// Authenticate middleware
func Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := security.ValidateToken(ctx); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		next(ctx)

	}
}
