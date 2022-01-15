package viber

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

// MessageWithSms represents a Viber plus SMS message.
type MessageWithSms struct {
	Message
	SmsText string `json:"text_sms"` // SmsText is an alternative SMS message text for cases when Viber message is not sent.
}

// NewMessage creates new Message.
func NewMessage() *Message {
	return &Message{}
}

// NewMessageWithSms creates new MessageWithSms.
func NewMessageWithSms() *MessageWithSms {
	return &MessageWithSms{}
}

// SetSender sets message sender.
func (m *Message) SetSender(sender string) *Message {
	m.Sender = sender
	return m
}

// SetReceiver sets message receiver.
func (m *Message) SetReceiver(receiver string) *Message {
	m.Receiver = receiver
	return m
}

// SetMessageType sets message type.
func (m *Message) SetMessageType(messageType MessageType) *Message {
	m.MessageType = messageType
	return m
}

// SetText sets message text.
func (m *Message) SetText(text string) *Message {
	m.Text = text
	return m
}

// SetImageUrl sets image url for promotional message with button caption and button action.
func (m *Message) SetImageUrl(imageUrl string) *Message {
	m.ImageUrl = imageUrl
	return m
}

// SetButtonCaption sets button caption.
func (m *Message) SetButtonCaption(buttonCaption string) *Message {
	m.ButtonCaption = buttonCaption
	return m
}

// SetButtonAction sets button action (an URL for transition when the button is pressed).
func (m *Message) SetButtonAction(buttonAction string) *Message {
	m.ButtonAction = buttonAction
	return m
}

// SetSourceType sets message source type (sending procedure).
func (m *Message) SetSourceType(sourceType MessageSourceType) *Message {
	m.SourceType = sourceType
	return m
}

// SetCallbackUrl sets callback URL (an URL for message status callback).
func (m *Message) SetCallbackUrl(callbackUrl string) *Message {
	m.CallbackUrl = callbackUrl
	return m
}

// SetValidityPeriod sets message validity period (life time of a message, in seconds).
func (m *Message) SetValidityPeriod(validityPeriod int) *Message {
	m.ValidityPeriod = validityPeriod
	return m
}

// AddSmsText adds SMS text to the message (alternative SMS message text for cases when Viber message is not sent).
func (m *Message) AddSmsText(smsText string) *MessageWithSms {
	msgSms := MessageWithSms{
		Message: *m,
	}
	msgSms.SmsText = smsText
	return &msgSms
}
