/*
 * Copyright (c) 2022 eBay Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */
package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	sdk "github.com/ebay/event-notification-golang-sdk.git/lib/notification"
	pojo "github.com/ebay/event-notification-golang-sdk.git/lib/pojo"
)

var config = new(pojo.Config)
var message = new(pojo.Message)
var metadata = new(pojo.Metadata)
var response = new(pojo.Response)
var payloadData = new(pojo.PayloadData)
var notification = new(pojo.Notification)
var payload = new(pojo.Payload)
var signature string
var Config *pojo.Config

func TestMain(m *testing.M) {
	fmt.Println("*****Test Suite Started*****")
	// exec setUp function
	setUp()
	// exec test and this returns an exit code to pass to os
	retCode := m.Run()
	// exec tearDown function
	tearDown()
	// If exit code is distinct of zero,
	// the test will be failed (red)
	os.Exit(retCode)
}

// setUp function to add initial setup
func setUp() {
	fmt.Println("Loading... Mock Data")

	config.VerificationToken = "71745723-d031-455c-bfa5-f90d11b4f20a"
	config.Endpoint = "http://www.testendpoint.com/webhook"

	metadata = &pojo.Metadata{Topic: "MARKETPLACE_ACCOUNT_DELETION",
		SchemaVersion: "1.0",
		Deprecated:    false}

	response = &pojo.Response{Key: "-----BEGIN PUBLIC KEY-----MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEZhhxXKtR+TOvtDbgTPCkSof02qgBB7IsYOyf76ilExJ/upAa/vKIKheOoCyOpcLmi4t0b4uepb7LLjmMr90FUg==-----END PUBLIC KEY-----",
		Algorithm: "ECDSA",
		Digest:    "SHA1"}

	payloadData = &pojo.PayloadData{Username: "test_user",
		UserID:    "ma8vp1jySJC",
		EiasToken: "nY+sHZ2PrBmdj6wVnY+sEZ2PrA2dj6wJnY+gAZGEpwmdj6x9nY+seQ=="}

	notification = &pojo.Notification{NotificationID: "49feeaeb-4982-42d9-a377-9645b8479411_33f7e043-fed8-442b-9d44-791923bd9a6d",
		EventDate:           "2021-03-19T20:43:59.462Z",
		PublishDate:         "2021-03-19T20:43:59.679Z",
		PublishAttemptCount: 1,
		PayloadData:         *payloadData}

	message = &pojo.Message{Metadata: *metadata, Notification: *notification}

	payload = &pojo.Payload{Signature: "eyJhbGciOiJlY2RzYSIsImtpZCI6Ijk5MzYyNjFhLTdkN2ItNDYyMS1hMGYxLTk2Y2NiNDI4YWY0OSIsInNpZ25hdHVyZSI6Ik1FWUNJUUNmeGZJV3V4bVdjSUJRSjljNS9YN2lHREpxczJSQ0dzQkVhQWppbnlycmZBSWhBSVY2d0djVGlCdVY1S0pVaWYyaG9reXJMK1E5c3NIa2FkK214Mm5FRTI1dyIsImRpZ2VzdCI6IlNIQTEifQ==",
		Message:   *message,
		Response:  *response,
		PublicKey: "9936261a-7d7b-4621-a0f1-96ccb428af49"}
}

// tearDown function to clean up
func tearDown() {
	fmt.Println("*****Test Suite Completed*****")
}

func loadTestData(key string) {
	var testData map[string]interface{}
	// read file
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println("Failed to read test file:", err)
	}

	// unmarshall it
	err = json.Unmarshal(data, &testData)
	if err != nil {
		fmt.Println("Failed to unmarshall test file:", err)
	}

	currentData := testData[key].(map[string]interface{})

	ds, _ := json.Marshal(currentData["message"])
	json.Unmarshal(ds, &message)

	ds, _ = json.Marshal(currentData["signature"])
	json.Unmarshal(ds, &signature)
}

func loadConfigData(*pojo.Config) {
	// read file
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	data, err := ioutil.ReadFile(filepath.Dir(pwd) + "/examples/config.json")
	if err != nil {
		fmt.Print("Failed to read config file:", err)
	}

	// unmarshall it
	err = json.Unmarshal(data, &Config)
	if err != nil {
		fmt.Println("Failed to unmarshall config file:", err)
	}
}

func TestValidateAndProcess(t *testing.T) {
	methods := sdk.ValidateAndProcess
	if !IsFunc(methods) {
		t.Errorf("ValidateAndProcess is not an function")
	}
}

func TestValidateEndpoint(t *testing.T) {
	methods := sdk.ValidateEndpoint
	if !IsFunc(methods) {
		t.Errorf("ValidateEndpoint is not an function")
	}
}

func TestValidateEndpointChallengeCodeEmpty(t *testing.T) {
	err, _ := sdk.ValidateEndpoint("", &pojo.Config{})
	if !strings.EqualFold(err, `The "challengeCode" is required.`) {
		t.Errorf("challengeCode can't be empty")
	}
}

func TestValidateEndpointConfigEmpty(t *testing.T) {
	err, _ := sdk.ValidateEndpoint("Something", nil)
	if !strings.EqualFold(err, `Please provide the config.`) {
		t.Errorf("Config can't be empty")
	}
}

func TestValidateEndpointConfigEndpoint(t *testing.T) {
	config.Endpoint = ""
	err, _ := sdk.ValidateEndpoint("Something", config)
	if !strings.EqualFold(err, `The "endpoint" is required.`) {
		t.Errorf(`The "endpoint" is required.`)
	}
}

func TestValidateEndpointConfigVerificationToken(t *testing.T) {
	config.Endpoint = "123"
	config.VerificationToken = ""
	err, _ := sdk.ValidateEndpoint("Something", config)
	if !strings.EqualFold(err, `The "verificationToken" is required.`) {
		t.Errorf(`The "verificationToken" is required.`)
	}
}

func TestValidateAndProcessMessage(t *testing.T) {
	err, _ := sdk.ValidateAndProcess(nil, "signature", config, "QA")
	if !strings.EqualFold(err, `Please provide the message.`) {
		t.Errorf(`Please provide the message.`)
	}
}

func TestValidateAndProcessSignature(t *testing.T) {
	err, _ := sdk.ValidateAndProcess(message, "", config, "QA")
	if !strings.EqualFold(err, `Please provide the signature.`) {
		t.Errorf(`Please provide the signature.`)
	}
}

func TestValidateAndProcessConfig(t *testing.T) {
	err, _ := sdk.ValidateAndProcess(message, "signature", nil, "QA")
	if !strings.EqualFold(err, `Please provide the config.`) {
		t.Errorf(`Please provide the config.`)
	}
}

func TestValidateAndProcessConfigClientId(t *testing.T) {
	err, _ := sdk.ValidateAndProcess(message, "signature", config, "QA")
	if !strings.EqualFold(err, `Please provide the Client ID.`) {
		t.Errorf(`Please provide the Client ID.`)
	}
}

func TestValidateAndProcessConfigClientSecret(t *testing.T) {
	config.Production.ClientID = "clientId"
	config.Sandbox.ClientID = "clientId"
	err, _ := sdk.ValidateAndProcess(message, "signature", config, "QA")
	if !strings.EqualFold(err, `Please provide the Client secret.`) {
		t.Errorf(`Please provide the Client secret.`)
	}
}

func TestValidateAndProcessConfigClientEnv(t *testing.T) {
	config.Production.ClientSecret = "clientSecret"
	config.Sandbox.ClientSecret = "clientSecret"
	err, _ := sdk.ValidateAndProcess(message, "signature", config, "QA")
	if !strings.EqualFold(err, `Please provide the Environment.`) {
		t.Errorf(`Please provide the Environment.`)
	}
}

func TestValidateSignatureValidSuccess(t *testing.T) {
	loadTestData("VALID")
	loadConfigData(Config)
	err, _ := sdk.ValidateAndProcess(message, signature, Config, "PRODUCTION")
	if !strings.EqualFold(err, "") {
		t.Errorf(`Failed to process`)
	}
}

func TestValidateSignatureInvalidSuccess(t *testing.T) {
	loadTestData("INVALID")
	loadConfigData(Config)
	err, _ := sdk.ValidateAndProcess(message, signature, Config, "PRODUCTION")
	if !strings.EqualFold(err, "412") {
		t.Errorf(`Failed to process`)
	}
}

func TestValidateSignatureMismatchSuccess(t *testing.T) {
	loadTestData("SIGNATURE_MISMATCH")
	loadConfigData(Config)
	err, _ := sdk.ValidateAndProcess(message, signature, Config, "PRODUCTION")
	if !strings.EqualFold(err, "412") {
		t.Errorf(`Failed to process`)
	}
}

func IsFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}
