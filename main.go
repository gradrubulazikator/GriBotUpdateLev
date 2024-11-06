package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"GriBotUpdateLev/internal/config"
	"GriBotUpdateLev/internal/updater"
)

func main() {
	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	bot.Debug = true // включить режим отладки, если нужно
	log.Println("Запуск GriBotUpdateLev...")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	// Обработка обновлений
	for update := range updates {
		updater.ProcessUpdates(bot, update)
	}
}

