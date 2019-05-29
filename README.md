# slack-alert (slacka)

[![GoDoc](https://godoc.org/github.com/takuoki/slack-alert/slacka?status.svg)](https://godoc.org/github.com/takuoki/slack-alert/slacka)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A Golang package for alerts using Slack.

## Installation

```bash
go get github.com/takuoki/slack-alert/slacka
```

## Usage

### Create New Client

```go
client := slacka.New("projectName", "serviceName", "icon_emoji")
client.SetURL("URL_KEY1", "https://hooks.slack.com/services/xxx/xxx/xxxxx")
client.SetURL("URL_KEY2", "https://hooks.slack.com/services/yyy/yyy/yyyyy")
client.SetErrorURL("https://hooks.slack.com/services/zzz/zzz/zzzzz")
```

### Send Message

```go
err := client.Send("URL_KEY", "This is message!")
```

You can change `username` and `icon_emoji` with options.

```go
err := client.Send("URL_KEY", "This is message!",
  slacka.WebhookUsername("new username"),
  slacka.WebhookIconEmoji("new_icon_emoji"))
```

### Send Error Message

Every method has a method of format type like `Debugf`.

```go
// Debug
err := client.Debug("This is message!")
```

```go
// Info
err := client.Info("This is message!")
```

```go
// Warn
err := client.Warn("This is message!")
```

```go
// Error
err := client.Error("This is message!")
```

```go
// Fatal
err := client.Fatal("This is message!")
```
