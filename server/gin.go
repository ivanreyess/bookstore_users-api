package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartGinApplication start a web server based on gin framework
func StartGinApplication() {
	mapUrls()
	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      router,            // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		fmt.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	cancelFunc()
	_ = s.Shutdown(ctx)
}
