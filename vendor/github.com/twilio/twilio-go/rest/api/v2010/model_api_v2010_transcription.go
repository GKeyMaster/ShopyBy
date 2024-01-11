/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// ApiV2010Transcription struct for ApiV2010Transcription
type ApiV2010Transcription struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the transcription.
	ApiVersion *string `json:"api_version,omitempty"`
	// The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *string `json:"date_created,omitempty"`
	// The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// The duration of the transcribed audio in seconds.
	Duration *string `json:"duration,omitempty"`
	// The charge for the transcript in the currency associated with the account. This value is populated after the transcript is complete so it may not be available immediately.
	Price *float32 `json:"price,omitempty"`
	// The currency in which `price` is measured, in [ISO 4127](https://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`).
	PriceUnit *string `json:"price_unit,omitempty"`
	// The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) from which the transcription was created.
	RecordingSid *string `json:"recording_sid,omitempty"`
	// The unique string that that we created to identify the Transcription resource.
	Sid    *string `json:"sid,omitempty"`
	Status *string `json:"status,omitempty"`
	// The text content of the transcription.
	TranscriptionText *string `json:"transcription_text,omitempty"`
	// The transcription type. Can only be: `fast`.
	Type *string `json:"type,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
}

func (response *ApiV2010Transcription) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid        *string      `json:"account_sid"`
		ApiVersion        *string      `json:"api_version"`
		DateCreated       *string      `json:"date_created"`
		DateUpdated       *string      `json:"date_updated"`
		Duration          *string      `json:"duration"`
		Price             *interface{} `json:"price"`
		PriceUnit         *string      `json:"price_unit"`
		RecordingSid      *string      `json:"recording_sid"`
		Sid               *string      `json:"sid"`
		Status            *string      `json:"status"`
		TranscriptionText *string      `json:"transcription_text"`
		Type              *string      `json:"type"`
		Uri               *string      `json:"uri"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = ApiV2010Transcription{
		AccountSid:        raw.AccountSid,
		ApiVersion:        raw.ApiVersion,
		DateCreated:       raw.DateCreated,
		DateUpdated:       raw.DateUpdated,
		Duration:          raw.Duration,
		PriceUnit:         raw.PriceUnit,
		RecordingSid:      raw.RecordingSid,
		Sid:               raw.Sid,
		Status:            raw.Status,
		TranscriptionText: raw.TranscriptionText,
		Type:              raw.Type,
		Uri:               raw.Uri,
	}

	responsePrice, err := client.UnmarshalFloat32(raw.Price)
	if err != nil {
		return err
	}
	response.Price = responsePrice

	return
}
