package bot

import (
	"TormentaBot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateMainMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📜 Magias", "/magias"),
			tgbotapi.NewInlineKeyboardButtonData("⚡ Poderes", "/poderes"),
		),

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📚 Regras", "/regras"),
			tgbotapi.NewInlineKeyboardButtonData("🛡️ Condições", "/condicoes"),
		),
	)

}

func CreateSpellsListKeyboard(magias []models.Magia) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, magia := range magias {
		btn := tgbotapi.NewInlineKeyboardButtonData(magia.Nome, "/magia_"+magia.Nome)
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(btn))
	}

	rows = append(rows, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("↩️ Voltar ao Menu", "/menu"),
	))

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func CreateBackKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("↩️ Voltar ao Menu", "/menu"),
		),
	)
}
