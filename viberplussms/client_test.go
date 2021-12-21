package viberplussms_test

import (
	"testing"

	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber"
	"github.com/IT-DecisionTelecom/decisiontelecom-go/viberplussms"
	"github.com/jarcoal/httpmock"
)

func TestSendMessage(t *testing.T) {
	var inputData = []struct {
		responseStatus    int
		response          string
		expectedMessageId viber.MessageId
		expectedError     error
	}{
		{200, `{"message_id":429}`, viber.MessageId(429), nil},
		{
			200,
			`{"name":"Invalid Parameter: source_addr","message":"Empty parameter or parameter validation error","code":1,"status":400}`,
			-1,
			viber.Error{
				Name:    "Invalid Parameter: source_addr",
				Message: "Empty parameter or parameter validation error",
				Code:    1,
				Status:  400,
			},
		},
		{401, `Some response content`, -1, viber.Error{Name: "Unauthorized", Status: 401}},
	}

	client := viberplussms.NewClient("")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, input := range inputData {
		t.Run(input.response, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://web.it-decision.com/v1/api/send-viber",
				httpmock.NewStringResponder(input.responseStatus, input.response))

			msgId, err := client.SendMessage(viberplussms.Message{})
			if err != input.expectedError {
				t.Errorf("FAIL. Expected error '%+v', but got '%+v'", input.expectedError, err)
			}

			if msgId != input.expectedMessageId {
				t.Errorf("FAIL. Expected messageId '%d', but got '%d'", input.expectedMessageId, msgId)
			}
		})
	}
}

func TestGetMessageStatus(t *testing.T) {
	var inputData = []struct {
		responseStatus     int
		response           string
		expectedMsgReceipt *viberplussms.MessageReceipt
		expectedError      error
	}{
		{
			200,
			`{"message_id":429,"status":1,"sms_message_id":36478,"sms_message_status":2}`,
			&viberplussms.MessageReceipt{
				MessageId:        429,
				Status:           viber.Delivered,
				SmsMessageId:     36478,
				SmsMessageStatus: viberplussms.Delivered,
			},
			nil,
		},
		{
			200,
			`{"name":"Invalid Parameter: source_addr","message":"Empty parameter or parameter validation error","code":1,"status":400}`,
			nil,
			viber.Error{
				Name:    "Invalid Parameter: source_addr",
				Message: "Empty parameter or parameter validation error",
				Code:    1,
				Status:  400,
			},
		},
		{401, `Some response content`, nil, viber.Error{Name: "Unauthorized", Status: 401}},
	}

	client := viberplussms.NewClient("")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, input := range inputData {
		t.Run(input.response, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://web.it-decision.com/v1/api/receive-viber",
				httpmock.NewStringResponder(input.responseStatus, input.response))

			msgReceipt, err := client.GetMessageStatus(0)
			if err != input.expectedError {
				t.Errorf("FAIL. Expected error '%+v', but got '%+v'", input.expectedError, err)
			}

			if input.expectedMsgReceipt != nil &&
				(msgReceipt.MessageId != input.expectedMsgReceipt.MessageId ||
					msgReceipt.Status != input.expectedMsgReceipt.Status ||
					msgReceipt.SmsMessageId != input.expectedMsgReceipt.SmsMessageId ||
					msgReceipt.SmsMessageStatus != input.expectedMsgReceipt.SmsMessageStatus) {
				t.Errorf("FAIL. Expected message receipt '%+v', but got '%+v'", input.expectedMsgReceipt, msgReceipt)
			}
		})
	}
}