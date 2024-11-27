

package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type IntegrationApiService service
/*
IntegrationApiService 17_997
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 17_997_101 | UnknownExceptionWhenTryingFetchOneIntegration |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *IntegrationApiIntegrationGetGetOpts - Optional Parameters:
     * @param "Id" (optional.Interface of string) - 
@return GetOneIntegrationResponse
*/

type IntegrationApiIntegrationGetGetOpts struct {
    Id optional.Interface
}

func (a *IntegrationApiService) IntegrationGetGet(ctx context.Context, localVarOptionals *IntegrationApiIntegrationGetGetOpts) (GetOneIntegrationResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue GetOneIntegrationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/integration/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarQueryParams.Add("Id", parameterToString(localVarOptionals.Id.Value(), ""))
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
			var v GetOneIntegrationResponse
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
IntegrationApiService 17_999
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 17_999_101 | UnknownExceptionWhenTryingFetchIntegrations | Unknown exception when trying to fetch integrations. |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *IntegrationApiIntegrationListGetOpts - Optional Parameters:
     * @param "Page" (optional.Int32) - 
     * @param "PerPage" (optional.Int32) - 
     * @param "SearchKeyword" (optional.String) - 
     * @param "CategoryId" (optional.Interface of string) - 
@return IntegrationListResponse
*/

type IntegrationApiIntegrationListGetOpts struct {
    Page optional.Int32
    PerPage optional.Int32
    SearchKeyword optional.String
    CategoryId optional.Interface
}

func (a *IntegrationApiService) IntegrationListGet(ctx context.Context, localVarOptionals *IntegrationApiIntegrationListGetOpts) (IntegrationListResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue IntegrationListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/integration/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("Page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("PerPage", parameterToString(localVarOptionals.PerPage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchKeyword.IsSet() {
		localVarQueryParams.Add("SearchKeyword", parameterToString(localVarOptionals.SearchKeyword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CategoryId.IsSet() {
		localVarQueryParams.Add("CategoryId", parameterToString(localVarOptionals.CategoryId.Value(), ""))
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
			var v IntegrationListResponse
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
