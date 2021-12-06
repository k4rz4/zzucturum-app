package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"zzucturum-app/pkg/http-transport"
	"zzucturum-app/cmd/api/routes"
)

var logger *logrus.Logger

func run(args []string) int {

	//bindAddress := flag.String("ip", "0.0.0.0", "IP address to bind")
	bindAddress := "0.0.0.0"
	//listenPort := flag.Int("port", 25478, "port number to listen on")
	listenPort := 25478
	tlsListenPort := flag.Int("tlsport", 25443, "port number to listen on with TLS")
	//tokenFlag := flag.String("token", "", "specify the security token (it is automatically generated if empty)")
	tokenFlag := "f297d33fec0caba3e11d"
	logLevelFlag := flag.String("loglevel", "info", "logging level")
	certFile := flag.String("cert", "", "path to certificate file")
	keyFile := flag.String("key", "", "path to key file")

	flag.Parse()

	if logLevel, err := logrus.ParseLevel(*logLevelFlag); err != nil {
		logrus.WithError(err).Error("failed to parse logging level, so set to default")
	} else {
		logger.Level = logLevel
	}
	token := tokenFlag
	if token == "" {
		count := 10
		b := make([]byte, count)
		if _, err := rand.Read(b); err != nil {
			logger.WithError(err).Fatal("could not generate token")
			return 1
		}
		token = fmt.Sprintf("%x", b)
		logger.WithField("token", token).Warn("token generated")
	}

	tlsEnabled := *certFile != "" && *keyFile != ""

	eventHandler := http_transport.NewServer(token, router.Routes)
	http.Handle("/", eventHandler)

	errors := make(chan error)

	go func() {
		logger.WithFields(logrus.Fields{
			"ip":               bindAddress,
			"port":             listenPort,
			"token":            token,
		}).Info("start listening")

		if err := http.ListenAndServe(fmt.Sprintf("%s:%d", bindAddress, listenPort), nil); err != nil {
			errors <- err
		}
	}()

	if tlsEnabled {
		go func() {
			logger.WithFields(logrus.Fields{
				"cert": *certFile,
				"key":  *keyFile,
				"port": *tlsListenPort,
			}).Info("start listening TLS")

			if err := http.ListenAndServeTLS(fmt.Sprintf("%s:%d", bindAddress, tlsListenPort), *certFile, *keyFile, nil); err != nil {
				errors <- err
			}
		}()
	}

	err := <-errors
	logger.WithError(err).Info("closing server")

	return 0
}

func main() {
	logger = logrus.New()
	logger.Info("starting up server")

	result := run(os.Args)
	os.Exit(result)
}
