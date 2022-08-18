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
 This package include service calls
 */
 package service

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	lru "github.com/hashicorp/golang-lru"

	constants "github.com/ebay/event-notification-golang-sdk.git/lib/constants"
	pojo "github.com/ebay/event-notification-golang-sdk.git/lib/pojo"
)

var m = make(map[string]pojo.Environment)
var cache, error = lru.New(100)

//Get App Token
//Input
//	request config
//Returns
//	app token string
func getAppToken(req *(pojo.CustomEnvironment)) string {

	var encodedStr string
	encodedStr = constants.Basic + b64.URLEncoding.EncodeToString([]byte(req.ClientID+":"+req.ClientSecret))

	u, _ := url.ParseRequestURI("https://" + req.BaseURL)
	u.Path = constants.IdentifyPath
	urlStr := u.String()

	data := url.Values{}
	data.Set(constants.GrantType, constants.ClientCredentials)
	data.Set(constants.Scope, constants.APIScope)

	client := &http.Client{}
	r, _ := http.NewRequest(constants.Post, urlStr, strings.NewReader(data.Encode()))
	r.Header.Add(constants.Authorization, encodedStr)
	r.Header.Add(constants.ContentType, constants.ContentTypeApplication)

	resp, _ := client.Do(r)

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return fmt.Sprint(res[constants.AccessToken])
}

//GetPublicKey is used to get pblic key for provided config
//Input
//	keyId
//	config details
//Returns
//	public key
func GetPublicKey(keyID string, config *pojo.CustomEnvironment) *pojo.Response {

	publicKeyVal, isPresent := cache.Get(keyID)
	if isPresent {
		publicKey := publicKeyVal.(pojo.Response)
		return &publicKey
	}

	var notifyEndpoint string
	if config.Environment == constants.EnvironmentSandbox {
		notifyEndpoint = constants.NotificationAPIEndpointSandbox
	} else {
		notifyEndpoint = constants.NotificationAPIEndpointProduction
	}

	token := getAppToken(config)

	client := &http.Client{}
	r, _ := http.NewRequest(constants.Get, notifyEndpoint+keyID, nil)
	r.Header.Add(constants.Authorization, constants.Bearer+token)
	r.Header.Add(constants.ContentType, constants.ContentTypeApplication)

	resp, _ := client.Do(r)

	var res pojo.Response
	json.NewDecoder(resp.Body).Decode(&res)

	cache.Add(keyID, res)

	return &res
}
