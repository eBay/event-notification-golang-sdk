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
	constants "github.com/ebay/event-notification-golang-sdk.git/lib/constants"
	pojo "github.com/ebay/event-notification-golang-sdk.git/lib/pojo"
)

//Processor is generice processor for message processing by topics
type Processor interface {
	Process(*pojo.Message)
}

var obj Processor

//GetProcessor is used to get processor for specified topic
//Input
//	topic to be processed
//Returns
//	processor for the topic
func GetProcessor(topic string) Processor {
	switch topic {
	case constants.TopicsMarketplaceAccountDeletion:
		obj := AccountDeletionMessageProcessor{}
		return obj
	default:
		panic("Message processor not registered for " + topic)
	}
}
