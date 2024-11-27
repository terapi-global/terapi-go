

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

type TenantApiService service
/*
TenantApiService 20_996
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_996_101 | TenantNotFound | Tenant Not Found |  | 20_996_102 | UnknownExceptionWhileFetchingTenant | Unknown exception when checking tenant existence. |  | 20_996_103 | UnknownExceptionWhileDecliningTenant | Unknown exception when declining tenant. |  | 20_996_104 | TenantIsNotInInvitedStatus |  |  | 20_996_105 | UnknownExceptionWhileCheckingIfTenantIsInCorrectStatus |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
@return DeclineTenantInvitationResponse
*/
func (a *TenantApiService) TenantDeclineTenantIdGet(ctx context.Context, id string) (DeclineTenantInvitationResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DeclineTenantInvitationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tenant/decline-tenant/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", fmt.Sprintf("%v", id), -1)

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
			var v DeclineTenantInvitationResponse
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
TenantApiService 20_998
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_998_101 | ApplicationIntegrationNotFound | Application integration not found. |  | 20_998_102 | UnknownExceptionWhenCheckApplicationExistence | Unknown exception when checking application existence. |  | 20_998_103 | UnknownExceptionWhenCreateTenant | Unknown exception when createing tenant. |  | 20_998_104 | UnknownExceptionWhenSendingEmailToInvitedEmailAddress | Unknown exception when sending email to invited email address. |  | 20_998_105 | UserNotFound | User not found. |  | 20_998_106 | UnknownExceptionWhenInsertApplication | Unknown exception when inserting application. |  | 20_998_107 | UnknownExceptionWhenCreateUserTenant | Unknown exception when creating user tenant. |  | 20_998_108 | UnknownExceptionWhenCreatingEmailContent | Unknown exception when creating email content. |  | 20_998_109 | UnknownExceptionWhenCreatingEmailAcceptInvitationUrl |  |  | 20_998_110 | UnknownExceptionWhenAddingTenantInKeycloak |  |  | 20_998_111 | UnknownExceptionWhenAddingTenantInOrganization |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TenantApiTenantInviteTenantByApplicationIntegrationIdPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of TenantInvitetenantbyapplicationintegrationidBody) - 
@return InviteTenantByApplicationIntegrationIdResponse
*/

type TenantApiTenantInviteTenantByApplicationIntegrationIdPostOpts struct {
    Body optional.Interface
}

func (a *TenantApiService) TenantInviteTenantByApplicationIntegrationIdPost(ctx context.Context, localVarOptionals *TenantApiTenantInviteTenantByApplicationIntegrationIdPostOpts) (InviteTenantByApplicationIntegrationIdResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InviteTenantByApplicationIntegrationIdResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tenant/invite-tenant-by-application-integration-id"

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
			var v InviteTenantByApplicationIntegrationIdResponse
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
TenantApiService 20_998
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_998_101 | ApplicationIntegrationNotFound | Application integration not found. |  | 20_998_102 | UnknownExceptionWhenCheckApplicationExistence | Unknown exception when checking application existence. |  | 20_998_103 | UnknownExceptionWhenCreateTenant | Unknown exception when createing tenant. |  | 20_998_104 | UnknownExceptionWhenSendingEmailToInvitedEmailAddress | Unknown exception when sending email to invited email address. |  | 20_998_105 | UserNotFound | User not found. |  | 20_998_106 | UnknownExceptionWhenInsertApplication | Unknown exception when inserting application. |  | 20_998_107 | UnknownExceptionWhenCreateUserTenant | Unknown exception when creating user tenant. |  | 20_998_108 | UnknownExceptionWhenCreatingEmailContent | Unknown exception when creating email content. |  | 20_998_109 | UnknownExceptionWhenCreatingEmailAcceptInvitationUrl |  |  | 20_998_110 | UnknownExceptionWhenAddingTenantInKeycloak |  |  | 20_998_111 | UnknownExceptionWhenAddingTenantInOrganization |  |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TenantApiTenantInviteTenantPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of TenantInvitetenantBody) - 
@return InviteTenantResponse
*/

type TenantApiTenantInviteTenantPostOpts struct {
    Body optional.Interface
}

func (a *TenantApiService) TenantInviteTenantPost(ctx context.Context, localVarOptionals *TenantApiTenantInviteTenantPostOpts) (InviteTenantResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InviteTenantResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tenant/invite-tenant"

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
			var v InviteTenantResponse
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
TenantApiService 20_997
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_997_101 | TenantNotFound | Tenant not found. |  | 20_997_102 | UnknownExceptionWhenCheckTenantExistence | Unknown exception when checking tenant existence. |  | 20_997_103 | UnknownExceptionWhenUpdateTenantConnection | Unknown exception when updating tenant connection. |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TenantApiTenantUpdateTenantConnectionPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of TenantUpdatetenantconnectionBody) - 
@return UpdateTenantConnectionResponse
*/

type TenantApiTenantUpdateTenantConnectionPostOpts struct {
    Body optional.Interface
}

func (a *TenantApiService) TenantUpdateTenantConnectionPost(ctx context.Context, localVarOptionals *TenantApiTenantUpdateTenantConnectionPostOpts) (UpdateTenantConnectionResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UpdateTenantConnectionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tenant/update-tenant-connection"

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
			var v UpdateTenantConnectionResponse
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
TenantApiService 20_999
### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_999_101 | UnknownExceptionWhenTryingFetchTenants | Unknown exception when trying to fetch tenants. |  | 20_999_102 | UserNotFound | User not found. |  | 20_999_103 | UnknownExceptionWhenInsertApplication | Unknown exception when inserting application. |
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TenantApiTenantUserTenantListGetOpts - Optional Parameters:
     * @param "InvitationStatus" (optional.Interface of InvitationStatus1) -    0 &#x3D; Invited  1 &#x3D; Accepted  2 &#x3D; Declined
     * @param "Page" (optional.Int32) - 
     * @param "PerPage" (optional.Int32) - 
@return UserTenantListResponse
*/

type TenantApiTenantUserTenantListGetOpts struct {
    InvitationStatus optional.Interface
    Page optional.Int32
    PerPage optional.Int32
}

func (a *TenantApiService) TenantUserTenantListGet(ctx context.Context, localVarOptionals *TenantApiTenantUserTenantListGetOpts) (UserTenantListResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue UserTenantListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tenant/user-tenant-list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.InvitationStatus.IsSet() {
		localVarQueryParams.Add("InvitationStatus", parameterToString(localVarOptionals.InvitationStatus.Value(), ""))
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
			var v UserTenantListResponse
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
