// Option function is a way to pass optional paremeters to a function
// that's used to construct or modify
package main

import "fmt"

// PersonOption presents a funcion that updates person property
type PersonOption func(p *Person)

// Person presents information of a person
type Person struct {
	Name, Address string
	Age           int
}

// NewPerson creates new a person with information options
func NewPerson(opts ...PersonOption) *Person {
	p := &Person{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// WithName set name of a person
func WithName(name string) PersonOption {
	return func(p *Person) {
		p.Name = name
	}
}

// WithAddress set address of a person
func WithAddress(address string) PersonOption {
	return func(p *Person) {
		p.Address = address
	}
}

// WithAge set age of a person
func WithAge(age int) PersonOption {
	return func(p *Person) {
		p.Age = age
	}
}

func main() {
	// Create new a new person with optional field values
	p := NewPerson(WithAge(3), WithName("Chicken"))
	fmt.Printf("Person information: %v\n", *p)

	// Create new a new client with optional request modifiers
	client := NewClient(WithAuth("JHJSjhT"), WithLang("en"))
	req := client.Request()
	req.Body["username"] = "kenvina"
	req.Body["password"] = "ken12821"

	fmt.Printf("Header of request: %v\n", req.Header)
	fmt.Printf("Body of request: %v\n", req.Body)
}

// Request presents a request information
type Request struct {
	Header, Body map[string]string
}

// RequestModifier can modifies a request
type RequestModifier func(req *Request)

// Client applies request modifiers to a request
type Client struct {
	modifiers []RequestModifier
}

// NewClient create a new client with request modifiers
func NewClient(modifiers ...RequestModifier) *Client {
	return &Client{modifiers}
}

// Request constructs a new request
func (c *Client) Request() *Request {
	req := &Request{
		make(map[string]string),
		make(map[string]string),
	}

	for _, modifier := range c.modifiers {
		modifier(req)
	}

	return req
}

// WithAuth adds the authorization header to the req
func WithAuth(token string) RequestModifier {
	return func(req *Request) {
		req.Header["authorization"] = token
	}
}

// WithLang adds the language header to the req
func WithLang(lang string) RequestModifier {
	return func(req *Request) {
		req.Header["language"] = lang
	}
}
