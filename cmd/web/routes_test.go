package main

import (
	"testing"

	"github.com/Saksham0239/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {

	var app config.AppConfig

	res := routes(&app)

	switch v := res.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Errorf("type is not chi.Mux , type is %T ", v)
	}
}
