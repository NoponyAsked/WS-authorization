package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	for i := 0; i < 15; i++ {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		y := r.Intn(1000)
		hello(y)
		t.Logf("%d - ", y)
		if y%2 == 0 {
			t.Error("gotta be hello, but no hello")
		} else {
			t.Log("hello, world")
		}

	}
}
