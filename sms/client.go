// Package sms contains types and functions for sending SMS messages.
package sms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const baseUrl = "https://web.it-decision.com/ru/js"

// Message represents an SMS message.
type Message struct {
	ReceiverPhone string // ReceiverPhone is a message receiver phone number (MSISDN Destination)
	Sender        string // Sender is a message sender. Could be a mobile phone number (including a country code) or an alphanumeric string.
	Text          string // Text is a message body.
	Delivery      bool   // Delivery should be true if a caller needs to obtain the delivery receipt in the future (by message id).
}

// NewMessage creates new message
func NewMessage(receiver string, sender string, text string, delivery bool) *Message {
	return &Message{
		ReceiverPhone: receiver,
		Sender:        sender,
		Text:          text,
		Delivery:      delivery,
	}
}

// Balance represents user money balance.
type Balance struct {
	BalanceAmount float64 `json:"balance"`
	CreditAmount  float64 `json:"credit"`
	Currency      string  `json:"currency"`
}

// ErrorCode represents an error code of the SMS operation.
type ErrorCode int

const (
	InvalidNumber ErrorCode = iota + 40
	IncorrectSender
	InvalidMessageId
	IncorrectJson
	InvalidLoginOrPassword
	UserLocked
	EmptyText
	EmptyLogin
	EmptyPassword
	NotEnoughMoney
	AuthorizationError
	InvalidPhoneNumber
)

// String returns the error code description.
func (code ErrorCode) String() string {
	errors := []string{
		"InvalidNumber",
		"IncorrectSender",
		"InvalidMessageId",
		"IncorrectJson",
		"InvalidLoginOrPassword",
		"UserLocked",
		"EmptyText",
		"EmptyLogin",
		"EmptyPassword",
		"NotEnoughMoney",
		"AuthorizationError",
		"InvalidPhoneNumber",
	}
	if int(code-40) < len(errors) {
		return errors[code-40]
	}

	return fmt.Sprintf("Unknown error code: %d", code)
}

// MessageStatus specifies status of the SMS message in the system.
type MessageStatus int

const (
	Unknown MessageStatus = iota
	_
	Delivered
	Expired
	_
	Undeliverable
	Accepted
)

// String returns the message status description.
func (s MessageStatus) String() string {
	switch s {
	case Unknown:
		return "Unknown"
	case Delivered:
		return "Delivered"
	case Undeliverable:
		return "Undeliverable"
	case Accepted:
		return "Accepted"
	default:
		return "Invalid status"
	}
}

// Error represents error which may occur while working with SMS messages.
// It holds SMS error code.
type Error struct {
	Code ErrorCode
}

// Error implements error interface.
func (e Error) Error() string {
	return e.Code.String()
}

type emptyValueFunction func() (int64, error)

// Client is used to work with SMS messages.
type Client struct {
	Login    string
	Password string
}

// NewClient creates new SMS client instance.
func NewClient(login string, password string) *Client {
	return &Client{Login: login, Password: password}
}

// SendMessage sends SMS message.
func (client *Client) SendMessage(message *Message) (int64, error) {
	var dlr = 0
	if message.Delivery {
		dlr = 1
	}
	url := fmt.Sprintf("%s/send?login=%s&password=%s&phone=%s&sender=%s&text=%s&dlr=%d",
		baseUrl, client.Login, client.Password, url.QueryEscape(message.ReceiverPhone), url.QueryEscape(message.Sender), url.QueryEscape(message.Text), dlr)

	responseBody, err := makeHttpRequest(url)
	if err != nil {
		return -1, err
	}

	return getIntValueFromListResponseBody(responseBody, "msgid", nil)
}

// GetMessageStatus returns SMS message delivery status.
func (smsClient *Client) GetMessageStatus(messageId int64) (MessageStatus, error) {
	url := fmt.Sprintf("%s/state?login=%s&password=%s&msgid=%d", baseUrl, smsClient.Login, smsClient.Password, messageId)

	responseBody, err := makeHttpRequest(url)
	if err != nil {
		return -1, err
	}

	emptyValueFunc := func() (int64, error) {
		return int64(Unknown), nil
	}

	status, err := getIntValueFromListResponseBody(responseBody, "status", emptyValueFunc)
	if err != nil {
		return -1, err
	}

	return MessageStatus(status), nil
}

// GetBalance returns user balance information.
func (smsClient *Client) GetBalance() (*Balance, error) {
	url := fmt.Sprintf("%s/balance?login=%s&password=%s", baseUrl, smsClient.Login, smsClient.Password)

	responseBody, err := makeHttpRequest(url)
	if err != nil {
		return nil, err
	}

	// replace symbols in the response body so it's possible to parse it as json
	// regexp removes quotation marks ("") around the numbers, so they could be parsed as float
	regex := regexp.MustCompile(`"([-+]?[0-9]*\.?[0-9]+)"`)
	relplacedContent := strings.ReplaceAll(strings.ReplaceAll(responseBody, "[", "{"), "]", "}")
	relplacedContent = regex.ReplaceAllString(relplacedContent, "$1")

	var balance Balance
	if err := json.Unmarshal([]byte(relplacedContent), &balance); err != nil {
		return nil, err
	}

	return &balance, nil
}

func makeHttpRequest(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Process unsuccessful status codes
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf("an error occurred while processing request. Response code: %d (%s)",
			response.StatusCode, http.StatusText(response.StatusCode))
	}

	if strings.Contains(string(bodyBytes), "error") {
		errorCode, err := getIntValueFromListResponseBody(string(bodyBytes), "error", nil)
		if err != nil {
			return "", err
		}

		return "", Error{Code: ErrorCode(errorCode)}
	}

	return string(bodyBytes), nil
}

func getIntValueFromListResponseBody(responseBody string, keyProperty string, emptyValueFunc emptyValueFunction) (int64, error) {
	var processedSplit []string
	split := strings.Split(strings.Trim(strings.Trim(responseBody, "["), "]"), ",")
	for _, s := range split {
		processedSplit = append(processedSplit, strings.Trim(s, "\""))
	}

	if strings.Compare(processedSplit[0], keyProperty) != 0 {
		return -1, fmt.Errorf("invalid response: unknown key '%s'", processedSplit[0])
	}

	if processedSplit[1] == "" && emptyValueFunc != nil {
		return emptyValueFunc()
	} else {
		return strconv.ParseInt(processedSplit[1], 10, 64)
	}
}
