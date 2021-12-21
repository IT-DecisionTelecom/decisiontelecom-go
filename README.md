IT-Decision Telecom Go SDK
===============================

Convenient Go client for IT-Decision Telecom messaging API.

[![test](https://github.com/IT-DecisionTelecom/decisiontelecom-go/actions/workflows/test.yml/badge.svg)](https://github.com/IT-DecisionTelecom/decisiontelecom-go/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Requirements
-----

- [Sign up](https://web.it-decision.com/site/signup) for a free IT-Decision Telecom account
- Request login and password to send SMS messages and access key to send Viber messages
- You should have an application written in Go to make use of this SDK

Installation
-----

The easiest way to use the IT-Decision Telecom SDK in your Go project is to install it using _go get_:

```
go get github.com/IT-DecisionTelecom/decisiontelecom-go
```

Usage
-----

We have put some self-explanatory usage examples in the *examples* folder,
but here is a quick reference on how IT-Decision Telecom clients work.
First, you need to import DecisionTelecom package which corresponds to your needs: 

```go
import "github.com/IT-DecisionTelecom/decisiontelecom-go/sms"
import "github.com/IT-DecisionTelecom/decisiontelecom-go/viber"
import "github.com/IT-DecisionTelecom/decisiontelecom-go/viberplussms"
```

Then, create an instance of a required client. Be sure to use real login, password and access key.

```go
smsClient := sms.NewClient("<YOUR_LOGIN>", "<YOUR_PASSWORD>")
viberClient := viber.NewClient("<YOUR_ACCESS_KEY>")
viberSmsClient := viberplussms.NewClient("<YOUR_ACCESS_KEY>")
```

Now you can use created client to perform some operations. For example, this is how you can get your SMS balance:

```go
// Request SMS balance information
balance, err := smsClient.GetBalance()
if err != nil {
    // Handle error.
    return
}

// Process results.
fmt.Printf("balance information: Balance: %f, Credit: %f, Currency: %s\n",
			balance.BalanceAmount, balance.CreditAmount, balance.Currency)
```

Please see other examples in the _examples_ folder for a complete overview of all available SDK calls.

### Error handling
All client methods return an error along with the desired result. Returned error might be a specific DecisionTelecom error.
SMS client methods might return error code, Viber and Viber plus SMS client methods might return `Error` object.

An error handling example is shown below. Let's look how we can gain more in-depth insight in what exactly went wrong:

```go
_, err := smsClient.GetBalance()
if err != nil {
    smsError, ok := err.(sms.Error)
    if !ok {
        // A non-DecisionTelecom error occurred (like connection error).
        fmt.Printf("error while getting SMS balance: %+v\n", err)
    } else {
        // DecisionTelecom error occurred.
        fmt.Printf("error while getting SMS balance, error code: %d (%s)\n", 
            smsError.Code, smsError.Code.String())
    }
}
```

#### SMS errors
SMS client methods return errors in form of the error code. Here are all possible error codes:

- 40 - Invalid number
- 41 - Incorrect sender
- 42 - Invalid message ID
- 43 - Incorrect JSON
- 44 - Invalid login or password
- 45 - User locked
- 46 - Empty text
- 47 - Empty login
- 48 - Empty password
- 49 - Not enough money to send a message
- 50 - Authentication error

#### Viber errors
Viber and Viber plus SMS client methods return errors in form of a struct with the `Name`, `Message`, `Code` and `Status` properties.

If underlying API request returns unsuccessful status code (like 401 Unauthorized),
then client methods will return error with only `Name` and `Status` properties set:

```json
{
  "name": "Unauthorized",
  "status": 401
}
```

Known Viber errors are:

```json
{
  "name": "Too Many Requests",
  "message": "Rate limit exceeded",
  "code": 0,
  "status": 429
}
```

```json
{
  "name": "Invalid Parameter: [param_name]",
  "message": "Empty parameter or parameter validation error",
  "code": 1,
  "status": 400
}
```

```json
{
  "name": "Internal server error",
  "message": "The server encountered an unexpected condition which prevented it from fulfilling the request",
  "code": 2,
  "status": 500
}
```

```json
{
  "name": "Topup balance is required",
  "message": "User balance is empty",
  "code": 3,
  "status": 402
}
```