package viber

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// BaseClient is a base client for Viber and Viber plus SMS operations.
type BaseClient struct {
	ApiKey string
}

// SendMessage sends Viber message.
func (cl *BaseClient) SendMessage(message interface{}) (MessageId, error) {
	url := fmt.Sprintf("%s/send-viber", baseUrl)
	responseBody, err := cl.makeHttpRequest(url, message)
	if err != nil {
		return -1, err
	}

	var responseMap map[string]int64
	if err := json.Unmarshal(responseBody, &responseMap); err != nil {
		return -1, err
	}

	msgId, ok := responseMap[messageIdPropertyName]
	if !ok {
		return -1, fmt.Errorf("invalid response: property '%s' was not found", messageIdPropertyName)
	}

	return MessageId(msgId), nil
}

// GetMessageStatus
func (cl *BaseClient) GetMessageStatusResponse(messageId MessageId, result interface{}) error {
	url := fmt.Sprintf("%s/receive-viber", baseUrl)
	request := map[string]MessageId{messageIdPropertyName: messageId}

	responseBody, err := cl.makeHttpRequest(url, request)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return err
	}

	return nil
}

// MakeHttpRequest performs HTTP request to the Viber endpoints and returns response body.
func (cl *BaseClient) makeHttpRequest(url string, requestContent interface{}) ([]byte, error) {
	jsonRequest, _ := json.Marshal(requestContent)
	accessKeyBase64 := base64.StdEncoding.EncodeToString([]byte(cl.ApiKey))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Basic "+accessKeyBase64)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	httpClient := &http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	bodyStr := string(bodyBytes)

	// Process unsuccessful status codes
	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		return nil, Error{Status: response.StatusCode, Name: http.StatusText(response.StatusCode)}
	}

	// If response contains "name", "message", "code" and "status" words, treat it as a ViberError
	if strings.Contains(bodyStr, "name") && strings.Contains(bodyStr, "message") &&
		strings.Contains(bodyStr, "code") && strings.Contains(bodyStr, "status") {
		var viberError Error
		if err := json.Unmarshal(bodyBytes, &viberError); err != nil {
			return nil, err
		}

		return nil, viberError
	}

	return bodyBytes, nil
}
