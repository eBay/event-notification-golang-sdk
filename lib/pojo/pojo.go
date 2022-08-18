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
 *
This package contains all required pojo
*/
package pojo

//Config is configuration file object
type Config struct {
	Sandbox           Environment `json:"SANDBOX"`
	Production        Environment `json:"PRODUCTION"`
	Endpoint          string      `json:"endpoint"`
	VerificationToken string      `json:"verificationToken"`
}

//Environment is configuration environment specific file
type Environment struct {
	BaseURL      string `json:"baseUrl"`
	RedirectURI  string `json:"redirectUri"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	DevID        string `json:"devid"`
}

//CustomEnvironment is configuration environment specific file
type CustomEnvironment struct {
	BaseURL      string `json:"baseUrl"`
	RedirectURI  string `json:"redirectUri"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	DevID        string `json:"devid"`
	Environment  string
}

//Payload is payload object
type Payload struct {
	Signature string   `json:"signature"`
	Message   Message  `json:"message"`
	Response  Response `json:"response"`
	PublicKey string   `json:"public_key"`
}

//Message is norification message
type Message struct {
	Metadata     Metadata     `json:"metadata"`
	Notification Notification `json:"notification"`
}

//Metadata is notification metadata
type Metadata struct {
	Topic         string `json:"topic"`
	SchemaVersion string `json:"schemaVersion"`
	Deprecated    bool   `json:"deprecated"`
}

//Notification is notification object
type Notification struct {
	NotificationID      string      `json:"notificationId"`
	EventDate           string      `json:"eventDate"`
	PublishDate         string      `json:"publishDate"`
	PublishAttemptCount int         `json:"publishAttemptCount"`
	PayloadData         PayloadData `json:"data"`
}

//PayloadData is user payload
type PayloadData struct {
	Username  string `json:"username"`
	UserID    string `json:"userId"`
	EiasToken string `json:"eiasToken"`
}

//Response is response object
type Response struct {
	Key       string `json:"key"`
	Algorithm string `json:"algorithm"`
	Digest    string `json:"digest"`
}

//XeBaySignatureHeader is pojo for signature creation
type XeBaySignatureHeader struct {
	Alg       string `json:"alg"`
	Digest    string `json:"digest"`
	Kid       string `json:"kid"`
	Signature string `json:"signature"`
}

const (
	//SANDBOX Environment string
	SANDBOX string = "SANDBOX"

	//PRODUCTION Environment string
	PRODUCTION = "PRODUCTION"
)
