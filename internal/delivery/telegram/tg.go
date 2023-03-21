package telegram

import (
	"context"
	"fmt"
	"test-tasks/tg-bot/internal/entity"
	"test-tasks/tg-bot/internal/usecase"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	clientUseCase  usecase.ClientUseCase
	messageUseCase usecase.MessageUseCase
	log            *zap.Logger
	statUsecase    usecase.Stat
	ctxTimeout     time.Duration
	weatherUseCaes usecase.Weather
}

func New(b *tgbotapi.BotAPI, clientUseCase usecase.ClientUseCase, messageUseCase usecase.MessageUseCase, log *zap.Logger, statUsecase usecase.Stat, ctxTimeout time.Duration, weatherUseCaes usecase.Weather) *Bot {
	b.Debug = false
	return &Bot{
		bot:            b,
		clientUseCase:  clientUseCase,
		messageUseCase: messageUseCase,
		log:            log,
		statUsecase:    statUsecase,
		ctxTimeout:     ctxTimeout,
		weatherUseCaes: weatherUseCaes,
	}
}

func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg, err := b.HandleMessage(update)
			if err != nil {
				b.log.Error("error b.HandleMessage", zap.Error(err))
			}

			_, err = b.bot.Send(msg)
			if err != nil {
				b.log.Error("error bot Send", zap.Error(err))
			}
		}
	}
}

func (b *Bot) SaveInfo(ctx context.Context, m *tgbotapi.Message) error {
	client := &entity.Client{
		ClientId:  m.Chat.ID,
		FirstName: m.Chat.FirstName,
		LastName:  m.Chat.LastName,
		Username:  m.Chat.UserName,
	}

	client, err := b.clientUseCase.CreateOrUpdate(ctx, client)
	if err != nil {
		return err
	}
	err = b.messageUseCase.Create(ctx, &entity.Message{
		ClientId: client.ClientId,
		Text:     m.Text,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) HandleMessage(u tgbotapi.Update) (msg tgbotapi.MessageConfig, err error) {
	msg = tgbotapi.NewMessage(u.Message.Chat.ID, "")

	ctx, cancel := context.WithTimeout(context.Background(), b.ctxTimeout)
	defer cancel()
	defer func() {
		if err != nil {
			msg.Text = "internal error, " + err.Error()
		}
	}()

	err = b.SaveInfo(ctx, u.Message)
	if err != nil {
		b.log.Error("error save info", zap.Error(err))
		return
	}

	switch u.Message.Text {
	case "/start":
		msg, err = b.setupKeyboard(u)
	case "Статистика":
		msg, err = b.GetStats(ctx, u)
	default:
		msg, err = b.GetData(ctx, u)
	}

	return
}

func (b *Bot) GetStats(ctx context.Context, u tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "")

	stat, err := b.statUsecase.GetAllStats(ctx, u.Message.Chat.ID)
	if err != nil {
		return msg, err
	}

	msg.Text = fmt.Sprintf(`
Статистика:
Количество запросов: %d
Время первого запроса: %s
Время последнего запроса: %s`, stat.CountRequests, stat.FirstRequest, stat.LastRequest)

	return msg, nil
}

func (b *Bot) GetData(ctx context.Context, u tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "")

	w, err := b.weatherUseCaes.CurrentWeather(ctx, u.Message.Text)
	if err != nil {
		return msg, err
	}

	msg.Text = fmt.Sprintf(`
Локация:
    %s, %s
    %s
Погода:
    Температура по цельсию: %.2f
    По ощущениям: %.2f
    Описание: %s
    Скорость ветра: %.2f
    Влажность: %d
    Облачность: %d
    Порывы ветра: %.2f
	`, w.Country, w.Name, w.Time, w.TemperatureC, w.FeelslikeC, w.ConditionText, w.WindKph, w.Humidity, w.Cloud, w.GustKph)
	return msg, nil
}

func (b *Bot) setupKeyboard(u tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Тестовое задание")
	menu := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Статистика")),
	)
	msg.ReplyMarkup = menu
	return msg, nil
}
