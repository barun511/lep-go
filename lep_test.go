package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	hello := HelloWorld()
	if hello != "something" {
		t.Errorf("Test failed");
	}
}