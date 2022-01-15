package viber

// MessageWithSms represents a Viber plus SMS message.
type MessageWithSms struct {
	Message
	SmsText string `json:"text_sms"` // SmsText is an alternative SMS message text for cases when Viber message is not sent.
}

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

// MessageReceiptWithSms represents Id and status of the particular Viber plus SMS message.
type MessageReceiptWithSms struct {
	MessageId        MessageId        `json:"message_id"`         // Id of the Viber message which status should be got (sent in the last 5 days).
	Status           MessageStatus    `json:"status"`             // Viber message status
	SmsMessageId     int64            `json:"sms_message_id"`     // SMS message Id (if available, only for transactional messages)
	SmsMessageStatus SmsMessageStatus `json:"sms_message_status"` // SMS message status (if available, only for transactional messages)
}

// ViberPlusSmsClient is used to work with Viber plus SMS messages.
type ViberPlusSmsClient struct {
	base *baseClient
}

// NewViberPlusSmsClient creates new Viber plus SMS client instance.
func NewViberPlusSmsClient(apiKey string) *ViberPlusSmsClient {
	return &ViberPlusSmsClient{
		base: &baseClient{ApiKey: apiKey},
	}
}

// SendMessage sends Viber plus SMS message.
func (cl *ViberPlusSmsClient) SendMessage(message MessageWithSms) (MessageId, error) {
	return cl.base.SendMessage(message)
}

// GetMessageStatus returns Viber plus SMS message status.
func (client *ViberPlusSmsClient) GetMessageStatus(messageId MessageId) (*MessageReceiptWithSms, error) {
	messageReceipt := &MessageReceiptWithSms{}
	if err := client.base.GetMessageStatusResponse(messageId, messageReceipt); err != nil {
		return nil, err
	}

	return messageReceipt, nil
}
