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


// SiteAssetsAPIService SiteAssetsAPI service
type SiteAssetsAPIService service

type ApiCreateSiteAssetRequest struct {
	ctx context.Context
	ApiService *SiteAssetsAPIService
	siteId string
	assetCreateParams *AssetCreateParams
}

// 
func (r ApiCreateSiteAssetRequest) AssetCreateParams(assetCreateParams AssetCreateParams) ApiCreateSiteAssetRequest {
	r.assetCreateParams = &assetCreateParams
	return r
}

func (r ApiCreateSiteAssetRequest) Execute() (*AssetCreateResult, *http.Response, error) {
	return r.ApiService.CreateSiteAssetExecute(r)
}

/*
CreateSiteAsset Method for CreateSiteAsset

Creates an asset for the site.
_This endpoint is only available to certain sites._

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId The site ID
 @return ApiCreateSiteAssetRequest
*/
func (a *SiteAssetsAPIService) CreateSiteAsset(ctx context.Context, siteId string) ApiCreateSiteAssetRequest {
	return ApiCreateSiteAssetRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return AssetCreateResult
func (a *SiteAssetsAPIService) CreateSiteAssetExecute(r ApiCreateSiteAssetRequest) (*AssetCreateResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssetCreateResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SiteAssetsAPIService.CreateSiteAsset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/assets"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetCreateParams == nil {
		return localVarReturnValue, nil, reportError("assetCreateParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.assetCreateParams
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

type ApiDeleteSiteAssetRequest struct {
	ctx context.Context
	ApiService *SiteAssetsAPIService
	assetId string
	siteId string
}

func (r ApiDeleteSiteAssetRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSiteAssetExecute(r)
}

/*
DeleteSiteAsset Method for DeleteSiteAsset

Deletes an asset.
_This endpoint is only available to certain sites._

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetId The asset ID
 @param siteId The site ID
 @return ApiDeleteSiteAssetRequest
*/
func (a *SiteAssetsAPIService) DeleteSiteAsset(ctx context.Context, assetId string, siteId string) ApiDeleteSiteAssetRequest {
	return ApiDeleteSiteAssetRequest{
		ApiService: a,
		ctx: ctx,
		assetId: assetId,
		siteId: siteId,
	}
}

// Execute executes the request
func (a *SiteAssetsAPIService) DeleteSiteAssetExecute(r ApiDeleteSiteAssetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SiteAssetsAPIService.DeleteSiteAsset")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/assets/{asset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"asset_id"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

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

type ApiGetSiteAssetInfoRequest struct {
	ctx context.Context
	ApiService *SiteAssetsAPIService
	assetId string
	siteId string
}

func (r ApiGetSiteAssetInfoRequest) Execute() (*Asset, *http.Response, error) {
	return r.ApiService.GetSiteAssetInfoExecute(r)
}

/*
GetSiteAssetInfo Method for GetSiteAssetInfo

Returns an asset for the site.
_This endpoint is only available to certain sites._

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetId The asset ID
 @param siteId The site ID
 @return ApiGetSiteAssetInfoRequest
*/
func (a *SiteAssetsAPIService) GetSiteAssetInfo(ctx context.Context, assetId string, siteId string) ApiGetSiteAssetInfoRequest {
	return ApiGetSiteAssetInfoRequest{
		ApiService: a,
		ctx: ctx,
		assetId: assetId,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return Asset
func (a *SiteAssetsAPIService) GetSiteAssetInfoExecute(r ApiGetSiteAssetInfoRequest) (*Asset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Asset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SiteAssetsAPIService.GetSiteAssetInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/assets/{asset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"asset_id"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

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

type ApiGetSiteAssetPublicSignatureRequest struct {
	ctx context.Context
	ApiService *SiteAssetsAPIService
	assetId string
	siteId string
}

func (r ApiGetSiteAssetPublicSignatureRequest) Execute() (*GetSiteAssetPublicSignature200Response, *http.Response, error) {
	return r.ApiService.GetSiteAssetPublicSignatureExecute(r)
}

/*
GetSiteAssetPublicSignature Method for GetSiteAssetPublicSignature

Returns a URL of the asset. If the asset isn't publicly visible, returns
the signed URL of the asset.
_This endpoint is only available to certain sites._

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetId The asset ID
 @param siteId The site ID
 @return ApiGetSiteAssetPublicSignatureRequest
*/
func (a *SiteAssetsAPIService) GetSiteAssetPublicSignature(ctx context.Context, assetId string, siteId string) ApiGetSiteAssetPublicSignatureRequest {
	return ApiGetSiteAssetPublicSignatureRequest{
		ApiService: a,
		ctx: ctx,
		assetId: assetId,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return GetSiteAssetPublicSignature200Response
func (a *SiteAssetsAPIService) GetSiteAssetPublicSignatureExecute(r ApiGetSiteAssetPublicSignatureRequest) (*GetSiteAssetPublicSignature200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSiteAssetPublicSignature200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SiteAssetsAPIService.GetSiteAssetPublicSignature")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/assets/{asset_id}/public_signature"
	localVarPath = strings.Replace(localVarPath, "{"+"asset_id"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

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

type ApiListSiteAssetsRequest struct {
	ctx context.Context
	ApiService *SiteAssetsAPIService
	siteId string
	filter *string
	search *string
}

// filter
func (r ApiListSiteAssetsRequest) Filter(filter string) ApiListSiteAssetsRequest {
	r.filter = &filter
	return r
}

// search
func (r ApiListSiteAssetsRequest) Search(search string) ApiListSiteAssetsRequest {
	r.search = &search
	return r
}

func (r ApiListSiteAssetsRequest) Execute() ([]Asset, *http.Response, error) {
	return r.ApiService.ListSiteAssetsExecute(r)
}

/*
ListSiteAssets Method for ListSiteAssets

Returns a list of assets for the site.
_This endpoint is only available to certain sites._

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId The site ID
 @return ApiListSiteAssetsRequest
*/
func (a *SiteAssetsAPIService) ListSiteAssets(ctx context.Context, siteId string) ApiListSiteAssetsRequest {
	return ApiListSiteAssetsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return []Asset
func (a *SiteAssetsAPIService) ListSiteAssetsExecute(r ApiListSiteAssetsRequest) ([]Asset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Asset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SiteAssetsAPIService.ListSiteAssets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/assets"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
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

type ApiUpdateSiteAssetRequest struct {
	ctx context.Context
	ApiService *SiteAssetsAPIService
	assetId string
	siteId string
	state *string
}

// The state of the asset
func (r ApiUpdateSiteAssetRequest) State(state string) ApiUpdateSiteAssetRequest {
	r.state = &state
	return r
}

func (r ApiUpdateSiteAssetRequest) Execute() (*Asset, *http.Response, error) {
	return r.ApiService.UpdateSiteAssetExecute(r)
}

/*
UpdateSiteAsset Method for UpdateSiteAsset

Updates a state of the asset.
_This endpoint is only available to certain sites._

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assetId The asset ID
 @param siteId The site ID
 @return ApiUpdateSiteAssetRequest
*/
func (a *SiteAssetsAPIService) UpdateSiteAsset(ctx context.Context, assetId string, siteId string) ApiUpdateSiteAssetRequest {
	return ApiUpdateSiteAssetRequest{
		ApiService: a,
		ctx: ctx,
		assetId: assetId,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return Asset
func (a *SiteAssetsAPIService) UpdateSiteAssetExecute(r ApiUpdateSiteAssetRequest) (*Asset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Asset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SiteAssetsAPIService.UpdateSiteAsset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/assets/{asset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"asset_id"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.state == nil {
		return localVarReturnValue, nil, reportError("state is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "")
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
