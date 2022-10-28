package main

import (
	"context"
	"echoinit/apps"
	"echoinit/routers"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	routers.Route(apps.Echo())
	go func() {
		port := apps.ApplicationData().Port
		retry := apps.ApplicationData().MaxStartRetry
		for startCnt := 0; startCnt < retry+1; startCnt++ {
			if e := apps.Echo().Start(":" + strconv.Itoa(port)); e != nil && e != http.ErrServerClosed {
				apps.Logs.Debug("start failed: ", e.Error())
				if startCnt >= retry {
					apps.Logs.Fatal("retry count is arrived at limit")
				}
				apps.Logs.Debug("retry starting echo. retrycount: ", startCnt+1)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	apps.Logs.Info(apps.ApplicationData().Name + " started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	apps.Logs.Info(" [*] application is terminating... ")
	shutdownCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	/*
		for the graceful shutdown.
		please retireve any required resources under here and above apps.Close().
		if needed timeout for shutdown procedure, please use shutdownCtx.
	*/
	apps.Close(shutdownCtx)
	apps.Logs.Info("Bye")
}
