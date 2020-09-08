package main

import "testing"

func TestNewServer(t *testing.T) {
	s, err := NewServer()
	if s == nil || err != nil {
		t.Error("failed to create new server")
	}
}
