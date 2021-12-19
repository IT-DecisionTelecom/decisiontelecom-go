package viber

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// MakeHttpRequest performs HTTP request to the Viber endpoints and returns response body.
func MakeHttpRequest(apiKey string, url string, requestContent interface{}) (string, error) {
	jsonRequest, _ := json.Marshal(requestContent)
	accessKeyBase64 := base64.StdEncoding.EncodeToString([]byte(apiKey))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Basic "+accessKeyBase64)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	httpClient := &http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	bodyStr := string(bodyBytes)

	// Process unsuccessful status codes
	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		return "", Error{Status: response.StatusCode, Name: http.StatusText(response.StatusCode)}
	}

	// If response contains "name", "message", "code" and "status" words, treat it as a ViberError
	if strings.Contains(bodyStr, "name") && strings.Contains(bodyStr, "message") &&
		strings.Contains(bodyStr, "code") && strings.Contains(bodyStr, "status") {
		var viberError Error
		if err := json.Unmarshal(bodyBytes, &viberError); err != nil {
			return "", err
		}

		return "", viberError
	}

	return bodyStr, nil
}
