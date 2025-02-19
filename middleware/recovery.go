package middleware

import (
	"net/http"
	"social-todo-list/common"

	"github.com/gin-gonic/gin"
)

func Recovery() func(*gin.Context) {
	return func(ctx *gin.Context) {
		defer func ()  {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrorInternalServer(err))
				}
			}
		}()

		ctx.Next()
	}
}