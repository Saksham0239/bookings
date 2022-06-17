package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) { //before running the test run this method for setup enivronment for the tests

	os.Exit(m.Run())
}

//used as a custom http.Hanlder
type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
