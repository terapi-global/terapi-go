
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ApiProxyApiService service
/*
ApiProxyApiService 21_997
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xTenantId
 * @param integrationName
 * @param endpoint
 * @param optional nil or *ApiProxyApiApiProxyEndpointDeleteOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ApiproxyEndpointBody6) - 
@return CallActionResponse
*/

type ApiProxyApiApiProxyEndpointDeleteOpts struct {
    Body optional.Interface
}

func (a *ApiProxyApiService) ApiProxyEndpointDelete(ctx context.Context, xTenantId string, integrationName string, endpoint string, localVarOptionals *ApiProxyApiApiProxyEndpointDeleteOpts) (CallActionResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CallActionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api-proxy/{endpoint}"
	localVarPath = strings.Replace(localVarPath, "{"+"integrationName"+"}", fmt.Sprintf("%v", integrationName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpoint"+"}", fmt.Sprintf("%v", endpoint), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["X-Tenant-Id"] = parameterToString(xTenantId, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CallActionResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ApiProxyApiService 21_997
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param integrationName
 * @param endpoint
 * @param xTenantId
@return CallActionResponse
*/
func (a *ApiProxyApiService) ApiProxyEndpointGet(ctx context.Context, integrationName string, endpoint string, xTenantId string) (CallActionResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CallActionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api-proxy/{endpoint}"
	localVarPath = strings.Replace(localVarPath, "{"+"integrationName"+"}", fmt.Sprintf("%v", integrationName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpoint"+"}", fmt.Sprintf("%v", endpoint), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["X-Tenant-Id"] = parameterToString(xTenantId, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CallActionResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ApiProxyApiService 21_997
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xTenantId
 * @param integrationName
 * @param endpoint
 * @param optional nil or *ApiProxyApiApiProxyEndpointPatchOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ApiproxyEndpointBody9) - 
@return CallActionResponse
*/

type ApiProxyApiApiProxyEndpointPatchOpts struct {
    Body optional.Interface
}

func (a *ApiProxyApiService) ApiProxyEndpointPatch(ctx context.Context, xTenantId string, integrationName string, endpoint string, localVarOptionals *ApiProxyApiApiProxyEndpointPatchOpts) (CallActionResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CallActionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api-proxy/{endpoint}"
	localVarPath = strings.Replace(localVarPath, "{"+"integrationName"+"}", fmt.Sprintf("%v", integrationName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpoint"+"}", fmt.Sprintf("%v", endpoint), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["X-Tenant-Id"] = parameterToString(xTenantId, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CallActionResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ApiProxyApiService 21_997
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xTenantId
 * @param integrationName
 * @param endpoint
 * @param optional nil or *ApiProxyApiApiProxyEndpointPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ApiproxyEndpointBody3) - 
@return CallActionResponse
*/

type ApiProxyApiApiProxyEndpointPostOpts struct {
    Body optional.Interface
}

func (a *ApiProxyApiService) ApiProxyEndpointPost(ctx context.Context, xTenantId string, integrationName string, endpoint string, localVarOptionals *ApiProxyApiApiProxyEndpointPostOpts) (CallActionResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CallActionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api-proxy/{endpoint}"
	localVarPath = strings.Replace(localVarPath, "{"+"integrationName"+"}", fmt.Sprintf("%v", integrationName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpoint"+"}", fmt.Sprintf("%v", endpoint), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["X-Tenant-Id"] = parameterToString(xTenantId, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CallActionResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ApiProxyApiService 21_997
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xTenantId
 * @param integrationName
 * @param endpoint
 * @param optional nil or *ApiProxyApiApiProxyEndpointPutOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ApiproxyEndpointBody) - 
@return CallActionResponse
*/

type ApiProxyApiApiProxyEndpointPutOpts struct {
    Body optional.Interface
}

func (a *ApiProxyApiService) ApiProxyEndpointPut(ctx context.Context, xTenantId string, integrationName string, endpoint string, localVarOptionals *ApiProxyApiApiProxyEndpointPutOpts) (CallActionResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CallActionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api-proxy/{endpoint}"
	localVarPath = strings.Replace(localVarPath, "{"+"integrationName"+"}", fmt.Sprintf("%v", integrationName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpoint"+"}", fmt.Sprintf("%v", endpoint), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["X-Tenant-Id"] = parameterToString(xTenantId, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CallActionResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ApiProxyApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ApiProxyApiApiProxyOauth2ChallengeGetOpts - Optional Parameters:
     * @param "TenantId" (optional.Interface of string) - 
@return OAuth2ChallengeDtoResponse
*/

type ApiProxyApiApiProxyOauth2ChallengeGetOpts struct {
    TenantId optional.Interface
}

func (a *ApiProxyApiService) ApiProxyOauth2ChallengeGet(ctx context.Context, localVarOptionals *ApiProxyApiApiProxyOauth2ChallengeGetOpts) (OAuth2ChallengeDtoResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue OAuth2ChallengeDtoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api-proxy/oauth2/challenge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.TenantId.IsSet() {
		localVarQueryParams.Add("tenantId", parameterToString(localVarOptionals.TenantId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v OAuth2ChallengeDtoResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ApiProxyApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param externalId
@return UnifiedContactResponse
*/
func (a *ApiProxyApiService) ApiProxyUnifiedContactsExternalIdGet(ctx context.Context, externalId string) (UnifiedContactResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UnifiedContactResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api-proxy/unified/contacts/{externalId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalId"+"}", fmt.Sprintf("%v", externalId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v UnifiedContactResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
ApiProxyApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return UnifiedContactsListResponse
*/
func (a *ApiProxyApiService) ApiProxyUnifiedContactsGet(ctx context.Context) (UnifiedContactsListResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UnifiedContactsListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api-proxy/unified/contacts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v UnifiedContactsListResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
