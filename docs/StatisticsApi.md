# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatisticsUserProfileGet**](StatisticsApi.md#StatisticsUserProfileGet) | **Get** /statistics/user-profile | 14_997

# **StatisticsUserProfileGet**
> UserProfileStatisticsResponse StatisticsUserProfileGet(ctx, )
14_997

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 14_997_101 | UserNotFound | User not found. |  | 14_997_102 | UserDoesNotHaveActiveSubscription | User does not have an active subscription. |  | 14_997_103 | UnknownExceptionWhenFetchUserProfileStatistics | Unknown exception when fetching user profile statistics. |

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserProfileStatisticsResponse**](UserProfileStatisticsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

