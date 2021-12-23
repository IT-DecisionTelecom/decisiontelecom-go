package main

import (
	"fmt"

	"github.com/IT-DecisionTelecom/decisiontelecom-go/sms"
	"github.com/IT-DecisionTelecom/decisiontelecom-go/viber"
	"github.com/IT-DecisionTelecom/decisiontelecom-go/viberplussms"
)

func main() {
	smsClientSendMessage()
	smsClientGetMessageStatus()
	smsClientGetBalance()

	viberClientSendMessage()
	viberClientGetMessageStatus()

	viberPlusSmsClientSendMessage()
	viberPlusSmsClientGetMessageStatus()
}

func smsClientSendMessage() {
	// Create new instance of the sms client.
	smsClient := sms.NewClient("<YOUR_LOGIN>", "<YOUR_PASSWORD>")

	// Create SMS message object
	message := sms.Message{
		ReceiverPhone: "380504444444",
		Sender:        "380505555555",
		Text:          "Test sms",
		Delivery:      true,
	}

	// Call client SendMessage method to send SMS message.
	msgId, err := smsClient.SendMessage(message)

	// Handle error if it has occurred while sending SMS message.
	if err != nil {
		smsError, ok := err.(sms.Error)
		if !ok {
			// A non-DecisionTelecom error occurred (like connection error).
			fmt.Printf("error while sending SMS message: %+v\n", err)
		} else {
			// DecisionTelecom error occurred.
			fmt.Printf("error while sending SMS message, error code: %d, error name: %s\n", smsError.Code, smsError.Code.String())
		}
	} else {
		// If no errors occurred, SendMessage method should return Id of the sent SMS message.
		fmt.Printf("message Id: %d\n", msgId)
	}
}

func smsClientGetMessageStatus() {
	// Create new instance of the sms client.
	smsClient := sms.NewClient("<YOUR_LOGIN>", "<YOUR_PASSWORD>")

	// Call client GetMessageStatus method to get SMS message status.
	status, err := smsClient.GetMessageStatus(5024173481)

	// Handle error if it has occurred while getting SMS message status.
	if err != nil {
		smsError, ok := err.(sms.Error)
		if !ok {
			// A non-DecisionTelecom error occurred (like connection error).
			fmt.Printf("error while getting SMS message status: %+v\n", err)
		} else {
			// DecisionTelecom error occurred.
			fmt.Printf("error while getting SMS message status, error code: %d, error name: %s\n", smsError.Code, smsError.Code.String())
		}
	} else {
		// If no errors occurred, GetMessageStatus method should return status of the sent SMS message.
		fmt.Printf("message status code: %d, message status name: %s\n", status, status.String())
	}
}

func smsClientGetBalance() {
	// Create new instance of the sms client.
	smsClient := sms.NewClient("<YOUR_LOGIN>", "<YOUR_PASSWORD>")

	// Call client GetBalance method to get SMS balance information.
	balance, err := smsClient.GetBalance()

	// Handle error if it has occurred while getting SMS balance.
	if err != nil {
		smsError, ok := err.(sms.Error)
		if !ok {
			// A non-DecisionTelecom error occurred (like connection error).
			fmt.Printf("error while getting SMS balance: %+v\n", err)
		} else {
			// DecisionTelecom error occurred.
			fmt.Printf("error while getting SMS balance, error code: %d, error name: %s\n", smsError.Code, smsError.Code.String())
		}
	} else {
		// If no errors occurred, GetBalance method should return status of the sent SMS message.
		fmt.Printf("balance information: Balance: %f, Credit: %f, Currency: %s\n",
			balance.BalanceAmount, balance.CreditAmount, balance.Currency)
	}
}

func viberClientSendMessage() {
	// Create new instance of the viber client.
	viberClient := viber.NewClient("<YOUR_ACCESS_KEY>")

	// Create viber message object
	message := viber.Message{
		Sender:         "Custom company",
		Receiver:       "380504444444",
		MessageType:    viber.TextImageButton,
		Text:           "Message content",
		ImageUrl:       "https://yourdomain.com/images/image.jpg",
		ButtonCaption:  "Join Us",
		ButtonAction:   "https://yourdomain.com/join-us",
		SourceType:     viber.Promotional,
		CallbackUrl:    "https://yourdomain.com/viber-callback",
		ValidityPeriod: 3600,
	}

	// Call client SendMessage method to send viber message.
	msgId, err := viberClient.SendMessage(message)

	// Handle error if it has occurred while sending viber message.
	if err != nil {
		viberError, ok := err.(viber.Error)
		if !ok {
			// A non-DecisionTelecom error occurred (like connection error).
			fmt.Printf("error while sending Viber message: %+v\n", err)
		} else {
			// DecisionTelecom error occurred.
			fmt.Printf("error while sending Viber message.\nerror name: %s\nerror message: %s\nerror code: %d\nerror status: %d\n",
				viberError.Name, viberError.Message, viberError.Code, viberError.Status)
		}
	} else {
		// If no errors occurred, SendMessage method should return Id of the sent Viber message.
		fmt.Printf("message Id: %d\n", msgId)
	}
}

func viberClientGetMessageStatus() {
	// Create new instance of the viber client.
	viberClient := viber.NewClient("<YOUR_ACCESS_KEY>")

	// Call client GetMessageStatus method to get viber message status.
	receipt, err := viberClient.GetMessageStatus(429)

	// Handle error if it has occurred while getting viber message status.
	if err != nil {
		viberError, ok := err.(viber.Error)
		if !ok {
			// A non-DecisionTelecom error occurred (like connection error).
			fmt.Printf("error while getting Viber message status: %+v\n", err)
		} else {
			// DecisionTelecom error occurred.
			fmt.Printf("error while getting Viber message status.\nerror name: %s\nerror message: %s\nerror code: %d\nerror status: %d\n",
				viberError.Name, viberError.Message, viberError.Code, viberError.Status)
		}
	} else {
		// If no errors occurred, GetMessageStatus method should return status of the sent Viber message.
		fmt.Printf("viber message status code: %d (%s)\n", receipt.Status, receipt.Status.String())
	}
}

func viberPlusSmsClientSendMessage() {
	// Create new instance of the viber plus SMS client.
	viberSmsClient := viberplussms.NewClient("<YOUR_ACCESS_KEY>")

	// Create viber plus SMS message object
	message := viberplussms.Message{
		Message: viber.Message{
			Sender:         "380504444444",
			Receiver:       "380504444444",
			MessageType:    viber.TextImageButton,
			Text:           "Viber Message",
			ImageUrl:       "https://yourdomain.com/images/image.jpg",
			ButtonCaption:  "Join Us",
			ButtonAction:   "https://yourdomain.com/join-us",
			SourceType:     viber.Promotional,
			CallbackUrl:    "https://yourdomain.com/viber-callback",
			ValidityPeriod: 50,
		},
		SmsText: "Text SMS Message",
	}

	// Call client SendMessage method to send viber plus SMS message.
	msgId, err := viberSmsClient.SendMessage(message)

	// Handle error if it has occurred while sending viber plus SMS message.
	if err != nil {
		viberError, ok := err.(viber.Error)
		if !ok {
			// A non-DecisionTelecom error occurred (like connection error).
			fmt.Printf("error while sending Viber plus SMS message: %+v\n", err)
		} else {
			// DecisionTelecom error occurred.
			fmt.Printf("error while sending Viber plus SMS message.\nerror name: %s\nerror message: %s\nerror code: %d\nerror status: %d\n",
				viberError.Name, viberError.Message, viberError.Code, viberError.Status)
		}
	} else {
		// If no errors occurred, SendMessage method should return Id of the sent Viber plus SMS message.
		fmt.Printf("message Id: %d\n", msgId)
	}
}

func viberPlusSmsClientGetMessageStatus() {
	// Create new instance of the viber plus SMS client.
	viberSmsClient := viberplussms.NewClient("<YOUR_ACCESS_KEY>")

	// Call client GetMessageStatus method to get viber plus SMS message status.
	receipt, err := viberSmsClient.GetMessageStatus(429)

	// Handle error if it has occurred while getting viber plus SMS message status.
	if err != nil {
		viberError, ok := err.(viber.Error)
		if !ok {
			// A non-DecisionTelecom error occurred (like connection error).
			fmt.Printf("error while getting Viber plus SMS message status: %+v\n", err)
		} else {
			// DecisionTelecom error occurred.
			fmt.Printf("error while getting Viber plus SMS message status.\nerror name: %s\nerror message: %s\nerror code: %d\nerror status: %d\n",
				viberError.Name, viberError.Message, viberError.Code, viberError.Status)
		}
	} else {
		// If no errors occurred, GetMessageStatus method should return status of the sent Viber plus SMS message.
		fmt.Printf(
			"viber message Id: %d\nviber message status: %d (%s)\nSMS message Id: %d\nSMS message status: %d (%s)\n",
			receipt.MessageId, receipt.Status, receipt.Status.String(), receipt.SmsMessageId, receipt.SmsMessageStatus, receipt.SmsMessageStatus.String())
	}
}
