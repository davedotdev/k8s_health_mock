package main

// LICENSE: This is private work owned by PacketPipe Ltd.
/*
	ThoughtZen

	This is a monolith that accepts webhooks, sends messages to an AI (OpenAI gpt-3.5) and sends text messages.
*/

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sethvargo/go-envconfig"

	log "github.com/sirupsen/logrus"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	var config ConfigData
	var err error

	// If a config flag is present, load a config file

	file := flag.String("config", "", "Config file")
	readyzfail := flag.Bool("readyz", false, "/readyz ok or fail (true/false)")
	livezfail := flag.Bool("livez", false, "/livez ok or fail (true/false)")
	flag.Parse()

	if *file == "" {
		// load from env vars
		ctxcfg := context.Background()
		if err := envconfig.Process(ctxcfg, &config); err != nil {
			log.Fatal(err)
		}

		if config.IsEmpty() {
			log.Error("env config data not found")
			os.Exit(1)
		}

	} else {
		err = config.Get(*file)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if *readyzfail {
		config.ReadyzFail = true
	}

	if *livezfail {
		config.LivezFail = true
	}

	// Set log level
	switch config.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	r := mux.NewRouter()

	srv := &http.Server{
		Addr:         "0.0.0.0:" + config.Port,
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      rateLimit(r), // Pass our instance of gorilla/mux in.
		// TLSConfig:    &tls.Config{},
	}

	// ================================================================
	// Server path logic lives here
	// ================================================================

	// This is publically accessible
	r.HandleFunc("/readyz", config.readyz)
	r.HandleFunc("/livez", config.livez)
	r.HandleFunc("/mode", config.mode)

	// Example code
	// r.Handle("/baz", config.MiddlewareHandler(injectActiveSession(config.HTTP_Protected_API_Presence())))
	// r.Handle("/foo/{UUID}", config.MiddlewareHandler(injectActiveSession(config.HTTP_Protected_API_Presence_VARS())))

	// ================================================================
	// End of path logic
	// ================================================================

	go func() {
		// if err := srv.ListenAndServeTLS("./certs/fullchain.pem", "./certs/privkey.pem"); err != nil {
		// 	log.Error(err)
		// }
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	log.Info("/livez fail: " + fmt.Sprint(config.LivezFail))
	log.Info("/readyz fail: " + fmt.Sprint(config.ReadyzFail))
	log.Info("Listening on port: " + config.Port)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	// ================================================================
	// Start listeners here
	// ================================================================
	// err = config.StartThing(ctx)
	// if err != nil {
	// 	log.Error(err)
	// }

	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
			srv.Shutdown(ctx)
		case <-ctx.Done():
			srv.Shutdown(ctx)
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()

	<-ctx.Done()
	log.Info("\nexiting...")
	os.Exit(0)
}
