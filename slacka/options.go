package slacka

import (
	"github.com/nlopes/slack"
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

// WebhookIconEmoji changes icon emoji of the webhook message.
func WebhookIconEmoji(iconEmoji string) WebhookOption {
	return func(msg *slack.WebhookMessage) {
		if msg == nil {
			return
		}
		msg.IconEmoji = ":" + iconEmoji + ":"
	}
}
