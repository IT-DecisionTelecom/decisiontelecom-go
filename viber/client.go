package viber

import (
	"encoding/json"

	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber/internal"
)

// Error represents error which may occur while working with Viber messages.
type Error struct {
	Name    string `json:"name"`    // Error name
	Message string `json:"message"` // Error message
	Code    int    `json:"code"`    // Error code
	Status  int    `json:"status"`  // Error status
}

// Error implements error interface.
func (e Error) Error() string {
	return e.Name
}

// MessageStatus represents Viber message status.
type MessageStatus uint16

const (
	Sent MessageStatus = iota
	Delivered
	ErrorStatus
	Rejected
	Undelivered
	Pending
	Unknown = iota + 20
)

// String returns the message status description.
func (s MessageStatus) String() string {
	switch s {
	case Sent:
		return "Sent"
	case Delivered:
		return "Delivered"
	case ErrorStatus:
		return "Error"
	case Rejected:
		return "Rejected"
	case Undelivered:
		return "Undelivered"
	case Pending:
		return "Pending"
	case Unknown:
		return "Unknown"
	default:
		return "Invalid status"
	}
}

// MessageReceipt represents Id and status of the particular Viber message.
type MessageReceipt struct {
	MessageId int64         `json:"message_id"` // Id of the Viber message which status should be got (sent in the last 5 days).
	Status    MessageStatus `json:"status"`     // Viber message status
}

// Client is used to work with Viber messages.
type Client struct {
	base *internal.BaseClient
}

// NewClient creates new Viber client instance.
func NewClient(apiKey string) *Client {
	return &Client{
		base: &internal.BaseClient{
			ApiKey:              apiKey,
			ParseViberErrorFunc: parseViberError,
		},
	}
}

// SendMessage sends Viber message.
func (client *Client) SendMessage(message *Message) (int64, error) {
	return client.base.SendMessage(message)
}

// GetMessageStatus returns Viber message status.
func (client *Client) GetMessageStatus(messageId int64) (*MessageReceipt, error) {
	messageReceipt := &MessageReceipt{}
	if err := client.base.GetMessageStatusResponse(messageId, messageReceipt); err != nil {
		return nil, err
	}

	return messageReceipt, nil
}

func parseViberError(responseBody []byte) error {
	var viberError Error
	if err := json.Unmarshal(responseBody, &viberError); err != nil {
		return err
	}

	return viberError
}
