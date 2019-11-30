package main

import "testing"

func TestGreet(t *testing.T) {
	var greeting = Greet()

	if greeting != "Hello world!" {
		t.Errorf("Greet was bad, got: %s, wanted: %s", greeting, "Hello world!")
	}
}