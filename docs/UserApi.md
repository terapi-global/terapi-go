# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCreateOrGetUserProfilePost**](UserApi.md#UserCreateOrGetUserProfilePost) | **Post** /user/create-or-get-user-profile | 14_998

# **UserCreateOrGetUserProfilePost**
> CreateUserProfileResponse UserCreateOrGetUserProfilePost(ctx, )
14_998

### Error Codes  | Code | Name | Message |  | :--- | :--- | :--- |  | 14_998_101 | UnknownExceptionWhenFetchJwtUserProfile | Unknown exception when fetching JWT user profile. |  | 14_998_102 | UnknownExceptionWhenCreateUserProfile | Unknown exception when creating user profile. |  | 14_998_103 | IncorrectDataInUserJwtPayload | Incorrect data in user JWT payload. |  | 14_998_104 | UnknownExceptionWhenSubscribeUserToFreePlan | Unknown exception when subscribing user to free plan. |  | 14_998_105 | UserAlreadyExistInDataStorage | Logged user not found in data storage. |  | 14_998_106 | UnknownExceptionWhenCheckUserExistence | Unknown exception when fetching logged user profile. |

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CreateUserProfileResponse**](CreateUserProfileResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

