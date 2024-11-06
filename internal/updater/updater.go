package updater

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"GriBotUpdateLev/internal/config"
)

// Функция отправки сообщения об обновлении
func SendUpdate(bot *tgbotapi.BotAPI, message string) {
	msg := tgbotapi.NewMessage(config.ChatID, message)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка отправки обновления: %v", err)
	}
}

// Основной функционал для обработки обновлений
func ProcessUpdates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Логика для обработки разных типов сообщений
	if update.Message != nil {
		SendUpdate(bot, "Новое сообщение: "+update.Message.Text)
	}
	// Добавьте дополнительную обработку, если потребуется
}

