package controllers

import (
	"_template_/api/schemas"
	"_template_/constants"
	"context"

	"github.com/laipz8200/i18n"
)

func Ping(ctx context.Context, _ any) (resp schemas.Message, code int, err error) {
	language := ctx.Value(constants.KeyLanguage)
	return schemas.Message(i18n.Lang(language.(string)).Sprintf("pong")), 200, nil
}
