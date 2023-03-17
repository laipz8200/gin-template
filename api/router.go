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
	"github.com/go-playground/validator"
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
func handle[Req any, Resp schemas.Response](fn func(ctx context.Context, req Req) (resp Resp, code int, err error)) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var lang string
		if headers, ok := c.Request.Header["Language"]; ok && len(headers) != 0 {
			lang = headers[0]
		}

		// Set language params.
		c.Set(string(constants.KeyLanguage), lang)

		// Move keys to standard context.
		ctx := c.Request.Context()
		for k, v := range c.Keys {
			ctx = context.WithValue(ctx, constants.ContextKey(k), v)
		}

		// Bind request parameters.
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

		// Validate request parameters.
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			errors := make([]string, 0)
			for _, err := range err.(validator.ValidationErrors) {
				tag := i18n.Lang(lang).Sprintf(err.Tag())
				param := i18n.Lang(lang).Sprintf(err.Param())
				errors = append(errors, i18n.Lang(lang).Sprintf("Validation error: key %s is %s %s, got %v", err.Field(), tag, param, err.Value()))
			}

			c.JSON(http.StatusBadRequest, schemas.ErrorMessage{
				Code:  http.StatusBadRequest,
				Error: strings.Join(errors, "; "),
			})
			return
		}

		// Execute controller function.
		select {
		case <-ctx.Done():
			c.Abort()
			return
		default:
			result, code, err := fn(ctx, req)
			if err != nil {
				if code == 0 {
					code = http.StatusInternalServerError
				}
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

			if code == 0 {
				code = http.StatusOK
			}
			c.JSON(code, result.ToResponse(code))
		}
	}
}
