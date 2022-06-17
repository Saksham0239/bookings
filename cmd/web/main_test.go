package main

import "testing"

func TestRun(t *testing.T) {

	err := run()

	if err != nil {
		t.Error("run() failed")
	}
}

/*

go test    for normal testing

go test -v  for verbose testing

go test -coverprofile=coverage for overall testing percentage

go tool cover -html=coverage   for getting testing coverage on html


*/
