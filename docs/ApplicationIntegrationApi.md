# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationIntegrationAddIntegrationToApplicationPost**](ApplicationIntegrationApi.md#ApplicationIntegrationAddIntegrationToApplicationPost) | **Post** /application-integration/add-integration-to-application | 18_999
[**ApplicationIntegrationApplicationIntegrationListByApplicationIdGet**](ApplicationIntegrationApi.md#ApplicationIntegrationApplicationIntegrationListByApplicationIdGet) | **Get** /application-integration/application-integration-list-by-application-id | 18_997

# **ApplicationIntegrationAddIntegrationToApplicationPost**
> AddIntegrationToApplicationResponse ApplicationIntegrationAddIntegrationToApplicationPost(ctx, optional)
18_999

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 18_999_101 | ApplicationNotFound | Application not found. |  | 18_999_102 | UnknownExceptionWhenCheckApplicationExistence | Unknown exception when checking if application exists. |  | 18_999_103 | IntegrationNotFound | Integration not found. |  | 18_999_104 | UnknownExceptionWhenCheckIntegrationExistence | Unknown exception when checking if integration exists. |  | 18_999_105 | UnknownExceptionWhenAddIntegrationToApplication | Unknown exception when adding integration to application. |  | 18_999_106 | IntegrationAlreadyAddedToTheApplication | You already map this integration to your application. |  | 18_999_107 | UnknownExceptionWhenCheckApplicationIntegrationUniqueness |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationIntegrationApiApplicationIntegrationAddIntegrationToApplicationPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationIntegrationApiApplicationIntegrationAddIntegrationToApplicationPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApplicationintegrationAddintegrationtoapplicationBody**](ApplicationintegrationAddintegrationtoapplicationBody.md)|  | 

### Return type

[**AddIntegrationToApplicationResponse**](AddIntegrationToApplicationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationIntegrationApplicationIntegrationListByApplicationIdGet**
> ApplicationIntegrationListByApplicationIdResponse ApplicationIntegrationApplicationIntegrationListByApplicationIdGet(ctx, optional)
18_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 18_998 | UnknownExceptionWhenTryingFetchApplicationIntegrations |  |  | 18_999 | ApplicationNotFound |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationIntegrationApiApplicationIntegrationApplicationIntegrationListByApplicationIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationIntegrationApiApplicationIntegrationApplicationIntegrationListByApplicationIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | [**optional.Interface of string**](.md)|  | 
 **page** | **optional.Int32**|  | 
 **perPage** | **optional.Int32**|  | 

### Return type

[**ApplicationIntegrationListByApplicationIdResponse**](ApplicationIntegrationListByApplicationIdResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

