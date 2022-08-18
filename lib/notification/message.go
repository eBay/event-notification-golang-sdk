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
This package implement two methods required for Event Notification Processing
 ValidateAndProcess - To validate signature and perform necessary action for received notification
 ValidateEndpoint - To Validate url endpoint readiness based on challenge code and response
*/
package notification

import (
	constants "github.com/ebay/event-notification-golang-sdk.git/lib/constants"
	helper "github.com/ebay/event-notification-golang-sdk.git/lib/helper"
	pojo "github.com/ebay/event-notification-golang-sdk.git/lib/pojo"
	processor "github.com/ebay/event-notification-golang-sdk.git/lib/processor"
	"strings"
)

//Returns CustomEnv object
//Input
//	env - environment details
//	environment - environment name
//Returns
//	customEnvironment - details of specified env
func getCustomEnv(env *pojo.Environment, environment string) *pojo.CustomEnvironment {
	return &pojo.CustomEnvironment{env.BaseURL, env.RedirectURI, env.ClientID, env.ClientSecret, env.DevID, environment}
}

//ValidateAndProcess is to validate request and process the message
//Input
//	message - message to be processed
//	signature - signature of sender
//	config - config details for processing
//	environment - environment name
//Returns
//	error
//	response body
func ValidateAndProcess(message *pojo.Message, signature string, config *pojo.Config, environment string) (string, string) {
	var err string
	if message == nil || &message.Metadata == nil || &message.Notification == nil {
		err = `Please provide the message.`
	} else if signature == "" {
		err = `Please provide the signature.`
	} else if config == nil {
		err = `Please provide the config.`
	} else if config.Production.ClientID == "" || config.Sandbox.ClientID == "" {
		err = `Please provide the Client ID.`
	} else if config.Production.ClientSecret == "" || config.Sandbox.ClientSecret == "" {
		err = `Please provide the Client secret.`
	} else if len(environment) == 0 || (environment != constants.EnvironmentProduction && environment != constants.EnvironmentSandbox) {
		err = `Please provide the Environment.`
	}

	if !strings.EqualFold(err, "") {
		return err, ""
	}
	var customEnv *pojo.CustomEnvironment
	if environment == constants.EnvironmentSandbox {
		customEnv = getCustomEnv(&config.Sandbox, environment)
	} else {
		customEnv = getCustomEnv(&config.Production, environment)
	}

	response := helper.ValidateSignature(message, signature, customEnv)
	if strings.EqualFold(response, constants.Success) {
		processor.GetProcessor(message.Metadata.Topic).Process(message);
		return "", constants.HTTPStatusCodeNoContent
	} else if strings.EqualFold(response, constants.Error) {
		return constants.HTTPStatusCodePreconditionFailed, ""
	}
	return constants.HTTPStatusCodeInternalServerError, ""
}

//ValidateEndpoint is to validate endpoint using challengeCode
//Input
//	challengeCode - challengeCode to be processed
//	config - config details for processing
//Returns
//	error
//	challenge response
func ValidateEndpoint(challengeCode string, config *pojo.Config) (string, string) {
	var err string
	if strings.EqualFold(challengeCode, "") {
		err = `The "challengeCode" is required.`
	} else if config == nil {
		err = `Please provide the config.`
	} else if strings.EqualFold(config.Endpoint, "") {
		err = `The "endpoint" is required.`
	} else if strings.EqualFold(config.VerificationToken, "") {
		err = `The "verificationToken" is required.`
	}

	if !strings.EqualFold(err, "") {
		return err, ""
	}

	return err, helper.GenerateChallengeResponse(challengeCode, config)
}
