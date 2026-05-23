package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	game := Game()
	if game == nil {
		t.Errorf("Wah wah no game")
	}
}