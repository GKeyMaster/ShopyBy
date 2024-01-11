/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// FlexV1InteractionChannel struct for FlexV1InteractionChannel
type FlexV1InteractionChannel struct {
	// The unique string created by Twilio to identify an Interaction Channel resource, prefixed with UO.
	Sid *string `json:"sid,omitempty"`
	// The unique string created by Twilio to identify an Interaction resource, prefixed with KD.
	InteractionSid *string `json:"interaction_sid,omitempty"`
	Type           *string `json:"type,omitempty"`
	Status         *string `json:"status,omitempty"`
	// The Twilio error code for a failed channel.
	ErrorCode *int `json:"error_code,omitempty"`
	// The error message for a failed channel.
	ErrorMessage *string                 `json:"error_message,omitempty"`
	Url          *string                 `json:"url,omitempty"`
	Links        *map[string]interface{} `json:"links,omitempty"`
}
