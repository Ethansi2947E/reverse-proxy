// telegram_bot.go
package telegrambot

import (
	"fmt"
	"net/http"
	"net/url"
)

// Bot represents a Telegram bot.
type Bot struct {
	token string
}

// NewBot creates a new instance of a Telegram bot with the provided token.
func NewBot(token string) *Bot {
	return &Bot{
		token: token,
	}
}

// SendMessage sends a message to a chat in Telegram.
func (b *Bot) SendMessage(chatID int64, text string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.token)
	values := url.Values{}
	values.Set("chat_id", fmt.Sprintf("%d", chatID))
	values.Set("text", text)

	resp, err := http.PostForm(endpoint, values)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return nil
}
