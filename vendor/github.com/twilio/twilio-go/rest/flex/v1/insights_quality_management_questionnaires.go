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

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateInsightsQuestionnaires'
type CreateInsightsQuestionnairesParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The name of this questionnaire
	Name *string `json:"Name,omitempty"`
	// The description of this questionnaire
	Description *string `json:"Description,omitempty"`
	// The flag to enable or disable questionnaire
	Active *bool `json:"Active,omitempty"`
	// The list of questions sids under a questionnaire
	QuestionSids *[]string `json:"QuestionSids,omitempty"`
}

func (params *CreateInsightsQuestionnairesParams) SetAuthorization(Authorization string) *CreateInsightsQuestionnairesParams {
	params.Authorization = &Authorization
	return params
}
func (params *CreateInsightsQuestionnairesParams) SetName(Name string) *CreateInsightsQuestionnairesParams {
	params.Name = &Name
	return params
}
func (params *CreateInsightsQuestionnairesParams) SetDescription(Description string) *CreateInsightsQuestionnairesParams {
	params.Description = &Description
	return params
}
func (params *CreateInsightsQuestionnairesParams) SetActive(Active bool) *CreateInsightsQuestionnairesParams {
	params.Active = &Active
	return params
}
func (params *CreateInsightsQuestionnairesParams) SetQuestionSids(QuestionSids []string) *CreateInsightsQuestionnairesParams {
	params.QuestionSids = &QuestionSids
	return params
}

// To create a Questionnaire
func (c *ApiService) CreateInsightsQuestionnaires(params *CreateInsightsQuestionnairesParams) (*FlexV1InsightsQuestionnaires, error) {
	path := "/v1/Insights/QualityManagement/Questionnaires"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Name != nil {
		data.Set("Name", *params.Name)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.Active != nil {
		data.Set("Active", fmt.Sprint(*params.Active))
	}
	if params != nil && params.QuestionSids != nil {
		for _, item := range *params.QuestionSids {
			data.Add("QuestionSids", item)
		}
	}

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InsightsQuestionnaires{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteInsightsQuestionnaires'
type DeleteInsightsQuestionnairesParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
}

func (params *DeleteInsightsQuestionnairesParams) SetAuthorization(Authorization string) *DeleteInsightsQuestionnairesParams {
	params.Authorization = &Authorization
	return params
}

// To delete the questionnaire
func (c *ApiService) DeleteInsightsQuestionnaires(QuestionnaireSid string, params *DeleteInsightsQuestionnairesParams) error {
	path := "/v1/Insights/QualityManagement/Questionnaires/{QuestionnaireSid}"
	path = strings.Replace(path, "{"+"QuestionnaireSid"+"}", QuestionnaireSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchInsightsQuestionnaires'
type FetchInsightsQuestionnairesParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
}

func (params *FetchInsightsQuestionnairesParams) SetAuthorization(Authorization string) *FetchInsightsQuestionnairesParams {
	params.Authorization = &Authorization
	return params
}

// To get the Questionnaire Detail
func (c *ApiService) FetchInsightsQuestionnaires(QuestionnaireSid string, params *FetchInsightsQuestionnairesParams) (*FlexV1InsightsQuestionnaires, error) {
	path := "/v1/Insights/QualityManagement/Questionnaires/{QuestionnaireSid}"
	path = strings.Replace(path, "{"+"QuestionnaireSid"+"}", QuestionnaireSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InsightsQuestionnaires{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInsightsQuestionnaires'
type ListInsightsQuestionnairesParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// Flag indicating whether to include inactive questionnaires or not
	IncludeInactive *bool `json:"IncludeInactive,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListInsightsQuestionnairesParams) SetAuthorization(Authorization string) *ListInsightsQuestionnairesParams {
	params.Authorization = &Authorization
	return params
}
func (params *ListInsightsQuestionnairesParams) SetIncludeInactive(IncludeInactive bool) *ListInsightsQuestionnairesParams {
	params.IncludeInactive = &IncludeInactive
	return params
}
func (params *ListInsightsQuestionnairesParams) SetPageSize(PageSize int) *ListInsightsQuestionnairesParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInsightsQuestionnairesParams) SetLimit(Limit int) *ListInsightsQuestionnairesParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of InsightsQuestionnaires records from the API. Request is executed immediately.
func (c *ApiService) PageInsightsQuestionnaires(params *ListInsightsQuestionnairesParams, pageToken, pageNumber string) (*ListInsightsQuestionnairesResponse, error) {
	path := "/v1/Insights/QualityManagement/Questionnaires"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IncludeInactive != nil {
		data.Set("IncludeInactive", fmt.Sprint(*params.IncludeInactive))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInsightsQuestionnairesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists InsightsQuestionnaires records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInsightsQuestionnaires(params *ListInsightsQuestionnairesParams) ([]FlexV1InsightsQuestionnaires, error) {
	response, errors := c.StreamInsightsQuestionnaires(params)

	records := make([]FlexV1InsightsQuestionnaires, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InsightsQuestionnaires records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInsightsQuestionnaires(params *ListInsightsQuestionnairesParams) (chan FlexV1InsightsQuestionnaires, chan error) {
	if params == nil {
		params = &ListInsightsQuestionnairesParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1InsightsQuestionnaires, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInsightsQuestionnaires(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInsightsQuestionnaires(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamInsightsQuestionnaires(response *ListInsightsQuestionnairesResponse, params *ListInsightsQuestionnairesParams, recordChannel chan FlexV1InsightsQuestionnaires, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Questionnaires
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListInsightsQuestionnairesResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInsightsQuestionnairesResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInsightsQuestionnairesResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInsightsQuestionnairesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateInsightsQuestionnaires'
type UpdateInsightsQuestionnairesParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The flag to enable or disable questionnaire
	Active *bool `json:"Active,omitempty"`
	// The name of this questionnaire
	Name *string `json:"Name,omitempty"`
	// The description of this questionnaire
	Description *string `json:"Description,omitempty"`
	// The list of questions sids under a questionnaire
	QuestionSids *[]string `json:"QuestionSids,omitempty"`
}

func (params *UpdateInsightsQuestionnairesParams) SetAuthorization(Authorization string) *UpdateInsightsQuestionnairesParams {
	params.Authorization = &Authorization
	return params
}
func (params *UpdateInsightsQuestionnairesParams) SetActive(Active bool) *UpdateInsightsQuestionnairesParams {
	params.Active = &Active
	return params
}
func (params *UpdateInsightsQuestionnairesParams) SetName(Name string) *UpdateInsightsQuestionnairesParams {
	params.Name = &Name
	return params
}
func (params *UpdateInsightsQuestionnairesParams) SetDescription(Description string) *UpdateInsightsQuestionnairesParams {
	params.Description = &Description
	return params
}
func (params *UpdateInsightsQuestionnairesParams) SetQuestionSids(QuestionSids []string) *UpdateInsightsQuestionnairesParams {
	params.QuestionSids = &QuestionSids
	return params
}

// To update the questionnaire
func (c *ApiService) UpdateInsightsQuestionnaires(QuestionnaireSid string, params *UpdateInsightsQuestionnairesParams) (*FlexV1InsightsQuestionnaires, error) {
	path := "/v1/Insights/QualityManagement/Questionnaires/{QuestionnaireSid}"
	path = strings.Replace(path, "{"+"QuestionnaireSid"+"}", QuestionnaireSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Active != nil {
		data.Set("Active", fmt.Sprint(*params.Active))
	}
	if params != nil && params.Name != nil {
		data.Set("Name", *params.Name)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.QuestionSids != nil {
		for _, item := range *params.QuestionSids {
			data.Add("QuestionSids", item)
		}
	}

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InsightsQuestionnaires{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
