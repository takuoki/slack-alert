// Package slacka is a package to send alert to Slack App.
package slacka

import "fmt"

// Client is a client struct of slacka package.
type Client struct {
	userame     string
	serviceName string
	iconURL     string
	urlMap      map[string]string
	errorURL    string
}

// New returns slacka client.
func New(userame, serviceName, iconURL string) *Client {
	return &Client{
		userame:     userame,
		serviceName: serviceName,
		iconURL:     iconURL,
		urlMap:      map[string]string{},
	}
}

// SetURL is a function to set Webhook URL to the client.
// If the key is already set, this function throws a panic.
func (c *Client) SetURL(key, url string) {
	if _, ok := c.urlMap[key]; ok {
		panic(fmt.Sprintf("the key is already set (key: %q)", key))
	}
	c.urlMap[key] = url
}

// SetErrorURL is a function to set Webhook URL for error alert to the client.
// If the error alert URL is already set, this function throws a panic.
func (c *Client) SetErrorURL(url string) {
	if c.errorURL != "" {
		panic("error url is already set")
	}
	c.errorURL = url
}
