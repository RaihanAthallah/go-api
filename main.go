package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inventory-management-tokobejo/go-api/config"
	"github.com/inventory-management-tokobejo/go-api/infrastructure/io"
	"github.com/inventory-management-tokobejo/go-api/infrastructure/registry"
	"github.com/inventory-management-tokobejo/go-api/infrastructure/router"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	config.ReadConfig()

	sentryLog := io.SentryLog{}
	if config.C.Sentry.IsActive {
		io.InitSentry(config.C.Sentry.Dsn, config.C.Server.Env)
	}

	loggerConfig := io.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      io.Debug,
		ConsoleJSONFormat: true,
		EnableFile:        true,
		FileLevel:         io.Info,
		FileJSONFormat:    true,
		FileLocation:      "log/core-services.log",
	}
	_, err := io.NewLogger(loggerConfig)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}

	isDbLogged := false
	if config.C.Server.Env == "stage" {
		gin.SetMode(gin.ReleaseMode)
		isDbLogged = true
	}

	db := io.InitDB(isDbLogged)
	reg := registry.NewRegistry(db)

	r := gin.Default()
	r = router.Routes(r, reg)

	fmt.Println("Server listen at http://localhost" + ": " + config.C.Server.Port)

	srv := &http.Server{
		Addr:    ":" + config.C.Server.Port,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sentryLog.CaptureError(fmt.Errorf("start Server has failed: %v", err))
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sentryLog.CaptureError(fmt.Errorf("shutdown Server has failed: %v", err))
	}

	fmt.Println("Server exiting")

}
