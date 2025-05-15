package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Bot struct {
	Client *tgbotapi.BotAPI
}

func NewBot(client *tgbotapi.BotAPI) *Bot {
	return &Bot{client}
}

func (b *Bot) Send(msg tgbotapi.Chattable) {
	if _, err := b.Client.Send(msg); err != nil {
	}
}
