package viber_test

import (
	"testing"

	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber"
	"github.com/jarcoal/httpmock"
)

func TestSendViberPlusSmsMessage(t *testing.T) {
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

	client := viber.NewViberPlusSmsClient("")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, input := range inputData {
		t.Run(input.response, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://web.it-decision.com/v1/api/send-viber",
				httpmock.NewStringResponder(input.responseStatus, input.response))

			msgId, err := client.SendMessage(viber.NewMessageWithSms())
			if err != input.expectedError {
				t.Errorf("FAIL. Expected error '%+v', but got '%+v'", input.expectedError, err)
			}

			if msgId != input.expectedMessageId {
				t.Errorf("FAIL. Expected messageId '%d', but got '%d'", input.expectedMessageId, msgId)
			}
		})
	}
}

func TestGetViberPlusSmsMessageStatus(t *testing.T) {
	var inputData = []struct {
		responseStatus     int
		response           string
		expectedMsgReceipt *viber.MessageReceiptWithSms
		expectedError      error
	}{
		{
			200,
			`{"message_id":429,"status":1,"sms_message_id":36478,"sms_message_status":2}`,
			&viber.MessageReceiptWithSms{
				MessageId:        429,
				Status:           viber.Delivered,
				SmsMessageId:     36478,
				SmsMessageStatus: viber.SmsDelivered,
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

	client := viber.NewViberPlusSmsClient("")

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
