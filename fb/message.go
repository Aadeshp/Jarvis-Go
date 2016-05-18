package fb

import (
	"encoding/json"
	"fmt"
)

type payload interface{}

type urlPayload struct {
	payload `json:"-"`

	Url string `json:"url"`
}

type buttonPayload struct {
	payload `json:"-"`

	TemplateType string   `json:"template_type"`
	Text         string   `json:"text"`
	Buttons      []Button `json:"buttons"`
}

type Button interface{}
type UrlButton struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}

type PostbackButton struct {
	Title   string `json:"title"`
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

type attachment struct {
	Type    string  `json:"type"`
	Payload payload `json:"payload"`
}

type message struct {
	Text       string      `json:"text"`
	Attachment *attachment `json:"attachment"`
}

func newMessage() *message {
	return &message{}
}

func (this *message) SetImageUrl(imageUrl string) {
	this.Attachment = &attachment{
		Type: "image",
		Payload: &urlPayload{
			Url: imageUrl,
		},
	}
}

func (this *message) SetTextAndButtons(text string, buttons []Button) {
	this.Attachment = &attachment{
		Type: "template",
		Payload: &buttonPayload{
			TemplateType: "button",
			Text:         text,
			Buttons:      buttons,
		},
	}
}

type Recipient struct {
	Id string `json:"id"`
}

func NewRecipient(sender string) *Recipient {
	return &Recipient{
		Id: sender,
	}
}

type BaseMessage struct {
	Recipient *Recipient `json:"recipient"`
	Message   *message   `json:"message"`
}

func (this *BaseMessage) Build() {
	ret, err := json.Marshal(this)
	if err != nil {
		return
	}

	fmt.Println(string(ret))
}

func NewTextMessage(sender string, message string) *BaseMessage {
	innerMsg := newMessage()
	innerMsg.Text = message

	msg := &BaseMessage{
		Recipient: NewRecipient(sender),
		Message:   innerMsg,
	}

	return msg
}

func NewImageMessage(sender string, imageUrl string) *BaseMessage {
	innerMsg := newMessage()
	innerMsg.SetImageUrl(imageUrl)

	msg := &BaseMessage{
		Recipient: NewRecipient(sender),
		Message:   innerMsg,
	}

	return msg
}

func NewButtonsMessage(sender string, message string, buttons []Button) *BaseMessage {
	innerMsg := newMessage()
	innerMsg.SetTextAndButtons(message, buttons)

	msg := &BaseMessage{
		Recipient: NewRecipient(sender),
		Message:   innerMsg,
	}

	return msg
}
