package slacka

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/nlopes/slack"
)

func (c *Client) send(url, message string, attachments []slack.Attachment, options ...WebhookOption) error {

	msg := slack.WebhookMessage{
		Username:    c.projectName,
		IconEmoji:   ":" + c.iconEmoji + ":",
		Text:        message,
		Attachments: attachments,
	}

	for _, opt := range options {
		opt(&msg)
	}

	if err := slack.PostWebhook(url, &msg); err != nil {
		return fmt.Errorf("Error occurred in slack.PostWebhook: %v", err)
	}

	return nil
}

// Send is a function to send message to Slack.
func (c *Client) Send(urlKey, message string, options ...WebhookOption) error {

	url, ok := c.urlMap[urlKey]
	if !ok {
		return errors.New("urlKey is not set to client")
	}

	return c.send(url, message, nil, options...)
}

func (c *Client) alert(message, severity, color string) error {

	b := &bytes.Buffer{}
	fmt.Fprintf(b, "*Message:* %s\n", message)
	fmt.Fprintf(b, "*Severity:* %s\n", severity)
	fmt.Fprintf(b, "*Timestamp:* %s\n", time.Now().Format(time.RFC3339))

	a := slack.Attachment{
		Text:  b.String(),
		Color: color,
	}

	return c.send(
		c.errorURL,
		fmt.Sprintf("Error occurred in the service `%s`.", c.serviceName),
		[]slack.Attachment{a},
	)
}

const (
	blue   = "#3AA3E3"
	green  = "#3AE37A"
	yellow = "#E3C73B"
	orange = "#F78B51"
	red    = "#E33B3B"
)

// Debug sends a debug message to Slack Error Channel.
func (c *Client) Debug(message string) error {
	return c.alert(message, "Debug", blue)
}

// Info sends a info message to Slack Error Channel.
func (c *Client) Info(message string) error {
	return c.alert(message, "Info", green)
}

// Warn sends a warning message to Slack Error Channel.
func (c *Client) Warn(message string) error {
	return c.alert(message, "Warn", yellow)
}

// Error sends a error message to Slack Error Channel.
func (c *Client) Error(message string) error {
	return c.alert(message, "Error", orange)
}

// Fatal sends a fatal message to Slack Error Channel.
func (c *Client) Fatal(message string) error {
	return c.alert(message, "Fatal", red)
}
