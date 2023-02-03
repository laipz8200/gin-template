package controllers

import (
	"_template_/constants"
	"context"

	"github.com/laipz8200/i18n"
)

func Ping(ctx context.Context, _ any) (resp string, code int, err error) {
	language := ctx.Value(constants.KEY_LANGUAGE)
	return i18n.Lang(language.(string)).Sprintf("pong"), 200, nil
}
