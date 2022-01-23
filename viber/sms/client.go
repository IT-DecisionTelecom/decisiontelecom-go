package sms

import (
	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber"
	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber/internal"
	types "github.com/IT-DecisionTelecom/decisiontelecom-go/viber/types"
)

// SmsMessageStatus represents SMS message status
type SmsMessageStatus uint16

const (
	SmsDelivered SmsMessageStatus = iota + 2
	SmsExpired
	_
	SmsUndeliverable
)

// String returns the SMS message status description.
func (s SmsMessageStatus) String() string {
	switch s {
	case SmsDelivered:
		return "Delivered"
	case SmsExpired:
		return "Expired"
	case SmsUndeliverable:
		return "Undeliverable"
	default:
		return "Invalid status"
	}
}

// MessageReceipt represents Id and status of the particular Viber plus SMS message.
type MessageReceipt struct {
	MessageId        types.MessageId     `json:"message_id"`         // Id of the Viber message which status should be got (sent in the last 5 days).
	Status           types.MessageStatus `json:"status"`             // Viber message status
	SmsMessageId     int64               `json:"sms_message_id"`     // SMS message Id (if available, only for transactional messages)
	SmsMessageStatus SmsMessageStatus    `json:"sms_message_status"` // SMS message status (if available, only for transactional messages)
}

// Message represents a Viber plus SMS message.
type Message struct {
	viber.Message
	SmsText string `json:"text_sms"` // SmsText is an alternative SMS message text for cases when Viber message is not sent.
}

// NewMessage creates new MessageWithSms.
func NewMessage() *Message {
	return &Message{}
}

// SetSmsText sets message SMS text (alternative SMS message text for cases when Viber message is not sent).
func (m *Message) SetSmsText(smsText string) {
	m.SmsText = smsText
}

// Client is used to work with Viber plus SMS messages.
type Client struct {
	base *internal.BaseClient
}

// NewClient creates new Viber plus SMS client instance.
func NewClient(apiKey string) *Client {
	return &Client{
		base: &internal.BaseClient{ApiKey: apiKey},
	}
}

// SendMessage sends Viber plus SMS message.
func (cl *Client) SendMessage(message *Message) (types.MessageId, error) {
	return cl.base.SendMessage(message)
}

// GetMessageStatus returns Viber plus SMS message status.
func (client *Client) GetMessageStatus(messageId types.MessageId) (*MessageReceipt, error) {
	messageReceipt := &MessageReceipt{}
	if err := client.base.GetMessageStatusResponse(messageId, messageReceipt); err != nil {
		return nil, err
	}

	return messageReceipt, nil
}
