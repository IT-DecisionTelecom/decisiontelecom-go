package viberplussms

import (
	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber"
)

// Message represents a Viber plus SMS message.
type Message struct {
	viber.Message
	SmsText string
}

type SmsMessageStatus uint16

const (
	Delivered SmsMessageStatus = iota + 2
	Expired
	_
	Undeliverable
)

type MessageReceipt struct {
	MessageId        viber.MessageId     `json:"message_id"`
	Status           viber.MessageStatus `json:"status"`
	SmsMessageId     int64               `json:"sms_message_id"`
	SmsMessageStatus SmsMessageStatus    `json:"sms_message_status"`
}

// Client is used to work with Viber plus SMS messages.
type Client struct {
	base *viber.BaseClient
}

// NewClient creates new Viber client instance.
func NewClient(apiKey string) *Client {
	return &Client{
		base: &viber.BaseClient{ApiKey: apiKey},
	}
}

func (cl *Client) SendMessage(message Message) (viber.MessageId, error) {
	return cl.base.SendMessage(message)
}

func (client *Client) GetMessageStatus(messageId viber.MessageId) (*MessageReceipt, error) {
	messageReceipt := &MessageReceipt{}
	if err := client.base.GetMessageStatusResponse(messageId, messageReceipt); err != nil {
		return nil, err
	}

	return messageReceipt, nil
}
