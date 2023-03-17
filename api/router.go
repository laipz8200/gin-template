package api

import (
	"_template_/api/controllers"
	"_template_/api/middleware"
	"_template_/config"
	"strings"
)

func (s *server) setup() {
	s.router.GET("/ping", handle(controllers.Ping))
	app := s.router.Group(strings.ToLower(config.AppName()))
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
