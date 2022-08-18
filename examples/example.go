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
 */
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	constants "github.com/ebay/event-notification-golang-sdk.git/lib/constants"
	sdk "github.com/ebay/event-notification-golang-sdk.git/lib/notification"
	"github.com/ebay/event-notification-golang-sdk.git/lib/pojo"
	"io/ioutil"
	"net/http"
	"strings"
)

//Process Notification request
//This function handles processing message notification resquest and respond accordingly
//Input
//	gin.Context - Request/Response context
func processNotification(c *gin.Context) {
	var body pojo.Message
	env := "PRODUCTION"
	json.NewDecoder(c.Request.Body).Decode(&body)
	signature := c.Request.Header[constants.XEbaySignature][0]
	err, responseCode := sdk.ValidateAndProcess(&body, signature, Config, env)
	if strings.EqualFold(err, constants.HTTPStatusCodePreconditionFailed) {
		fmt.Println(`Signature validation failed`)
		c.JSON(http.StatusInternalServerError, "Signature validation processing failure")
	} else if !strings.EqualFold(err, "") {
		fmt.Println(`Something went wrong`)
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	} else {
		fmt.Println(responseCode)
		c.JSON(http.StatusOK, "Successful")
	}
}

//Get challenge code
//This function handles in get challenge code request and respond accordingly
//Input
//	gin.Context - Request/Response context
func getChallengeCode(c *gin.Context) {
	challengeCode := c.Query("challenge_code")
	if !strings.EqualFold(challengeCode, "") {
		// challengeResponse := challengeCode
		err, challengeResponse := sdk.ValidateEndpoint(challengeCode, Config)
		if !strings.EqualFold(err, "") {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, gin.H{
			"challengeResponse": challengeResponse,
		})

	} else {
		c.JSON(http.StatusInternalServerError, "Missing 'challenge_code' request param")
	}

}

//To load config data
//This functions takes in config object and load data from config.json file.
//Input
//	Config - reference to Config object
//Returns
//	error response on reading/unmarshalling failure
func loadConfig(*pojo.Config) string {

	// read file
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Print("Failed to read config file:", err)
		return "Failed to read config file"
	}

	// unmarshall it
	err = json.Unmarshal(data, &Config)
	if err != nil {
		fmt.Println("Failed to unmarshall config file:", err)
		return "Failed to unmarshall config file"
	}
	return ""
}

func main() {
	err := loadConfig(Config)
	if !strings.EqualFold(err, "") {
		panic("Failed to load config file")
	}

	router := gin.Default()
	router.POST("/webhook", processNotification)
	router.GET("/webhook", getChallengeCode)

	router.Run("localhost:8080")
}

var Config *pojo.Config
