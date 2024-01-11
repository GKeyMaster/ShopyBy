/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// SupersimV1BillingPeriod struct for SupersimV1BillingPeriod
type SupersimV1BillingPeriod struct {
	// The SID of the Billing Period.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) the Super SIM belongs to.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Super SIM the Billing Period belongs to.
	SimSid *string `json:"sim_sid,omitempty"`
	// The start time of the Billing Period specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	StartTime *time.Time `json:"start_time,omitempty"`
	// The end time of the Billing Period specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EndTime    *time.Time `json:"end_time,omitempty"`
	PeriodType *string    `json:"period_type,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}