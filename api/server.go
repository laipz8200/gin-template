package api

import (
	"_template_/config"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
	router *gin.RouterGroup
	srv    *http.Server
)

func Init() {
	if !config.Debug() {
		gin.SetMode(gin.ReleaseMode)
	}
	engine = gin.Default()
	router = engine.Group("")
	srv = &http.Server{
		Handler: engine,
	}
}

// Run
func Run(addr string) (err error) {
	setup()
	srv.Addr = addr
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		return
	}

	return
}
