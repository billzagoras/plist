package server

import (
	"plist/internal/app/ptlist"
)

// Application structure holds all services.
type Application struct {
	PtList *ptlist.Service
}

// Close function.
func (a *Application) Close() error {
	return nil
}

// NewApplication constructor creates all services with inner dependencies.
func NewApplication() *Application {
	ptlistService := ptlist.NewService()

	return &Application{
		PtList: ptlistService,
	}
}
