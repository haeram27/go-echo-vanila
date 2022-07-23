package main

import (
	"echoinit/apps"
	"echoinit/routers"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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
	signal.Notify(quit, os.Interrupt)

	<-quit
	apps.Logs.Info("to confirm shutdown, please send interrupt again")
	<-quit
	apps.Logs.Info("terminating application")
	apps.Close()
	apps.Logs.Info("Bye")
}
