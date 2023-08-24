package main

import (
	"avito_test/internals/app"
	"avito_test/internals/cfg"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
)


func main (){
	logger := new(zerolog.Logger)

	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background()) 

	c := make(chan os.Signal, 1) 

	signal.Notify(c, os.Interrupt,syscall.SIGTERM)

	server := app.NewServer(config, ctx,logger) 

 
	go func() { 
		oscall := <-c 

		logger.Info().Msg(fmt.Sprintf("system call:%+v", oscall))

		if err := server.Shutdown(); err != nil { 
			logger.Err(err)
		}

		cancel() 
	}()
	if err := server.Serve(); err != nil {
		logger.Err(err)
	}
}


