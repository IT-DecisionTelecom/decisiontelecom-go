package viber

const baseUrl = "https://web.it-decision.com/v1/api"
const messageIdPropertyName = "message_id"

// MessageId specifies an Id of the Viber message.
type MessageId int64

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
	MessageId MessageId     `json:"message_id"` // Id of the Viber message which status should be got (sent in the last 5 days).
	Status    MessageStatus `json:"status"`     // Viber message status
}

// ViberClient is used to work with Viber messages.
type ViberClient struct {
	base *baseClient
}

// NewViberClient creates new Viber client instance.
func NewViberClient(apiKey string) *ViberClient {
	return &ViberClient{
		base: &baseClient{ApiKey: apiKey},
	}
}

// SendMessage sends Viber message.
func (client *ViberClient) SendMessage(message *Message) (MessageId, error) {
	return client.base.SendMessage(message)
}

// GetMessageStatus returns Viber message status.
func (client *ViberClient) GetMessageStatus(messageId MessageId) (*MessageReceipt, error) {
	messageReceipt := &MessageReceipt{}
	if err := client.base.GetMessageStatusResponse(messageId, messageReceipt); err != nil {
		return nil, err
	}

	return messageReceipt, nil
}
