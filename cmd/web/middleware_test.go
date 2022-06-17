package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) { //Testing if the type is myHanlder or not

	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing

	default:
		t.Errorf("The value is not of type http.Hanlder , but of type %T", v)
	}
}

func TestSessionLoad(t *testing.T) { //Testing if the type is myHanlder or not

	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing

	default:
		t.Errorf("The value is not of type http.Hanlder , but of type %T", v)
	}
}
