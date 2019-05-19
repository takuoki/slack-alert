// Package slacka is a package to send alert to Slack App.
package slacka

// Client is a client struct of slacka package.
type Client struct {
	projectName string
	serviceName string
	iconEmoji   string
	urlMap      map[string]string
	errorURL    string
}

// New returns slacka client.
func New(projectName, serviceName, iconEmoji string) *Client {
	return &Client{
		projectName: projectName,
		serviceName: serviceName,
		iconEmoji:   iconEmoji,
		urlMap:      map[string]string{},
	}
}

// SetURL is a function to set Webhook URL to the client.
// If the key is already set, this function overwrites it.
func (c *Client) SetURL(key, url string) {
	c.urlMap[key] = url
}

// SetErrorURL is a function to set Webhook URL for error alert to the client.
// If the error alert URL is already set, this function overwrites it.
func (c *Client) SetErrorURL(url string) {
	c.errorURL = url
}
