package main

import "fmt"

// An interface that defines the Get() method.
type Client interface {
	Get() string
}

// Actual implementation of the client.
type ConcreteClient struct {
}

// Implementing Get() from the Client interface.
func (c ConcreteClient) Get() string {
	return "Hello from ConcreteClient"
}

// This is the mock implementation we will use in the tests.
type MockClient struct {
}

// Implementing Get() from the Client interface.
func (c *MockClient) Get() string {
	return "Hello from MockClient"
}

// Note how we are using the interface type Client here.
type Service struct {
	client Client
}

// Example to show how dependency injection works.
func NewService(c Client) *Service {
	return &Service{
		client: c,
	}
}

func main() {
	// Use the concrete implementation.
	c := &ConcreteClient{}
	s := NewService(c)
	out := s.client.Get()
	fmt.Println(out)
}
