package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dimoynwa/product-api/handlers"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/product", ph)

	serv := &http.Server{
		Addr:         *bindAddress,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		l.Printf("Starting server on port %v...\n ", serv.Addr)
		if err := serv.ListenAndServe(); err != nil {
			l.Fatalf("error starting the server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Printf("Received terminal signal %v, Graceful shutdown\n", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	serv.Shutdown(tc)
}
