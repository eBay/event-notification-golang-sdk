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
package constants

const (
	AccessToken                       = "access_token"
	APIScope                          = "https://api.ebay.com/oauth/api_scope"
	Authorization                     = "Authorization"
	Basic                             = "Basic "
	Bearer                            = "bearer "
	ClientCredentials                 = "client_credentials"
	ContentType                       = "Content-Type"
	ContentTypeApplication            = "application/x-www-form-urlencoded"
	Get                               = "GET"
	GrantType                         = "grant_type"
	IdentifyPath                      = "/identity/v1/oauth2/token"
	KeyEnd                            = "-----END PUBLIC KEY-----"
	KeyStart                          = "-----BEGIN PUBLIC KEY-----"
	NotificationAPIEndpointProduction = "https://api.ebay.com/commerce/notification/v1/public_key/"
	NotificationAPIEndpointSandbox    = "https://api.sandbox.ebay.com/commerce/notification/v1/public_key/"
	Post                              = "POST"
	Scope                             = "scope"
	Success                           = "Success"
	TopicsMarketplaceAccountDeletion  = "MARKETPLACE_ACCOUNT_DELETION"
	XEbaySignature                    = "X-Ebay-Signature"
	EnvironmentSandbox                = "SANDBOX"
	EnvironmentProduction             = "PRODUCTION"
	Error                             = "Error"
	HTTPStatusCodeNoContent           = "204"
	HTTPStatusCodeOk                  = "200"
	HTTPStatusCodePreconditionFailed  = "412"
	HTTPStatusCodeInternalServerError = "500"
)
