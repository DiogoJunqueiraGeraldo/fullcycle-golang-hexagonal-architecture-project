package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/adapters/web/handler"
	"github.com/DiogoJunqueiraGeraldo/fullcycle-golang-hexagonal-architecture-project/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func NewWebServer(service application.ProductServiceInterface) *Webserver {
	return &Webserver{Service: service}
}

func (w *Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandler(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
