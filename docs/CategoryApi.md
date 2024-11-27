# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoryListGet**](CategoryApi.md#CategoryListGet) | **Get** /category/list | 19_999

# **CategoryListGet**
> CategoryListResponse CategoryListGet(ctx, optional)
19_999

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 19_999_101 | UnknownExceptionWhenTryingFetchCategories | Unknown exception when fetching category. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CategoryApiCategoryListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoryApiCategoryListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **perPage** | **optional.Int32**|  | 

### Return type

[**CategoryListResponse**](CategoryListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

