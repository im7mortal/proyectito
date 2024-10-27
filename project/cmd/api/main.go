// Created by Petr Lozhkin

// api micro service

package main

import (
	"context"
	"errors"
	"github.com/golang/glog"
	"github.com/im7mortal/project/pkg/keygen"
	"github.com/im7mortal/project/pkg/server"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

const gracefulShutdownD = time.Second * 10
const defaultPort = ":3002"

func main() {

	// Main context
	world, endOfTheWorld := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		glog.Infof("Got signal %s; Shutdown all operations", sig.String())
		endOfTheWorld()
	}()

	sdkEngine := server.New(keygen.NewKeyGenerator(), "public")

	srv := &http.Server{
		Addr:    defaultPort,
		Handler: sdkEngine.GetMainEngine(),
	}
	go func() {
		if r := recover(); r != nil {
			glog.Errorf("panic %s \nstacktrace from panic: \n%s", r, string(debug.Stack()))
		}
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			glog.Error(err)
			endOfTheWorld()
		}
	}()
	select {
	case <-world.Done():

	}

	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownD)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Error("Server forced to shutdown:", err)
	}
	glog.Error("correct exit")
}
