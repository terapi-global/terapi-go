

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

type ApplicationIntegrationApiService service
/*
ApplicationIntegrationApiService 18_999
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 18_999_101 | ApplicationNotFound | Application not found. |  | 18_999_102 | UnknownExceptionWhenCheckApplicationExistence | Unknown exception when checking if application exists. |  | 18_999_103 | IntegrationNotFound | Integration not found. |  | 18_999_104 | UnknownExceptionWhenCheckIntegrationExistence | Unknown exception when checking if integration exists. |  | 18_999_105 | UnknownExceptionWhenAddIntegrationToApplication | Unknown exception when adding integration to application. |  | 18_999_106 | IntegrationAlreadyAddedToTheApplication | You already map this integration to your application. |  | 18_999_107 | UnknownExceptionWhenCheckApplicationIntegrationUniqueness |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ApplicationIntegrationApiApplicationIntegrationAddIntegrationToApplicationPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of ApplicationintegrationAddintegrationtoapplicationBody) - 
@return AddIntegrationToApplicationResponse
*/

type ApplicationIntegrationApiApplicationIntegrationAddIntegrationToApplicationPostOpts struct {
    Body optional.Interface
}

func (a *ApplicationIntegrationApiService) ApplicationIntegrationAddIntegrationToApplicationPost(ctx context.Context, localVarOptionals *ApplicationIntegrationApiApplicationIntegrationAddIntegrationToApplicationPostOpts) (AddIntegrationToApplicationResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue AddIntegrationToApplicationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/application-integration/add-integration-to-application"

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
			var v AddIntegrationToApplicationResponse
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
ApplicationIntegrationApiService 18_997
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 18_998 | UnknownExceptionWhenTryingFetchApplicationIntegrations |  |  | 18_999 | ApplicationNotFound |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ApplicationIntegrationApiApplicationIntegrationApplicationIntegrationListByApplicationIdGetOpts - Optional Parameters:
     * @param "ApplicationId" (optional.Interface of string) - 
     * @param "Page" (optional.Int32) - 
     * @param "PerPage" (optional.Int32) - 
@return ApplicationIntegrationListByApplicationIdResponse
*/

type ApplicationIntegrationApiApplicationIntegrationApplicationIntegrationListByApplicationIdGetOpts struct {
    ApplicationId optional.Interface
    Page optional.Int32
    PerPage optional.Int32
}

func (a *ApplicationIntegrationApiService) ApplicationIntegrationApplicationIntegrationListByApplicationIdGet(ctx context.Context, localVarOptionals *ApplicationIntegrationApiApplicationIntegrationApplicationIntegrationListByApplicationIdGetOpts) (ApplicationIntegrationListByApplicationIdResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ApplicationIntegrationListByApplicationIdResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/application-integration/application-integration-list-by-application-id"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ApplicationId.IsSet() {
		localVarQueryParams.Add("ApplicationId", parameterToString(localVarOptionals.ApplicationId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("Page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("PerPage", parameterToString(localVarOptionals.PerPage.Value(), ""))
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
			var v ApplicationIntegrationListByApplicationIdResponse
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
