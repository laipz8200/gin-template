package api

import (
	"_template_/api/controllers"
	"_template_/api/middleware"
	"_template_/api/schemas"
	"_template_/config"
	"_template_/constants"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/laipz8200/i18n"
)

func setup() {
	router.GET("/ping", handle(controllers.Ping))
	app := router.Group(strings.ToLower(config.AppName()))
	{
		private := app.Group("")
		private.Use(middleware.AuthMiddleware)
		{
			private.GET("")
		}

		public := app.Group("")
		{
			public.GET("")
		}
	}
}

// handle
func handle[Req any, Resp any](fn func(ctx context.Context, req Req) (resp Resp, code int, err error)) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var lang string
		if headers, ok := c.Request.Header["Language"]; ok && len(headers) != 0 {
			lang = headers[0]
		}

		// Set language
		c.Set(constants.KeyLanguage, lang)
		ctx := c.Request.Context()

		// Bind request parameters
		var req Req

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, schemas.ErrorMessage{
				Code:  http.StatusBadRequest,
				Error: i18n.Lang(lang).Sprintf("Bad request: %s", err.Error()),
			})
			return
		}

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, schemas.ErrorMessage{
				Code:  http.StatusBadRequest,
				Error: i18n.Lang(lang).Sprintf("Bad request: %s", err.Error()),
			})
			return
		}

		// Execute controller function
		select {
		case <-ctx.Done():
			c.Abort()
			return
		default:
			resp, code, err := fn(ctx, req)
			if err != nil {
				c.JSON(code, schemas.ErrorMessage{
					Code:  code,
					Error: i18n.Lang(lang).Sprintf(err.Error()),
				})
				return
			}

			if code == http.StatusNoContent {
				c.Status(code)
				return
			}

			c.JSON(code, schemas.Response{
				Code: code,
				Data: resp,
			})
		}
	}
}
