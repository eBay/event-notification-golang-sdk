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
This package include message processor to process message based on their topics
*/
package processor

import (
	"fmt"

	pojo "github.com/ebay/event-notification-golang-sdk.git/lib/pojo"
)

//AccountDeletionMessageProcessor is to process account deletion
type AccountDeletionMessageProcessor struct {
}

//Implemenation for processing account deletion messages
//Input
//	message to be processed
func (a AccountDeletionMessageProcessor) Process(message *pojo.Message) {
	fmt.Println("Accoutn deletion processing")
	data := message.Notification.PayloadData
	fmt.Println(fmt.Sprintf(`\n==========================\nUser ID: %s`, data.UserID))
	fmt.Println(fmt.Sprintf("Username: %s\n==========================\n", data.Username))
}
