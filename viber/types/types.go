package types

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
