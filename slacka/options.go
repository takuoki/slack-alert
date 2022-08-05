package slacka

import (
	"github.com/slack-go/slack"
)

// WebhookOption changes some parameters of the webhook message.
type WebhookOption func(*slack.WebhookMessage)

// WebhookUsername changes username of the webhook message.
func WebhookUsername(username string) WebhookOption {
	return func(msg *slack.WebhookMessage) {
		if msg == nil {
			return
		}
		msg.Username = username
	}
}

// WebhookIconURL changes icon url of the webhook message.
func WebhookIconURL(iconURL string) WebhookOption {
	return func(msg *slack.WebhookMessage) {
		if msg == nil {
			return
		}
		msg.IconURL = iconURL
	}
}
