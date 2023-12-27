package main

import "testing"

func TestSayHello(t *testing.T) {
	name := "Ramu"
	got := sayHello(name)
	want := "Hello, " + name
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
