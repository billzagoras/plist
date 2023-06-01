package server

import (
	"context"
	"log"
	"net/http"
	"plist/server/modules/ptlists"
	"plist/utils"
	"time"

	"github.com/gorilla/mux"
)

// ApplicationServer a blueprint of a PowerFactors application server.
type ApplicationServer struct {
	httpServer *http.Server
	Router     http.Handler
	app        *Application
}

// Setup initiates the application server.
func (server *ApplicationServer) Setup() {
	router := mux.NewRouter()

	app := NewApplication()

	// router.Use(middlewares.MetricsMiddleware(metrics))

	server.app = app

	// Register Routes
	ptlists.Setup(router, app.PtList)

	server.Router = router
}

// Run executed the application server.
func (server *ApplicationServer) Run(port string) error {
	server.httpServer = &http.Server{
		Addr:              "0.0.0.0:" + port,
		Handler:           server.Router,
		ReadTimeout:       10 * time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      60 * time.Second,
	}
	log.Println(server.httpServer)

	if err := server.httpServer.ListenAndServe(); utils.CheckErr(err) && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Close terminates the application server.
func (server *ApplicationServer) Close() error {
	server.app.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return server.httpServer.Shutdown(ctx)
}
