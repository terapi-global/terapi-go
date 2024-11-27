# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationGetGet**](IntegrationApi.md#IntegrationGetGet) | **Get** /integration/get | 17_997
[**IntegrationListGet**](IntegrationApi.md#IntegrationListGet) | **Get** /integration/list | 17_999

# **IntegrationGetGet**
> GetOneIntegrationResponse IntegrationGetGet(ctx, optional)
17_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 17_997_101 | UnknownExceptionWhenTryingFetchOneIntegration |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IntegrationApiIntegrationGetGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationApiIntegrationGetGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**optional.Interface of string**](.md)|  | 

### Return type

[**GetOneIntegrationResponse**](GetOneIntegrationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IntegrationListGet**
> IntegrationListResponse IntegrationListGet(ctx, optional)
17_999

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 17_999_101 | UnknownExceptionWhenTryingFetchIntegrations | Unknown exception when trying to fetch integrations. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IntegrationApiIntegrationListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationApiIntegrationListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **perPage** | **optional.Int32**|  | 
 **searchKeyword** | **optional.String**|  | 
 **categoryId** | [**optional.Interface of string**](.md)|  | 

### Return type

[**IntegrationListResponse**](IntegrationListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

