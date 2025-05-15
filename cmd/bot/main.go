package main

import (
	"TormentaBot/internal/bot"
	"TormentaBot/internal/config"
	"TormentaBot/internal/data"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	cfg := config.Load()

	if err := data.LoadSpells(); err != nil {
		log.Panicf("Falha ao carregar magias: %v", err) // Mostra o erro detalhado
	}

	tbotClient, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic("Erro ao iniciar o bot: ", err)
	}

	tbotClient.Debug = cfg.Debug

	myBot := bot.NewBot(tbotClient)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := tbotClient.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			myBot.HandleMessages(update)
		} else if update.CallbackQuery != nil {
			myBot.HandleCallbacks(update.CallbackQuery)
		}
	}
}
