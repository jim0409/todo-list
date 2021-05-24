package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	. "todo-list/logger"
	. "todo-list/models"

	"todo-list/config"
	"todo-list/router"

	"github.com/gin-gonic/gin"
)

var (
	confPath    = flag.String("config", "./config/app.dev.ini", "config location")
	checkcommit = flag.Bool("version", false, "burry code for check version")

	confInfo     *config.Config
	gitcommitnum string
)

func checkComimit() {
	log.Println(gitcommitnum)
}

func Init() error {
	flag.Parse()
	// if there is a needed to check git commit num ... print it out and exit
	if *checkcommit {
		checkComimit()
		os.Exit(1)
	}

	// read config and pass variables ...
	var err error
	confInfo, err = config.InitConfig(confPath)
	if err != nil {
		return fmt.Errorf("Init config err: %v", err)
	}

	// initialize logger
	if err = InitLog(confInfo.LogConf.LogPath, confInfo.LogConf.LogLevel); err != nil {
		return fmt.Errorf("init logger err: %v", err)
	}

	// initialize both mysql and redis db
	if err = InitDb(&confInfo.DbConf); err != nil {
		return fmt.Errorf("init db err: %v", err)
	}

	return nil
}

func main() {
	//catch global panic
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic err: %v", err)
		}
	}()

	err := Init()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	route := gin.Default()
	router.ApiRouter(route)

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%v", confInfo.HttpPort),
		Handler: route,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Log.Info(fmt.Sprintf("http listen : %v\n", err))
			panic(err)
		}
	}()

	gracefulShutdown()
}

func gracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	Log.Info("awaiting signal")
	<-done
	Log.Info("exiting")
}
