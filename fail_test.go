package main

import "testing"

func TestIntentionalFailure(t *testing.T) {
    t.Fatalf("intentional failure")
}
