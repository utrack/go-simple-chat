package main

import (
	"flag"
	"github.com/Sirupsen/logrus"
	"github.com/kennygrant/sanitize"
	"github.com/utrack/go-simple-chat/client/bot"
	"github.com/utrack/go-simple-chat/hub"
	"github.com/utrack/go-simple-chat/interface/http"
	"html/template"
	"net/http"
)

var httpAddr = flag.String("addr", ":8080", "HTTP socket listening address")
var logLevel = flag.String("log", "info", "Logging level: Debug,Info,Warn,Error,Fatal")
var tmplPath = flag.String("static", "assets/static/chat.tmpl", "Path to the page's template")

func main() {
	flag.Parse()
	lvl, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetLevel(lvl)
	logrus.WithField("level", lvl.String()).Warn("Logging level was set")

	logrus.Info("Starting up")
	logrus.Info("Initiating the Hub...")

	h := hub.NewHub(hub.DefaultNameChecker, sanitize.HTML, loggerFunc)
	h.Run()
	h.RegisterClient(clientBot.NewBot(), "ChuckServ")

	staticTemplate := template.Must(template.ParseFiles(*tmplPath))
	http.HandleFunc(`/`, serveStatic(staticTemplate))
	http.HandleFunc(`/ws`, ifaceHttp.ServeWs(h))

	logrus.WithField("addr", *httpAddr).Info("Starting HTTP server")
	err = http.ListenAndServe(*httpAddr, nil)
	if err != nil {
		logrus.WithError(err).Fatal("Error when starting HTTP server!")
	}
}
