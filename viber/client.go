package viber

const baseUrl = "https://web.it-decision.com/v1/api"
const messageIdPropertyName = "message_id"

// MessageType represents a Viber message type.
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
	Sender         string            `json:"source_addr"`      // Sender is a message sender (from whom message is sent).
	Receiver       string            `json:"destination_addr"` // Receiver is a message receiver (to whom message is sent).
	MessageType    MessageType       `json:"message_type"`     // Message type is a message type.
	Text           string            `json:"text"`             // Text is a message in the UTF8 format.
	ImageUrl       string            `json:"image"`            // ImageUrl is URL of an image for promotional message with button caption and button action.
	ButtonCaption  string            `json:"button_caption"`   // ButtonCaption is a button caption in the UTF8 format.
	ButtonAction   string            `json:"button_action"`    // ButtonAction is an URL for transition when the button is pressed.
	SourceType     MessageSourceType `json:"source_type"`      // SourceType is a message sending procedure.
	CallbackUrl    string            `json:"callback_url"`     // CallbackUrl is an URL for message status callback.
	ValidityPeriod int               `json:"validity_period"`  // ValidityPeriod is a life time of a message (in seconds).
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
	MessageId MessageId     `json:"message_id"`
	Status    MessageStatus `json:"status"`
}

// Client is used to work with Viber messages.
type Client struct {
	base *BaseClient
}

// NewClient creates new Viber client instance.
func NewClient(apiKey string) *Client {
	return &Client{
		base: &BaseClient{ApiKey: apiKey},
	}
}

// SendMessage sends Viber message.
func (client *Client) SendMessage(message Message) (MessageId, error) {
	return client.base.SendMessage(message)
}

// GetMessageStatus returns Viber message status.
func (client *Client) GetMessageStatus(messageId MessageId) (*MessageReceipt, error) {
	messageReceipt := &MessageReceipt{}
	if err := client.base.GetMessageStatusResponse(messageId, messageReceipt); err != nil {
		return nil, err
	}

	return messageReceipt, nil
}
