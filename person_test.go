package main

import (
	"testing"
)

func TestPersonAdultCheck(t *testing.T) {
	p := person{"Bob", 20}
	if !p.isAdult() {
		t.Fatal("Failed age check - supposed to be adult")
	}

	p = person{"Bobby", 12}
	if p.isAdult() {
		t.Fatal("Failed age check - not supposed to be adult")
	}
}
