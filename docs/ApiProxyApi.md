# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProxyEndpointDelete**](ApiProxyApi.md#ApiProxyEndpointDelete) | **Delete** /api-proxy/{endpoint} | 21_997
[**ApiProxyEndpointGet**](ApiProxyApi.md#ApiProxyEndpointGet) | **Get** /api-proxy/{endpoint} | 21_997
[**ApiProxyEndpointPatch**](ApiProxyApi.md#ApiProxyEndpointPatch) | **Patch** /api-proxy/{endpoint} | 21_997
[**ApiProxyEndpointPost**](ApiProxyApi.md#ApiProxyEndpointPost) | **Post** /api-proxy/{endpoint} | 21_997
[**ApiProxyEndpointPut**](ApiProxyApi.md#ApiProxyEndpointPut) | **Put** /api-proxy/{endpoint} | 21_997
[**ApiProxyOauth2ChallengeGet**](ApiProxyApi.md#ApiProxyOauth2ChallengeGet) | **Get** /api-proxy/oauth2/challenge | 
[**ApiProxyUnifiedContactsExternalIdGet**](ApiProxyApi.md#ApiProxyUnifiedContactsExternalIdGet) | **Get** /api-proxy/unified/contacts/{externalId} | 
[**ApiProxyUnifiedContactsGet**](ApiProxyApi.md#ApiProxyUnifiedContactsGet) | **Get** /api-proxy/unified/contacts | 

# **ApiProxyEndpointDelete**
> CallActionResponse ApiProxyEndpointDelete(ctx, integrationName, endpoint, xTenantId, optional)
21_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationName** | **string**|  | 
  **endpoint** | **string**|  | 
  **xTenantId** | [**string**](.md)|  | 
 **optional** | ***ApiProxyApiApiProxyEndpointDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiProxyApiApiProxyEndpointDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ApiproxyEndpointBody6**](ApiproxyEndpointBody6.md)|  | 

### Return type

[**CallActionResponse**](CallActionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxyEndpointGet**
> CallActionResponse ApiProxyEndpointGet(ctx, integrationName, endpoint, xTenantId)
21_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationName** | **string**|  | 
  **endpoint** | **string**|  | 
  **xTenantId** | [**string**](.md)|  | 

### Return type

[**CallActionResponse**](CallActionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxyEndpointPatch**
> CallActionResponse ApiProxyEndpointPatch(ctx, integrationName, endpoint, xTenantId, optional)
21_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationName** | **string**|  | 
  **endpoint** | **string**|  | 
  **xTenantId** | [**string**](.md)|  | 
 **optional** | ***ApiProxyApiApiProxyEndpointPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiProxyApiApiProxyEndpointPatchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ApiproxyEndpointBody9**](ApiproxyEndpointBody9.md)|  | 

### Return type

[**CallActionResponse**](CallActionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxyEndpointPost**
> CallActionResponse ApiProxyEndpointPost(ctx, integrationName, endpoint, xTenantId, optional)
21_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationName** | **string**|  | 
  **endpoint** | **string**|  | 
  **xTenantId** | [**string**](.md)|  | 
 **optional** | ***ApiProxyApiApiProxyEndpointPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiProxyApiApiProxyEndpointPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ApiproxyEndpointBody3**](ApiproxyEndpointBody3.md)|  | 

### Return type

[**CallActionResponse**](CallActionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxyEndpointPut**
> CallActionResponse ApiProxyEndpointPut(ctx, integrationName, endpoint, xTenantId, optional)
21_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 21_997_101 | TenantNotFound |  |  | 21_997_102 | IntegrationNotFound |  |  | 21_997_103 | UnknownExceptionWhenCallingThirdPartyApi |  |  | 21_997_104 | UnknownExceptionWhenCallAction |  |  | 21_997_105 | IntegrationIsNotEnabled |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationName** | **string**|  | 
  **endpoint** | **string**|  | 
  **xTenantId** | [**string**](.md)|  | 
 **optional** | ***ApiProxyApiApiProxyEndpointPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiProxyApiApiProxyEndpointPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of ApiproxyEndpointBody**](ApiproxyEndpointBody.md)|  | 

### Return type

[**CallActionResponse**](CallActionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxyOauth2ChallengeGet**
> OAuth2ChallengeDtoResponse ApiProxyOauth2ChallengeGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiProxyApiApiProxyOauth2ChallengeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApiProxyApiApiProxyOauth2ChallengeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**optional.Interface of string**](.md)|  | 

### Return type

[**OAuth2ChallengeDtoResponse**](OAuth2ChallengeDtoResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxyUnifiedContactsExternalIdGet**
> UnifiedContactResponse ApiProxyUnifiedContactsExternalIdGet(ctx, externalId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalId** | **string**|  | 

### Return type

[**UnifiedContactResponse**](UnifiedContactResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProxyUnifiedContactsGet**
> UnifiedContactsListResponse ApiProxyUnifiedContactsGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UnifiedContactsListResponse**](UnifiedContactsListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

