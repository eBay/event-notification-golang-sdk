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
This package contains required helper functions
*/
package helper

import (
	"crypto/ecdsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"

	constants "github.com/ebay/event-notification-golang-sdk.git/lib/constants"
	pojo "github.com/ebay/event-notification-golang-sdk.git/lib/pojo"
	service "github.com/ebay/event-notification-golang-sdk.git/lib/service"
)

//Get XeBay Signature header takes in signature decode using base64 and returns it
//Returns base64 decoded signature
//Input
//	signatureHeader - base64 encoded signature
//Returns
//	base64 decoded signature
func getXeBaySignatureHeader(signatureHeader string) *pojo.XeBaySignatureHeader {
	rawDecodedText, err := base64.StdEncoding.DecodeString(signatureHeader)
	if err != nil {
		panic(err)
	}
	var signature pojo.XeBaySignatureHeader
	json.Unmarshal([]byte(rawDecodedText), &signature)
	return &signature
}

//The format key function convert key by adding newline before/after comments
//Input
//	key - unformatted key
//Returns
//	key - formatted key
func formatKey(key string) string {

	key = strings.Replace(key, constants.KeyStart, fmt.Sprintf("%s\n", constants.KeyStart), 1)
	key = strings.Replace(key, constants.KeyEnd, fmt.Sprintf("\n%s", constants.KeyEnd), 1)

	return key
}

//ValidateSignature is to validate signature used in request
//Returns string Success/Error
//Input
//	message - message details 
//	signatureHeader - base64 encoded signature
//	config - specific custom environment
//Returns
//	string Success/Error
func ValidateSignature(message *pojo.Message, signatureHeader string, config *pojo.CustomEnvironment) string {

	// Base64 decode the signatureHeader and convert to JSON
	xeBaySignature := getXeBaySignatureHeader(signatureHeader)

	// // Get the public key
	publicKey := service.GetPublicKey(xeBaySignature.Kid, config)

	var pubPEMData = []byte(formatKey(publicKey.Key))
	block, _ := pem.Decode(pubPEMData)
	if block == nil {
		fmt.Println("Invalid PEM Block")
		return constants.Error
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return constants.Error
	}

	pubKey := key.(*ecdsa.PublicKey)
	signature, err := base64.StdEncoding.DecodeString(xeBaySignature.Signature)
	if err != nil {
		fmt.Println(err)
		return constants.Error
	}

	byteArr, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return constants.Error
	}

	hash := sha1.Sum(byteArr)
	success := ecdsa.VerifyASN1(pubKey, hash[:], signature)

	if success != true {
		fmt.Println(err)
		return constants.Error
	}

	return constants.Success
}

//GenerateChallengeResponse is used to generate challenge response for given challenge code
//Input
//	challengeCode - challengeCode to be processed
//	config - config details for processing
//Returns
//	challenge response
func GenerateChallengeResponse(challengeCode string, config *pojo.Config) string {
	hasher := sha256.New()
	hasher.Write([]byte(challengeCode))
	hasher.Write([]byte(config.VerificationToken))
	hasher.Write([]byte(config.Endpoint))

	digest := hasher.Sum(nil)
	digestStr := hex.EncodeToString(digest)
	return digestStr
}
