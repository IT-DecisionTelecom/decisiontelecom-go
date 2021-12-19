package viber

import (
	"encoding/json"
	"fmt"
)

const baseUrl = "https://web.it-decision.com/v1/api"
const messageIdPropertyName = "message_id"

// MessageType represents a Viber message type
type MessageType uint16

const (
	TextOnly            MessageType = 106
	TextImageButton     MessageType = 108
	TextOnly2Way        MessageType = 206
	TextImageButton2Way MessageType = 208
)

// MessageSourceType represents message sending procedure.
type MessageSourceType uint16

const (
	Promotional MessageSourceType = iota + 1
	Transactional
)

// Message represents a Viber message.
type Message struct {
	// Sender is a message sender (from whom message is sent).
	Sender string
	// Receiver is a message receiver (to whom message is sent).
	Receiver string
	// Message type is a message type.
	MessageType MessageType
	// Text is a message in the UTF8 format.
	Text string
	// ImageUrl is URL of an image for promotional message with button caption and button action.
	ImageUrl string
	// ButtonCaption is a button caption in the UTF8 format.
	ButtonCaption string
	// ButtonAction is an URL for transition when the button is pressed.
	ButtonAction string
	// SourceType is a message sending procedure.
	SourceType MessageSourceType
	// CallbackUrl is an URL for message status callback.
	CallbackUrl string
	// ValidityPeriod is a life time of a message (in seconds).
	ValidityPeriod int
}

// MessageId specifies an Id of the Viber message.
type MessageId int64

// Error represents error which may occur while working with Viber messages.
type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  int    `json:"status"`
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

// MessageReceipt represents Id and status of the particular Viber message.
type MessageReceipt struct {
	MessageId MessageId     `json:"message_id"`
	Status    MessageStatus `json:"status"`
}

// Client is used to work with Viber messages.
type Client struct {
	ApiKey string
}

// NewClient creates new Viber client instance.
func NewClient(apiKey string) *Client {
	return &Client{ApiKey: apiKey}
}

// SendMessage sends Viber message
func (client *Client) SendMessage(message Message) (MessageId, error) {
	url := fmt.Sprintf("%s/send-viber", baseUrl)
	responseBody, err := MakeHttpRequest(client.ApiKey, url, message)
	if err != nil {
		return -1, err
	}

	var responseMap map[string]int64
	if err := json.Unmarshal([]byte(responseBody), &responseMap); err != nil {
		return -1, err
	}

	return MessageId(responseMap[messageIdPropertyName]), nil
}

// GetMessageStatus returns Viber message status
func (client *Client) GetMessageStatus(messageId MessageId) (*MessageReceipt, error) {
	url := fmt.Sprintf("%s/receive-viber", baseUrl)
	request := map[string]MessageId{messageIdPropertyName: messageId}

	responseBody, err := MakeHttpRequest(client.ApiKey, url, request)
	if err != nil {
		return nil, err
	}

	var messageReceipt MessageReceipt
	if err := json.Unmarshal([]byte(responseBody), &messageReceipt); err != nil {
		return nil, err
	}

	return &messageReceipt, nil
}
