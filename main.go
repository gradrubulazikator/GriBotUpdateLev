package main

import (
    "log"
    "time"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var updateMessage = "Ваше регулярное обновление! Проверьте статус ваших задач и целей."

func main() {
    bot, err := tgbotapi.NewBotAPI("7666734298:AAHHxhGhWVM7_hKGkJ6q2R4IV1p4Rw9Y6PA")
    if err != nil {
        log.Fatalf("Ошибка при подключении к боту: %v", err)
    }
    bot.Debug = true

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates, _ := bot.GetUpdatesChan(u)

    // ID чата пользователя, которому будут приходить регулярные обновления
    chatID := int64(6733740743)

    ticker := time.NewTicker(4 * time.Hour) // обновления каждые 4 часа
    defer ticker.Stop()

    log.Println("Бот запущен и ожидает команд...")

    go func() {
        for {
            select {
            case <-ticker.C:
                msg := tgbotapi.NewMessage(chatID, updateMessage)
                if _, err := bot.Send(msg); err != nil {
                    log.Printf("Ошибка при отправке сообщения: %v", err)
                } else {
                    log.Println("Сообщение успешно отправлено")
                }
            }
        }
    }()

    // Обработка команды /start
    for update := range updates {
        if update.Message == nil { // игнорируем не-текстовые сообщения
            continue
        }

        if update.Message.Text == "/start" {
            startMessage := "Добро пожаловать! Я буду отправлять вам регулярные обновления. Ожидайте первого сообщения!"
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, startMessage)
            if _, err := bot.Send(msg); err != nil {
                log.Printf("Ошибка при отправке сообщения: %v", err)
            }
        }
    }
}

