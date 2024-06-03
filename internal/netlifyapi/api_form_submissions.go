/*
Netlify's API documentation

Netlify is a hosting service for the programmable web. It understands your documents and provides an API to handle atomic deploys of websites, manage form submissions, inject JavaScript snippets, and much more. This is a REST-style API that uses JSON for serialization and OAuth 2 for authentication.   This document is an OpenAPI reference for the Netlify API that you can explore. For more detailed instructions for common uses, please visit the [online documentation](https://docs.netlify.com/api/get-started/). Visit our Community forum to join the conversation about [understanding and using Netlify’s API](https://community.netlify.com/t/common-issue-understanding-and-using-netlifys-api/160).   Additionally, we have two API clients for your convenience: - [Go Client](https://github.com/netlify/open-api#go-client) - [JS Client](https://github.com/netlify/js-client) 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netlifyapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// FormSubmissionsAPIService FormSubmissionsAPI service
type FormSubmissionsAPIService service

type ApiDeleteSubmissionRequest struct {
	ctx context.Context
	ApiService *FormSubmissionsAPIService
	submissionId string
}

func (r ApiDeleteSubmissionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSubmissionExecute(r)
}

/*
DeleteSubmission Method for DeleteSubmission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param submissionId The submission ID
 @return ApiDeleteSubmissionRequest
*/
func (a *FormSubmissionsAPIService) DeleteSubmission(ctx context.Context, submissionId string) ApiDeleteSubmissionRequest {
	return ApiDeleteSubmissionRequest{
		ApiService: a,
		ctx: ctx,
		submissionId: submissionId,
	}
}

// Execute executes the request
func (a *FormSubmissionsAPIService) DeleteSubmissionExecute(r ApiDeleteSubmissionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FormSubmissionsAPIService.DeleteSubmission")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/submissions/{submission_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"submission_id"+"}", url.PathEscape(parameterValueToString(r.submissionId, "submissionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListFormSubmissionRequest struct {
	ctx context.Context
	ApiService *FormSubmissionsAPIService
	submissionId string
	page *string
	perPage *string
}

// Number of per_page to skip when returning records
func (r ApiListFormSubmissionRequest) Page(page string) ApiListFormSubmissionRequest {
	r.page = &page
	return r
}

// Number of records to return
func (r ApiListFormSubmissionRequest) PerPage(perPage string) ApiListFormSubmissionRequest {
	r.perPage = &perPage
	return r
}

func (r ApiListFormSubmissionRequest) Execute() (*FormSubmission, *http.Response, error) {
	return r.ApiService.ListFormSubmissionExecute(r)
}

/*
ListFormSubmission Method for ListFormSubmission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param submissionId The submission ID
 @return ApiListFormSubmissionRequest
*/
func (a *FormSubmissionsAPIService) ListFormSubmission(ctx context.Context, submissionId string) ApiListFormSubmissionRequest {
	return ApiListFormSubmissionRequest{
		ApiService: a,
		ctx: ctx,
		submissionId: submissionId,
	}
}

// Execute executes the request
//  @return FormSubmission
func (a *FormSubmissionsAPIService) ListFormSubmissionExecute(r ApiListFormSubmissionRequest) (*FormSubmission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FormSubmission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FormSubmissionsAPIService.ListFormSubmission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/submissions/{submission_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"submission_id"+"}", url.PathEscape(parameterValueToString(r.submissionId, "submissionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListFormSubmissionsRequest struct {
	ctx context.Context
	ApiService *FormSubmissionsAPIService
	formId string
	page *string
	perPage *string
}

// Number of per_page to skip when returning records
func (r ApiListFormSubmissionsRequest) Page(page string) ApiListFormSubmissionsRequest {
	r.page = &page
	return r
}

// Number of records to return
func (r ApiListFormSubmissionsRequest) PerPage(perPage string) ApiListFormSubmissionsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiListFormSubmissionsRequest) Execute() ([]FormSubmission, *http.Response, error) {
	return r.ApiService.ListFormSubmissionsExecute(r)
}

/*
ListFormSubmissions Method for ListFormSubmissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param formId The form ID
 @return ApiListFormSubmissionsRequest
*/
func (a *FormSubmissionsAPIService) ListFormSubmissions(ctx context.Context, formId string) ApiListFormSubmissionsRequest {
	return ApiListFormSubmissionsRequest{
		ApiService: a,
		ctx: ctx,
		formId: formId,
	}
}

// Execute executes the request
//  @return []FormSubmission
func (a *FormSubmissionsAPIService) ListFormSubmissionsExecute(r ApiListFormSubmissionsRequest) ([]FormSubmission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []FormSubmission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FormSubmissionsAPIService.ListFormSubmissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/form/{form_id}/submissions"
	localVarPath = strings.Replace(localVarPath, "{"+"form_id"+"}", url.PathEscape(parameterValueToString(r.formId, "formId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListSiteSubmissionsRequest struct {
	ctx context.Context
	ApiService *FormSubmissionsAPIService
	siteId string
	page *string
	perPage *string
}

// Number of per_page to skip when returning records
func (r ApiListSiteSubmissionsRequest) Page(page string) ApiListSiteSubmissionsRequest {
	r.page = &page
	return r
}

// Number of records to return
func (r ApiListSiteSubmissionsRequest) PerPage(perPage string) ApiListSiteSubmissionsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiListSiteSubmissionsRequest) Execute() ([]FormSubmission, *http.Response, error) {
	return r.ApiService.ListSiteSubmissionsExecute(r)
}

/*
ListSiteSubmissions Method for ListSiteSubmissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId The site ID
 @return ApiListSiteSubmissionsRequest
*/
func (a *FormSubmissionsAPIService) ListSiteSubmissions(ctx context.Context, siteId string) ApiListSiteSubmissionsRequest {
	return ApiListSiteSubmissionsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return []FormSubmission
func (a *FormSubmissionsAPIService) ListSiteSubmissionsExecute(r ApiListSiteSubmissionsRequest) ([]FormSubmission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []FormSubmission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FormSubmissionsAPIService.ListSiteSubmissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/submissions"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}