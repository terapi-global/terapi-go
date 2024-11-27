# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthAuthenticatePost**](AuthApi.md#AuthAuthenticatePost) | **Post** /auth/authenticate | 23_999
[**AuthRefreshPost**](AuthApi.md#AuthRefreshPost) | **Post** /auth/refresh | 23_998

# **AuthAuthenticatePost**
> GetAuthenticationTokenResponse AuthAuthenticatePost(ctx, optional)
23_999

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 23_999_101 | UnknownExceptionWhenTryingFetchApplicationApiKey |  |  | 23_999_102 | FailedToFindApplicationApiKey |  |  | 23_999_103 | UnknownExceptionWhenTryingToGetUserAuthenticationToken |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthApiAuthAuthenticatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthApiAuthAuthenticatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secretKey** | **optional.String**|  | 

### Return type

[**GetAuthenticationTokenResponse**](GetAuthenticationTokenResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthRefreshPost**
> RefreshAuthenticationTokenResponse AuthRefreshPost(ctx, optional)
23_998

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 23_998_101 | UnknownExceptionWhenTryingToRefreshUserAuthenticationToken |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthApiAuthRefreshPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthApiAuthRefreshPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshToken** | **optional.String**|  | 

### Return type

[**RefreshAuthenticationTokenResponse**](RefreshAuthenticationTokenResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

