package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"goblog/initialize"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	initialize.ViperInit()
	initialize.ZeroLogInit()
	initialize.DBInit()

	router := initialize.RouterInit()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Err(err).Stack().Msgf("listen error")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server ...")
	initialize.DBDefer()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Err(err).Stack().Msg("Server Shutdown")
	}
	log.Info().Msg("Server exiting")
}
