package web

import (
	"context"
	"net/http"
	"template/internal/constants"
	"template/internal/web/controllers"

	"github.com/gin-gonic/gin"
	"github.com/laipz8200/i18n"
)

func setup() {
	engine.GET("/ping", handle(controllers.Ping))
}

func handle[Req any, Resp any](fn func(ctx context.Context, req Req) (resp Resp, code int, err error)) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var lang string
		if headers, ok := c.Request.Header["Language"]; ok && len(headers) != 0 {
			lang = headers[0]
		}

		ctx := context.WithValue(c.Request.Context(), constants.KEY_LANGUAGE, lang)

		var req Req
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": i18n.Lang(lang).Sprintf("Bad request: %s", err.Error()),
			})
			return
		}

		resp, code, err := fn(ctx, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  code,
				"error": i18n.Lang(lang).Sprintf("Internel server error: %s", err.Error()),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": resp,
		})
	}
}
