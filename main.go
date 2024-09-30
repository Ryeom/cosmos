package main

import (
	"fmt"
	"github.com/Ryeom/cosmos/log"
	"github.com/Ryeom/cosmos/mongo"
	"github.com/Ryeom/cosmos/router"
	"github.com/Ryeom/cosmos/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var EndSignal = make(chan os.Signal, 1)

func Cleanup() {
	signal.Notify(EndSignal, os.Interrupt)
	signal.Notify(EndSignal, syscall.SIGQUIT)
}

func init() {
	// [init 1] setting execution mode and environment
	err := util.MustSetArguments()
	if err != nil {
		panic(exitCode{1}) // 그냥 그자리에서 패닉띄우기
	}
	// [init 2] application log setting
	log.MustInitializeApplicationLog()
	// [init 3] Change different settings depending on the {mode}
	util.MustInitializeSetting()

	Cleanup()
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(log.CreateCustomLogConfig()))
	e.Use(router.Cors)

	mongo.InitialiseMongo()

	router.Initialize(e)

	go func() {
		<-EndSignal
		log.Logger.Error("APPLICATION END.")

		_ = log.ServerLogDesc.Close()
		_ = log.AccessLogDesc.Close()
	}()

	log.UpTime = time.Now().Format(log.DefaultDateFormat)
	port := viper.GetString("cosmos.port")
	log.Logger.Info("SERVICE PORT : ", port, " ...")
	log.Logger.Fatal(e.Start(":" + port))
}

type exitCode struct{ Code int }

func handleExit() {
	if e := recover(); e != nil {
		if exit, ok := e.(exitCode); ok {
			if exit.Code != 0 {
				fmt.Fprintln(os.Stderr, "Fail !!!!", time.Now().Format("2006-01-02 15:04:05.999999"))
			} else {
				fmt.Fprintln(os.Stderr, "Stop !!!!", time.Now().Format("2006-01-02T15:04:05.999999"))
			}
			os.Exit(exit.Code)
		}
		panic(e)
	}
}
