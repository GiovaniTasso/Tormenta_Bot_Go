package bot

import (
	"TormentaBot/internal/data"
	"TormentaBot/internal/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const (
	StateWaitingForSpellSearch = "waiting_for_spell_search"
)

func (b *Bot) SendSearchPrompt(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "🔍 Digite o nome da magia que deseja buscar:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️ Voltar", "/menu"),
		),
	)
	b.Send(msg)
}

func (b *Bot) HandleMessages(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Text == "/magias" || update.Message.Command() == "magias" {
		b.SendSearchPrompt(update.Message.Chat.ID)
		return
	}

	if update.Message.Text == "/menu" {
		b.SendMainMenu(update.Message.Chat.ID)
	}

	results := data.SearchSpells(update.Message.Text)
	if len(results) > 0 {
		b.SendSpellsList(update.Message.Chat.ID, results)
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🔍 Nenhuma magia encontrada. Tente novamente!")
		b.Send(msg)
	}

}

func (b *Bot) SendSpellsList(chatID int64, magias []models.Magia) {
	msg := tgbotapi.NewMessage(chatID, "📚 Magias encontradas:")
	msg.ReplyMarkup = CreateSpellsListKeyboard(magias)
	b.Send(msg)
}

func (b *Bot) HandleCallbacks(query *tgbotapi.CallbackQuery) {
	callbackData := query.Data

	switch {
	case callbackData == "/menu":
		b.SendMainMenu(query.Message.Chat.ID)

	case callbackData == "/magias":
		b.SendSearchPrompt(query.Message.Chat.ID)

	case strings.HasPrefix(callbackData, "/magia_"):
		magiaNome := strings.TrimPrefix(callbackData, "/magia_")
		magia, found := data.GetSpellByName(magiaNome)
		if !found {
			b.Send(tgbotapi.NewMessage(query.Message.Chat.ID, "❌ Magia não encontrada!"))
			return
		}
		b.SendSpellDetails(query.Message.Chat.ID, magia)

	}
}

func (b *Bot) SendSpellDetails(chatID int64, magia models.Magia) {
	text := fmt.Sprintf(
		"✨ *%s* (Nível %d)\n"+
			"🏫 _%s_\n"+
			"🔮 *Tipo:* %s\n\n"+
			"📝 *Execução:* %s\n"+
			"🎯 *Alcance:* %s\n"+
			"🔮 *Alvo:* %s\n"+
			"⏳ *Duração:* %s\n"+
			"🛡 *Resistência:* %s\n\n"+
			"%s",
		magia.Nome,
		magia.Nivel,
		magia.Escola,
		magia.Tipo,
		magia.Execucao,
		magia.Alcance,
		magia.Alvo,
		magia.Duracao,
		magia.Resistencia,
		magia.Descricao,
	)

	if len(magia.Aprimoramentos) > 0 {
		text += "\n\n🔧 *Aprimoramentos:*"
		for _, apr := range magia.Aprimoramentos {
			text += fmt.Sprintf("\n- *%s*: %s", apr.Custo, apr.Descricao)
		}
	}

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️ Voltar", "/magias"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🏠 Menu Principal", "/menu"),
		),
	)
	b.Send(msg)
}

func (b *Bot) SendMainMenu(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "🎲 **Menu Principal**")
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = CreateMainMenu()
	b.Send(msg)
}
