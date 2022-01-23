package viber_test

import (
	"fmt"
	"testing"

	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber"
	types "github.com/IT-DecisionTelecom/decisiontelecom-go/viber/types"
	"github.com/jarcoal/httpmock"
)

func TestSendViberMessage(t *testing.T) {
	var inputData = []struct {
		responseStatus    int
		response          string
		expectedMessageId int64
		expectedError     error
	}{
		{200, `{"message_id":429}`, 429, nil},
		{
			200,
			`{"name":"Invalid Parameter: source_addr","message":"Empty parameter or parameter validation error","code":1,"status":400}`,
			-1,
			types.Error{
				Name:    "Invalid Parameter: source_addr",
				Message: "Empty parameter or parameter validation error",
				Code:    1,
				Status:  400,
			},
		},
		{401, `Some response content`, -1, fmt.Errorf("an error occurred while processing request. Response code: 401 (Unauthorized)")},
	}

	client := viber.NewClient("")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, input := range inputData {
		t.Run(input.response, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://web.it-decision.com/v1/api/send-viber",
				httpmock.NewStringResponder(input.responseStatus, input.response))

			msgId, err := client.SendMessage(viber.NewMessage())
			if err != nil && err.Error() != input.expectedError.Error() {
				t.Errorf("FAIL. Expected error '%+v', but got '%+v'", input.expectedError, err)
			}

			if msgId != input.expectedMessageId {
				t.Errorf("FAIL. Expected messageId '%d', but got '%d'", input.expectedMessageId, msgId)
			}
		})
	}
}

func TestGetViberMessageStatus(t *testing.T) {
	var inputData = []struct {
		responseStatus     int
		response           string
		expectedMsgReceipt *viber.MessageReceipt
		expectedError      error
	}{
		{200, `{"message_id":429,"status":1}`, &viber.MessageReceipt{MessageId: 429, Status: types.Delivered}, nil},
		{
			200,
			`{"name":"Invalid Parameter: source_addr","message":"Empty parameter or parameter validation error","code":1,"status":400}`,
			nil,
			types.Error{
				Name:    "Invalid Parameter: source_addr",
				Message: "Empty parameter or parameter validation error",
				Code:    1,
				Status:  400,
			},
		},
		{401, `Some response content`, nil, fmt.Errorf("an error occurred while processing request. Response code: 401 (Unauthorized)")},
	}

	client := viber.NewClient("")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, input := range inputData {
		t.Run(input.response, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://web.it-decision.com/v1/api/receive-viber",
				httpmock.NewStringResponder(input.responseStatus, input.response))

			msgReceipt, err := client.GetMessageStatus(0)
			if err != nil && err.Error() != input.expectedError.Error() {
				t.Errorf("FAIL. Expected error '%+v', but got '%+v'", input.expectedError, err)
			}

			if input.expectedMsgReceipt != nil && (msgReceipt.MessageId != input.expectedMsgReceipt.MessageId || msgReceipt.Status != input.expectedMsgReceipt.Status) {
				t.Errorf("FAIL. Expected message receipt '%+v', but got '%+v'", input.expectedMsgReceipt, msgReceipt)
			}
		})
	}
}
