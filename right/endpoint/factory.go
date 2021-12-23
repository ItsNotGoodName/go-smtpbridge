package endpoint

import (
	"fmt"

	"github.com/ItsNotGoodName/smtpbridge/app"
)

func Factory(endpointType string, config map[string]string) (app.EndpointPort, error) {
	switch endpointType {
	case "telegram":
		return telegramFactory(config)
	case "mock":
		return NewMock()
	}

	return nil, fmt.Errorf("unknown endpointType %s: %v", endpointType, app.ErrInvalidEndpointType)
}

func telegramFactory(config map[string]string) (*Telegram, error) {
	token, ok := config["token"]
	if !ok {
		return nil, fmt.Errorf("telegram token not found: %v", app.ErrInvalidEndpointConfig)
	}
	chatID, ok := config["chat_id"]
	if !ok {
		return nil, fmt.Errorf("telegram chat_id not found: %v", app.ErrInvalidEndpointConfig)
	}
	return NewTelegram(token, chatID), nil
}