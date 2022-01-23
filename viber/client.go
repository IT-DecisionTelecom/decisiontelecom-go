package viber

import (
	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber/internal"
	types "github.com/IT-DecisionTelecom/decisiontelecom-go/viber/types"
)

// MessageReceipt represents Id and status of the particular Viber message.
type MessageReceipt struct {
	MessageId types.MessageId     `json:"message_id"` // Id of the Viber message which status should be got (sent in the last 5 days).
	Status    types.MessageStatus `json:"status"`     // Viber message status
}

// Client is used to work with Viber messages.
type Client struct {
	base *internal.BaseClient
}

// NewClient creates new Viber client instance.
func NewClient(apiKey string) *Client {
	return &Client{
		base: &internal.BaseClient{ApiKey: apiKey},
	}
}

// SendMessage sends Viber message.
func (client *Client) SendMessage(message *Message) (types.MessageId, error) {
	return client.base.SendMessage(message)
}

// GetMessageStatus returns Viber message status.
func (client *Client) GetMessageStatus(messageId types.MessageId) (*MessageReceipt, error) {
	messageReceipt := &MessageReceipt{}
	if err := client.base.GetMessageStatusResponse(messageId, messageReceipt); err != nil {
		return nil, err
	}

	return messageReceipt, nil
}
