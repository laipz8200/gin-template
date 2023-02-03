package api

import (
	"_template_/api/controllers"
	"_template_/constants"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laipz8200/i18n"
)

func setup() {
	engine.GET("/ping", handle(controllers.Ping))
}

// handle
func handle[Req any, Resp any](fn func(ctx context.Context, req Req) (resp Resp, code int, err error)) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var lang string
		if headers, ok := c.Request.Header["Language"]; ok && len(headers) != 0 {
			lang = headers[0]
		}

		// Set language
		ctx := context.WithValue(c.Request.Context(), constants.KEY_LANGUAGE, lang)

		// Bind request paramters
		var req Req
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"error": i18n.Lang(lang).Sprintf("Bad request: %s", err.Error()),
			})
			return
		}

		// Execute controller function
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
