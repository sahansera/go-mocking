package main

import (
	"testing"
)

func TestDoSomething(t *testing.T) {
	// Use the mock implementation.
	c := &MockClient{}
	s := NewService(c)
	out := s.client.Get()

	if out != "Hello from MockClient" {
		t.Errorf("Expected 'Hello from MockClient', got '%s'", out)
	}
}
