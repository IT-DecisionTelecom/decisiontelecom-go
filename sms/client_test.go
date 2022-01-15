package sms_test

import (
	"testing"

	"github.com/IT-DecisionTelecom/decisiontelecom-go/sms"
	"github.com/jarcoal/httpmock"
)

func TestSendMessage(t *testing.T) {
	var inputData = []struct {
		response          string
		expectedMessageId sms.MessageId
		expectedError     error
	}{
		{`["msgid","31885463"]`, sms.MessageId(31885463), nil},
		{`["error","44"]`, -1, sms.Error{sms.InvalidLoginOrPassword}},
	}

	smsClient := sms.NewClient("", "")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, input := range inputData {
		t.Run(input.response, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://web.it-decision.com/ru/js/send",
				httpmock.NewStringResponder(200, input.response))

			msgId, err := smsClient.SendMessage(sms.NewMessage("", "", "", true))
			if err != input.expectedError {
				t.Errorf("FAIL. Expected error '%v', but got '%v'", input.expectedError, err.Error())
			}

			if msgId != input.expectedMessageId {
				t.Errorf("FAIL. Expected messageId '%d', but got '%d'", input.expectedMessageId, msgId)
			}
		})
	}
}

func TestGetMessageStatus(t *testing.T) {
	var inputData = []struct {
		response       string
		expectedStatus sms.MessageStatus
		expectedError  error
	}{
		{`["status","2"]`, sms.Delivered, nil},
		{`["status",""]`, sms.Unknown, nil},
		{`["error","44"]`, -1, sms.Error{sms.InvalidLoginOrPassword}},
	}

	smsClient := sms.NewClient("", "")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, input := range inputData {
		t.Run(input.response, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://web.it-decision.com/ru/js/state",
				httpmock.NewStringResponder(200, input.response))

			status, err := smsClient.GetMessageStatus(0)
			if err != input.expectedError {
				t.Errorf("FAIL. Expected error '%v', but got '%v'", input.expectedError, err.Error())
			}

			if status != input.expectedStatus {
				t.Errorf("FAIL. Expected message status '%d', but got '%d'", input.expectedStatus, status)
			}
		})
	}
}

func TestGetBalance(t *testing.T) {
	var inputData = []struct {
		response        string
		expectedBalance *sms.Balance
		expectedError   error
	}{
		{
			`["balance":"-791.8391870","credit":"1000","currency":"EUR"]`,
			&sms.Balance{BalanceAmount: -791.8391870, CreditAmount: 1000, Currency: "EUR"},
			nil,
		},
		{
			`["balance":"348.879089","credit":"-5000.509409","currency":""]`,
			&sms.Balance{BalanceAmount: 348.879089, CreditAmount: -5000.509409, Currency: ""},
			nil,
		},
		{`["error","45"]`, nil, sms.Error{sms.UserLocked}},
	}

	smsClient := sms.NewClient("", "")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, input := range inputData {
		t.Run(input.response, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://web.it-decision.com/ru/js/balance",
				httpmock.NewStringResponder(200, input.response))

			balance, err := smsClient.GetBalance()
			if err != input.expectedError {
				t.Errorf("FAIL. Expected error '%v', but got '%v'", input.expectedError, err.Error())
			}

			if balance != nil &&
				(balance.BalanceAmount != input.expectedBalance.BalanceAmount ||
					balance.CreditAmount != input.expectedBalance.CreditAmount ||
					balance.Currency != input.expectedBalance.Currency) {
				t.Errorf("FAIL. Expected balance '%+v', but got '%+v'", input.expectedBalance, balance)
			}
		})
	}
}
