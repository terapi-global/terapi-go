# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationApikeyGeneratePost**](ApplicationApiKeyApi.md#ApplicationApikeyGeneratePost) | **Post** /application-apikey/generate | 16_999
[**ApplicationApikeyGetByApplicationIdGet**](ApplicationApiKeyApi.md#ApplicationApikeyGetByApplicationIdGet) | **Get** /application-apikey/get-by-application-id | 16_995
[**ApplicationApikeyRevokeDelete**](ApplicationApiKeyApi.md#ApplicationApikeyRevokeDelete) | **Delete** /application-apikey/revoke | 16_998

# **ApplicationApikeyGeneratePost**
> CreateApplicationApiKeyResponse ApplicationApikeyGeneratePost(ctx, optional)
16_999

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 16_999_101 | ApplicationNotFound | Application not found. |  | 16_999_102 | UnknownExceptionWhenCheckApplicationExistence | Unknown exception when checking if application exists. |  | 16_999_103 | UnknownExceptionWhenGenerateApplicationApiKey | Unknown exception when generating application API key. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiKeyApiApplicationApikeyGeneratePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiKeyApiApplicationApikeyGeneratePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApplicationapikeyGenerateBody**](ApplicationapikeyGenerateBody.md)|  | 

### Return type

[**CreateApplicationApiKeyResponse**](CreateApplicationApiKeyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationApikeyGetByApplicationIdGet**
> ApplicationApiKeyByApplicationIdResponse ApplicationApikeyGetByApplicationIdGet(ctx, optional)
16_995

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 16_995_101 | ApplicationApiKeyNotFound | Application API key not found. |  | 16_995_102 | UnknownExceptionWhenTryingFetchApplications | Unknown exception when trying to fetch applications. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiKeyApiApplicationApikeyGetByApplicationIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiKeyApiApplicationApikeyGetByApplicationIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | [**optional.Interface of string**](.md)|  | 

### Return type

[**ApplicationApiKeyByApplicationIdResponse**](ApplicationApiKeyByApplicationIdResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationApikeyRevokeDelete**
> RevokeApplicationApiKeyResponse ApplicationApikeyRevokeDelete(ctx, optional)
16_998

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 16_998_101 | ApplicationApiKeyNotFound | Application API key not found. |  | 16_998_102 | UnknownExceptionWhenCheckApplicationExistence | Unknown exception when checking if application exists. |  | 16_998_103 | UnknownExceptionWhenGenerateApplicationApiKey | Unknown exception when generating application API key. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiKeyApiApplicationApikeyRevokeDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiKeyApiApplicationApikeyRevokeDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**optional.Interface of string**](.md)|  | 

### Return type

[**RevokeApplicationApiKeyResponse**](RevokeApplicationApiKeyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

