package apps

import (
	"context"

	"github.com/spf13/viper"
)

type typeAppCtxKey string

type AppData struct {
	Name          string
	Port          int
	MaxStartRetry int
}

const (
	ctxKeyAppData   typeAppCtxKey = "CtxKeyAppData"
	applicationName string        = "echoinit"
)

func initContext() (context.Context, context.CancelFunc) {
	ctx := context.Background()

	// ===============================================================
	// load configuraiton
	viper.SetConfigName(applicationName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./apps/conf/")
	viper.AddConfigPath("$HOME/.config/" + applicationName)
	err := viper.ReadInConfig()
	if err != nil {
		Logs.Fatal(err)
	}

	pvEcho := viper.GetStringMap("Echo")
	port, ok := pvEcho["listenport"].(int)
	if !ok {
		port = 1323
	}
	Logs.Debug("echo port: ", port)

	retry, ok := pvEcho["maxstartretry"].(int)
	if !ok {
		retry = 1
	}
	Logs.Debug("echo retry: ", retry)

	// viper.SetConfigName("restapis")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("./apps/conf/apis/")
	// err = viper.MergeInConfig()
	// if err != nil {
	// 	Logs.Fatal(err)
	// }

	// pvRestApis := viper.GetStringMap("RestApis")
	// Logs.Debug("%+v", pvRestApis)

	data := &AppData{
		Name:          applicationName,
		Port:          port,
		MaxStartRetry: retry,
	}
	ctx = context.WithValue(ctx, ctxKeyAppData, data)

	return context.WithCancel(ctx)
}
