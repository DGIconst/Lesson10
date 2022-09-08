package hello

import "testing"

func TestHello(t *testing.T) {

	got := hello()
	want := "Hello there!"

	if got != want {

		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMorning(t *testing.T) {

	got := morning()
	want := "Good morning!"

	if got != want {

		t.Errorf("got %s, want %s", got, want)
	}
}
