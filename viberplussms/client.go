package viberplussms

import (
	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber"
)

// Message represents a Viber plus SMS message.
type Message struct {
	viber.Message
	SmsText string `json:"text_sms"` // SmsText is an alternative SMS message text for cases when Viber message is not sent.
}

// SmsMessageStatus represents SMS message status
type SmsMessageStatus uint16

const (
	Delivered SmsMessageStatus = iota + 2
	Expired
	_
	Undeliverable
)

// String returns the SMS message status description.
func (s SmsMessageStatus) String() string {
	switch s {
	case Delivered:
		return "Delivered"
	case Expired:
		return "Expired"
	case Undeliverable:
		return "Undeliverable"
	default:
		return "Invalid status"
	}
}

// MessageReceipt represents Id and status of the particular Viber plus SMS message.
type MessageReceipt struct {
	MessageId        viber.MessageId     `json:"message_id"`         // Id of the Viber message which status should be got (sent in the last 5 days).
	Status           viber.MessageStatus `json:"status"`             // Viber message status
	SmsMessageId     int64               `json:"sms_message_id"`     // SMS message Id (if available, only for transactional messages)
	SmsMessageStatus SmsMessageStatus    `json:"sms_message_status"` // SMS message status (if available, only for transactional messages)
}

// Client is used to work with Viber plus SMS messages.
type Client struct {
	base *viber.BaseClient
}

// NewClient creates new Viber plus SMS client instance.
func NewClient(apiKey string) *Client {
	return &Client{
		base: &viber.BaseClient{ApiKey: apiKey},
	}
}

// SendMessage sends Viber plus SMS message.
func (cl *Client) SendMessage(message Message) (viber.MessageId, error) {
	return cl.base.SendMessage(message)
}

// GetMessageStatus returns Viber plus SMS message status.
func (client *Client) GetMessageStatus(messageId viber.MessageId) (*MessageReceipt, error) {
	messageReceipt := &MessageReceipt{}
	if err := client.base.GetMessageStatusResponse(messageId, messageReceipt); err != nil {
		return nil, err
	}

	return messageReceipt, nil
}
