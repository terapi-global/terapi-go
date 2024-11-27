# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantDeclineTenantIdGet**](TenantApi.md#TenantDeclineTenantIdGet) | **Get** /tenant/decline-tenant/{id} | 20_996
[**TenantInviteTenantByApplicationIntegrationIdPost**](TenantApi.md#TenantInviteTenantByApplicationIntegrationIdPost) | **Post** /tenant/invite-tenant-by-application-integration-id | 20_998
[**TenantInviteTenantPost**](TenantApi.md#TenantInviteTenantPost) | **Post** /tenant/invite-tenant | 20_998
[**TenantUpdateTenantConnectionPost**](TenantApi.md#TenantUpdateTenantConnectionPost) | **Post** /tenant/update-tenant-connection | 20_997
[**TenantUserTenantListGet**](TenantApi.md#TenantUserTenantListGet) | **Get** /tenant/user-tenant-list | 20_999

# **TenantDeclineTenantIdGet**
> DeclineTenantInvitationResponse TenantDeclineTenantIdGet(ctx, id)
20_996

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_996_101 | TenantNotFound | Tenant Not Found |  | 20_996_102 | UnknownExceptionWhileFetchingTenant | Unknown exception when checking tenant existence. |  | 20_996_103 | UnknownExceptionWhileDecliningTenant | Unknown exception when declining tenant. |  | 20_996_104 | TenantIsNotInInvitedStatus |  |  | 20_996_105 | UnknownExceptionWhileCheckingIfTenantIsInCorrectStatus |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**DeclineTenantInvitationResponse**](DeclineTenantInvitationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantInviteTenantByApplicationIntegrationIdPost**
> InviteTenantByApplicationIntegrationIdResponse TenantInviteTenantByApplicationIntegrationIdPost(ctx, optional)
20_998

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_998_101 | ApplicationIntegrationNotFound | Application integration not found. |  | 20_998_102 | UnknownExceptionWhenCheckApplicationExistence | Unknown exception when checking application existence. |  | 20_998_103 | UnknownExceptionWhenCreateTenant | Unknown exception when createing tenant. |  | 20_998_104 | UnknownExceptionWhenSendingEmailToInvitedEmailAddress | Unknown exception when sending email to invited email address. |  | 20_998_105 | UserNotFound | User not found. |  | 20_998_106 | UnknownExceptionWhenInsertApplication | Unknown exception when inserting application. |  | 20_998_107 | UnknownExceptionWhenCreateUserTenant | Unknown exception when creating user tenant. |  | 20_998_108 | UnknownExceptionWhenCreatingEmailContent | Unknown exception when creating email content. |  | 20_998_109 | UnknownExceptionWhenCreatingEmailAcceptInvitationUrl |  |  | 20_998_110 | UnknownExceptionWhenAddingTenantInKeycloak |  |  | 20_998_111 | UnknownExceptionWhenAddingTenantInOrganization |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenantApiTenantInviteTenantByApplicationIntegrationIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantApiTenantInviteTenantByApplicationIntegrationIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TenantInvitetenantbyapplicationintegrationidBody**](TenantInvitetenantbyapplicationintegrationidBody.md)|  | 

### Return type

[**InviteTenantByApplicationIntegrationIdResponse**](InviteTenantByApplicationIntegrationIdResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantInviteTenantPost**
> InviteTenantResponse TenantInviteTenantPost(ctx, optional)
20_998

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_998_101 | ApplicationIntegrationNotFound | Application integration not found. |  | 20_998_102 | UnknownExceptionWhenCheckApplicationExistence | Unknown exception when checking application existence. |  | 20_998_103 | UnknownExceptionWhenCreateTenant | Unknown exception when createing tenant. |  | 20_998_104 | UnknownExceptionWhenSendingEmailToInvitedEmailAddress | Unknown exception when sending email to invited email address. |  | 20_998_105 | UserNotFound | User not found. |  | 20_998_106 | UnknownExceptionWhenInsertApplication | Unknown exception when inserting application. |  | 20_998_107 | UnknownExceptionWhenCreateUserTenant | Unknown exception when creating user tenant. |  | 20_998_108 | UnknownExceptionWhenCreatingEmailContent | Unknown exception when creating email content. |  | 20_998_109 | UnknownExceptionWhenCreatingEmailAcceptInvitationUrl |  |  | 20_998_110 | UnknownExceptionWhenAddingTenantInKeycloak |  |  | 20_998_111 | UnknownExceptionWhenAddingTenantInOrganization |  |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenantApiTenantInviteTenantPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantApiTenantInviteTenantPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TenantInvitetenantBody**](TenantInvitetenantBody.md)|  | 

### Return type

[**InviteTenantResponse**](InviteTenantResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantUpdateTenantConnectionPost**
> UpdateTenantConnectionResponse TenantUpdateTenantConnectionPost(ctx, optional)
20_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_997_101 | TenantNotFound | Tenant not found. |  | 20_997_102 | UnknownExceptionWhenCheckTenantExistence | Unknown exception when checking tenant existence. |  | 20_997_103 | UnknownExceptionWhenUpdateTenantConnection | Unknown exception when updating tenant connection. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenantApiTenantUpdateTenantConnectionPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantApiTenantUpdateTenantConnectionPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TenantUpdatetenantconnectionBody**](TenantUpdatetenantconnectionBody.md)|  | 

### Return type

[**UpdateTenantConnectionResponse**](UpdateTenantConnectionResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantUserTenantListGet**
> UserTenantListResponse TenantUserTenantListGet(ctx, optional)
20_999

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 20_999_101 | UnknownExceptionWhenTryingFetchTenants | Unknown exception when trying to fetch tenants. |  | 20_999_102 | UserNotFound | User not found. |  | 20_999_103 | UnknownExceptionWhenInsertApplication | Unknown exception when inserting application. |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenantApiTenantUserTenantListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantApiTenantUserTenantListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationStatus** | [**optional.Interface of InvitationStatus1**](.md)|   0 &#x3D; Invited  1 &#x3D; Accepted  2 &#x3D; Declined | 
 **page** | **optional.Int32**|  | 
 **perPage** | **optional.Int32**|  | 

### Return type

[**UserTenantListResponse**](UserTenantListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

