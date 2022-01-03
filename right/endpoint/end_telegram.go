package endpoint

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/ItsNotGoodName/smtpbridge/core"
)

type Telegram struct {
	Token  string
	ChatID string
}

func NewTelegram(token string, chatID string) *Telegram {
	return &Telegram{
		Token:  token,
		ChatID: chatID,
	}
}

type TelegramResponse struct {
	OK          bool   `json:"ok"`
	Description string `json:"description"`
}

func (t *Telegram) Send(msg *core.EndpointMessage) error {
	if len(msg.Attachments) == 0 {
		return t.sendMessage(msg.Text)
	}

	// TODO: use sendMediaGroup when more than 1 attachment
	for _, attachment := range msg.Attachments {
		err := t.sendPhoto(msg.Text, attachment.Name, attachment.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Telegram) sendMessage(text string) error {
	// Create and send request
	if len(text) > 4096 {
		text = text[:4096]
	}
	resp, err := http.PostForm(
		"https://api.telegram.org/bot"+t.Token+"/sendMessage",
		url.Values{
			"chat_id": {t.ChatID}, "text": {text},
		})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Parse response
	res := &TelegramResponse{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}
	if !res.OK {
		return errors.New(res.Description)
	}

	return nil
}

func (t *Telegram) sendPhoto(caption, name string, photo []byte) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Photo
	w, err := writer.CreateFormFile("photo", name)
	if err != nil {
		return err
	}
	w.Write(photo)

	// Caption
	w, err = writer.CreateFormField("caption")
	if err != nil {
		return err
	}
	if len(caption) > 1024 {
		caption = caption[:1024]
	}
	w.Write([]byte(caption))
	writer.Close()

	// Create request
	req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+t.Token+"/sendPhoto?chat_id="+t.ChatID, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Parse response
	res := &TelegramResponse{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}
	if !res.OK {
		return errors.New(res.Description)
	}

	return nil
}
