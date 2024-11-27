# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationCreatePost**](ApplicationApi.md#ApplicationCreatePost) | **Post** /application/create | 15_999
[**ApplicationDeleteDelete**](ApplicationApi.md#ApplicationDeleteDelete) | **Delete** /application/delete | 15_997
[**ApplicationListGet**](ApplicationApi.md#ApplicationListGet) | **Get** /application/list | 15_996
[**ApplicationUpdatePut**](ApplicationApi.md#ApplicationUpdatePut) | **Put** /application/update | 15_998

# **ApplicationCreatePost**
> CreateApplicationResponse ApplicationCreatePost(ctx, optional)
15_999

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 15_999_101 | ApplicationNameAlreadyExist | Application name already exists. |  | 15_999_102 | UnknownExceptionWhenValidateApplicationNameUniqueness | Unknown exception when validating application name uniqueness. |  | 15_999_103 | UnknownExceptionWhenInsertApplication | Unknown exception when inserting application. |  | 15_999_104 | UserNotFound | User not found. |  | 15_999_105 | UnknownExceptionWhenCreateOrganization |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiApplicationCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiApplicationCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApplicationCreateBody**](ApplicationCreateBody.md)|  | 

### Return type

[**CreateApplicationResponse**](CreateApplicationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationDeleteDelete**
> DeleteApplicationResponse ApplicationDeleteDelete(ctx, optional)
15_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 15_997_101 | ApplicationNotFound | Application not found. |  | 15_997_102 | UnknownExceptionWhenFetchApplicationExistence | Unknown exception when fetching application existence. |  | 15_997_103 | UnknownExceptionWhenDeleteApplication | Unknown exception when deleting application. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiApplicationDeleteDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiApplicationDeleteDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**optional.Interface of string**](.md)|  | 

### Return type

[**DeleteApplicationResponse**](DeleteApplicationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationListGet**
> UserApplicationListResponse ApplicationListGet(ctx, optional)
15_996

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 15_996_101 | UnknownExceptionWhenTryingFetchApplications | Unknown exception when trying to fetch applications. |  | 15_996_102 | UserNotFound | User not found. |  | 15_996_103 | UnknownExceptionWhenInsertApplication | Unknown exception when inserting application. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiApplicationListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiApplicationListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **perPage** | **optional.Int32**|  | 

### Return type

[**UserApplicationListResponse**](UserApplicationListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationUpdatePut**
> UpdateApplicationResponse ApplicationUpdatePut(ctx, optional)
15_998

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 15_998_101 | ApplicationNotFound | Application not found. |  | 15_998_102 | UnknownExceptionWhenFetchApplication | Unknown exception when fetching application. |  | 15_998_103 | UnknownExceptionWhenInsertApplication | Unknown exception when inserting application. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationApiApplicationUpdatePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationApiApplicationUpdatePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ApplicationUpdateBody**](ApplicationUpdateBody.md)|  | 

### Return type

[**UpdateApplicationResponse**](UpdateApplicationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

