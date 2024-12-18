## Overview

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiProxyApi* | [**ApiProxyEndpointDelete**](docs/ApiProxyApi.md#apiproxyendpointdelete) | **Delete** /api-proxy/{endpoint} | 21_997
*ApiProxyApi* | [**ApiProxyEndpointGet**](docs/ApiProxyApi.md#apiproxyendpointget) | **Get** /api-proxy/{endpoint} | 21_997
*ApiProxyApi* | [**ApiProxyEndpointPatch**](docs/ApiProxyApi.md#apiproxyendpointpatch) | **Patch** /api-proxy/{endpoint} | 21_997
*ApiProxyApi* | [**ApiProxyEndpointPost**](docs/ApiProxyApi.md#apiproxyendpointpost) | **Post** /api-proxy/{endpoint} | 21_997
*ApiProxyApi* | [**ApiProxyEndpointPut**](docs/ApiProxyApi.md#apiproxyendpointput) | **Put** /api-proxy/{endpoint} | 21_997
*ApiProxyApi* | [**ApiProxyOauth2ChallengeGet**](docs/ApiProxyApi.md#apiproxyoauth2challengeget) | **Get** /api-proxy/oauth2/challenge |
*ApiProxyApi* | [**ApiProxyUnifiedContactsExternalIdGet**](docs/ApiProxyApi.md#apiproxyunifiedcontactsexternalidget) | **Get** /api-proxy/unified/contacts/{externalId} | 
*ApiProxyApi* | [**ApiProxyUnifiedContactsGet**](docs/ApiProxyApi.md#apiproxyunifiedcontactsget) | **Get** /api-proxy/unified/contacts | 
*ApplicationApi* | [**ApplicationCreatePost**](docs/ApplicationApi.md#applicationcreatepost) | **Post** /application/create | 15_999
*ApplicationApi* | [**ApplicationDeleteDelete**](docs/ApplicationApi.md#applicationdeletedelete) | **Delete** /application/delete | 15_997
*ApplicationApi* | [**ApplicationListGet**](docs/ApplicationApi.md#applicationlistget) | **Get** /application/list | 15_996
*ApplicationApi* | [**ApplicationUpdatePut**](docs/ApplicationApi.md#applicationupdateput) | **Put** /application/update | 15_998
*ApplicationApiKeyApi* | [**ApplicationApikeyGeneratePost**](docs/ApplicationApiKeyApi.md#applicationapikeygeneratepost) | **Post** /application-apikey/generate | 16_999
*ApplicationApiKeyApi* | [**ApplicationApikeyGetByApplicationIdGet**](docs/ApplicationApiKeyApi.md#applicationapikeygetbyapplicationidget) | **Get** /application-apikey/get-by-application-id | 16_995
*ApplicationApiKeyApi* | [**ApplicationApikeyRevokeDelete**](docs/ApplicationApiKeyApi.md#applicationapikeyrevokedelete) | **Delete** /application-apikey/revoke | 16_998
*ApplicationIntegrationApi* | [**ApplicationIntegrationAddIntegrationToApplicationPost**](docs/ApplicationIntegrationApi.md#applicationintegrationaddintegrationtoapplicationpost) | **Post** /application-integration/add-integration-to-application | 18_999
*ApplicationIntegrationApi* | [**ApplicationIntegrationApplicationIntegrationListByApplicationIdGet**](docs/ApplicationIntegrationApi.md#applicationintegrationapplicationintegrationlistbyapplicationidget) | **Get** /application-integration/application-integration-list-by-application-id | 18_997
*AuthApi* | [**AuthAuthenticatePost**](docs/AuthApi.md#authauthenticatepost) | **Post** /auth/authenticate | 23_999
*AuthApi* | [**AuthRefreshPost**](docs/AuthApi.md#authrefreshpost) | **Post** /auth/refresh | 23_998
*CategoryApi* | [**CategoryListGet**](docs/CategoryApi.md#categorylistget) | **Get** /category/list | 19_999
*IntegrationApi* | [**IntegrationGetGet**](docs/IntegrationApi.md#integrationgetget) | **Get** /integration/get | 17_997
*IntegrationApi* | [**IntegrationListGet**](docs/IntegrationApi.md#integrationlistget) | **Get** /integration/list | 17_999
*StatisticsApi* | [**StatisticsUserProfileGet**](docs/StatisticsApi.md#statisticsuserprofileget) | **Get** /statistics/user-profile | 14_997
*TenantApi* | [**TenantDeclineTenantIdGet**](docs/TenantApi.md#tenantdeclinetenantidget) | **Get** /tenant/decline-tenant/{id} | 20_996
*TenantApi* | [**TenantInviteTenantByApplicationIntegrationIdPost**](docs/TenantApi.md#tenantinvitetenantbyapplicationintegrationidpost) | **Post** /tenant/invite-tenant-by-application-integration-id | 20_998
*TenantApi* | [**TenantInviteTenantPost**](docs/TenantApi.md#tenantinvitetenantpost) | **Post** /tenant/invite-tenant | 20_998
*TenantApi* | [**TenantUpdateTenantConnectionPost**](docs/TenantApi.md#tenantupdatetenantconnectionpost) | **Post** /tenant/update-tenant-connection | 20_997
*TenantApi* | [**TenantUserTenantListGet**](docs/TenantApi.md#tenantusertenantlistget) | **Get** /tenant/user-tenant-list | 20_999
*UserApi* | [**UserCreateOrGetUserProfilePost**](docs/UserApi.md#usercreateorgetuserprofilepost) | **Post** /user/create-or-get-user-profile | 14_998

## Documentation For Models

 - [AddIntegrationToApplicationError](docs/AddIntegrationToApplicationError.md)
 - [AddIntegrationToApplicationRequest](docs/AddIntegrationToApplicationRequest.md)
 - [AddIntegrationToApplicationRequestDto](docs/AddIntegrationToApplicationRequestDto.md)
 - [AddIntegrationToApplicationRequestErrorCodes](docs/AddIntegrationToApplicationRequestErrorCodes.md)
 - [AddIntegrationToApplicationResponse](docs/AddIntegrationToApplicationResponse.md)
 - [AllOfAddIntegrationToApplicationErrorCode](docs/AllOfAddIntegrationToApplicationErrorCode.md)
 - [AllOfAddIntegrationToApplicationErrorType_](docs/AllOfAddIntegrationToApplicationErrorType_.md)
 - [AllOfAddIntegrationToApplicationRequestDto](docs/AllOfAddIntegrationToApplicationRequestDto.md)
 - [AllOfAddIntegrationToApplicationResponseError_](docs/AllOfAddIntegrationToApplicationResponseError_.md)
 - [AllOfApplicationApiKeyByApplicationIdErrorCode](docs/AllOfApplicationApiKeyByApplicationIdErrorCode.md)
 - [AllOfApplicationApiKeyByApplicationIdErrorType_](docs/AllOfApplicationApiKeyByApplicationIdErrorType_.md)
 - [AllOfApplicationApiKeyByApplicationIdResponseData](docs/AllOfApplicationApiKeyByApplicationIdResponseData.md)
 - [AllOfApplicationApiKeyByApplicationIdResponseError_](docs/AllOfApplicationApiKeyByApplicationIdResponseError_.md)
 - [AllOfApplicationIntegrationDtoApplication](docs/AllOfApplicationIntegrationDtoApplication.md)
 - [AllOfApplicationIntegrationDtoIntegration](docs/AllOfApplicationIntegrationDtoIntegration.md)
 - [AllOfApplicationIntegrationListByApplicationIdErrorCode](docs/AllOfApplicationIntegrationListByApplicationIdErrorCode.md)
 - [AllOfApplicationIntegrationListByApplicationIdErrorType_](docs/AllOfApplicationIntegrationListByApplicationIdErrorType_.md)
 - [AllOfApplicationIntegrationListByApplicationIdResponseData](docs/AllOfApplicationIntegrationListByApplicationIdResponseData.md)
 - [AllOfApplicationIntegrationListByApplicationIdResponseError_](docs/AllOfApplicationIntegrationListByApplicationIdResponseError_.md)
 - [AllOfCallActionErrorCode](docs/AllOfCallActionErrorCode.md)
 - [AllOfCallActionErrorType_](docs/AllOfCallActionErrorType_.md)
 - [AllOfCallActionResponseData](docs/AllOfCallActionResponseData.md)
 - [AllOfCallActionResponseError_](docs/AllOfCallActionResponseError_.md)
 - [AllOfCategoryListErrorCode](docs/AllOfCategoryListErrorCode.md)
 - [AllOfCategoryListErrorType_](docs/AllOfCategoryListErrorType_.md)
 - [AllOfCategoryListResponseData](docs/AllOfCategoryListResponseData.md)
 - [AllOfCategoryListResponseError_](docs/AllOfCategoryListResponseError_.md)
 - [AllOfCreateApplicationApiKeyErrorCode](docs/AllOfCreateApplicationApiKeyErrorCode.md)
 - [AllOfCreateApplicationApiKeyErrorType_](docs/AllOfCreateApplicationApiKeyErrorType_.md)
 - [AllOfCreateApplicationApiKeyRequestDto](docs/AllOfCreateApplicationApiKeyRequestDto.md)
 - [AllOfCreateApplicationApiKeyResponseData](docs/AllOfCreateApplicationApiKeyResponseData.md)
 - [AllOfCreateApplicationApiKeyResponseError_](docs/AllOfCreateApplicationApiKeyResponseError_.md)
 - [AllOfCreateApplicationErrorCode](docs/AllOfCreateApplicationErrorCode.md)
 - [AllOfCreateApplicationErrorType_](docs/AllOfCreateApplicationErrorType_.md)
 - [AllOfCreateApplicationRequestDto](docs/AllOfCreateApplicationRequestDto.md)
 - [AllOfCreateApplicationResponseData](docs/AllOfCreateApplicationResponseData.md)
 - [AllOfCreateApplicationResponseError_](docs/AllOfCreateApplicationResponseError_.md)
 - [AllOfCreateUserProfileErrorCode](docs/AllOfCreateUserProfileErrorCode.md)
 - [AllOfCreateUserProfileErrorType_](docs/AllOfCreateUserProfileErrorType_.md)
 - [AllOfCreateUserProfileResponseData](docs/AllOfCreateUserProfileResponseData.md)
 - [AllOfCreateUserProfileResponseError_](docs/AllOfCreateUserProfileResponseError_.md)
 - [AllOfDeclineTenantInvitationErrorCode](docs/AllOfDeclineTenantInvitationErrorCode.md)
 - [AllOfDeclineTenantInvitationErrorType_](docs/AllOfDeclineTenantInvitationErrorType_.md)
 - [AllOfDeclineTenantInvitationResponseError_](docs/AllOfDeclineTenantInvitationResponseError_.md)
 - [AllOfDeleteApplicationErrorCode](docs/AllOfDeleteApplicationErrorCode.md)
 - [AllOfDeleteApplicationErrorType_](docs/AllOfDeleteApplicationErrorType_.md)
 - [AllOfDeleteApplicationResponseError_](docs/AllOfDeleteApplicationResponseError_.md)
 - [AllOfErrorType_](docs/AllOfErrorType_.md)
 - [AllOfGetAuthenticationTokenErrorCode](docs/AllOfGetAuthenticationTokenErrorCode.md)
 - [AllOfGetAuthenticationTokenErrorType_](docs/AllOfGetAuthenticationTokenErrorType_.md)
 - [AllOfGetAuthenticationTokenResponseData](docs/AllOfGetAuthenticationTokenResponseData.md)
 - [AllOfGetAuthenticationTokenResponseError_](docs/AllOfGetAuthenticationTokenResponseError_.md)
 - [AllOfGetOneIntegrationErrorCode](docs/AllOfGetOneIntegrationErrorCode.md)
 - [AllOfGetOneIntegrationErrorType_](docs/AllOfGetOneIntegrationErrorType_.md)
 - [AllOfGetOneIntegrationResponseData](docs/AllOfGetOneIntegrationResponseData.md)
 - [AllOfGetOneIntegrationResponseError_](docs/AllOfGetOneIntegrationResponseError_.md)
 - [AllOfIntegrationDtoReleaseStatus](docs/AllOfIntegrationDtoReleaseStatus.md)
 - [AllOfIntegrationListErrorCode](docs/AllOfIntegrationListErrorCode.md)
 - [AllOfIntegrationListErrorType_](docs/AllOfIntegrationListErrorType_.md)
 - [AllOfIntegrationListResponseData](docs/AllOfIntegrationListResponseData.md)
 - [AllOfIntegrationListResponseError_](docs/AllOfIntegrationListResponseError_.md)
 - [AllOfInviteTenantByApplicationIntegrationIdErrorCode](docs/AllOfInviteTenantByApplicationIntegrationIdErrorCode.md)
 - [AllOfInviteTenantByApplicationIntegrationIdErrorType_](docs/AllOfInviteTenantByApplicationIntegrationIdErrorType_.md)
 - [AllOfInviteTenantByApplicationIntegrationIdRequestDto](docs/AllOfInviteTenantByApplicationIntegrationIdRequestDto.md)
 - [AllOfInviteTenantByApplicationIntegrationIdResponseData](docs/AllOfInviteTenantByApplicationIntegrationIdResponseData.md)
 - [AllOfInviteTenantByApplicationIntegrationIdResponseError_](docs/AllOfInviteTenantByApplicationIntegrationIdResponseError_.md)
 - [AllOfInviteTenantErrorCode](docs/AllOfInviteTenantErrorCode.md)
 - [AllOfInviteTenantErrorType_](docs/AllOfInviteTenantErrorType_.md)
 - [AllOfInviteTenantRequestDto](docs/AllOfInviteTenantRequestDto.md)
 - [AllOfInviteTenantResponseData](docs/AllOfInviteTenantResponseData.md)
 - [AllOfInviteTenantResponseError_](docs/AllOfInviteTenantResponseError_.md)
 - [AllOfOAuth2ChallengeDtoResponseData](docs/AllOfOAuth2ChallengeDtoResponseData.md)
 - [AllOfOAuth2ChallengeDtoResponseError_](docs/AllOfOAuth2ChallengeDtoResponseError_.md)
 - [AllOfOAuth2ChallengeDtoTokens](docs/AllOfOAuth2ChallengeDtoTokens.md)
 - [AllOfRefreshAuthenticationTokenErrorCode](docs/AllOfRefreshAuthenticationTokenErrorCode.md)
 - [AllOfRefreshAuthenticationTokenErrorType_](docs/AllOfRefreshAuthenticationTokenErrorType_.md)
 - [AllOfRefreshAuthenticationTokenResponseData](docs/AllOfRefreshAuthenticationTokenResponseData.md)
 - [AllOfRefreshAuthenticationTokenResponseError_](docs/AllOfRefreshAuthenticationTokenResponseError_.md)
 - [AllOfResponseError_](docs/AllOfResponseError_.md)
 - [AllOfRevokeApplicationApiKeyErrorCode](docs/AllOfRevokeApplicationApiKeyErrorCode.md)
 - [AllOfRevokeApplicationApiKeyErrorType_](docs/AllOfRevokeApplicationApiKeyErrorType_.md)
 - [AllOfRevokeApplicationApiKeyResponseError_](docs/AllOfRevokeApplicationApiKeyResponseError_.md)
 - [AllOfTenantDtoApplicationIntegration](docs/AllOfTenantDtoApplicationIntegration.md)
 - [AllOfTenantDtoInvitationStatus](docs/AllOfTenantDtoInvitationStatus.md)
 - [AllOfUnifiedContactResponseData](docs/AllOfUnifiedContactResponseData.md)
 - [AllOfUnifiedContactResponseError_](docs/AllOfUnifiedContactResponseError_.md)
 - [AllOfUnifiedContactsListResponseData](docs/AllOfUnifiedContactsListResponseData.md)
 - [AllOfUnifiedContactsListResponseError_](docs/AllOfUnifiedContactsListResponseError_.md)
 - [AllOfUpdateApplicationErrorCode](docs/AllOfUpdateApplicationErrorCode.md)
 - [AllOfUpdateApplicationErrorType_](docs/AllOfUpdateApplicationErrorType_.md)
 - [AllOfUpdateApplicationRequestDto](docs/AllOfUpdateApplicationRequestDto.md)
 - [AllOfUpdateApplicationResponseData](docs/AllOfUpdateApplicationResponseData.md)
 - [AllOfUpdateApplicationResponseError_](docs/AllOfUpdateApplicationResponseError_.md)
 - [AllOfUpdateTenantConnectionErrorCode](docs/AllOfUpdateTenantConnectionErrorCode.md)
 - [AllOfUpdateTenantConnectionErrorType_](docs/AllOfUpdateTenantConnectionErrorType_.md)
 - [AllOfUpdateTenantConnectionRequestDto](docs/AllOfUpdateTenantConnectionRequestDto.md)
 - [AllOfUpdateTenantConnectionResponseError_](docs/AllOfUpdateTenantConnectionResponseError_.md)
 - [AllOfUserApplicationListErrorCode](docs/AllOfUserApplicationListErrorCode.md)
 - [AllOfUserApplicationListErrorType_](docs/AllOfUserApplicationListErrorType_.md)
 - [AllOfUserApplicationListResponseData](docs/AllOfUserApplicationListResponseData.md)
 - [AllOfUserApplicationListResponseError_](docs/AllOfUserApplicationListResponseError_.md)
 - [AllOfUserProfileStatisticsErrorCode](docs/AllOfUserProfileStatisticsErrorCode.md)
 - [AllOfUserProfileStatisticsErrorType_](docs/AllOfUserProfileStatisticsErrorType_.md)
 - [AllOfUserProfileStatisticsResponseData](docs/AllOfUserProfileStatisticsResponseData.md)
 - [AllOfUserProfileStatisticsResponseError_](docs/AllOfUserProfileStatisticsResponseError_.md)
 - [AllOfUserTenantListErrorCode](docs/AllOfUserTenantListErrorCode.md)
 - [AllOfUserTenantListErrorType_](docs/AllOfUserTenantListErrorType_.md)
 - [AllOfUserTenantListResponseData](docs/AllOfUserTenantListResponseData.md)
 - [AllOfUserTenantListResponseError_](docs/AllOfUserTenantListResponseError_.md)
 - [ApiproxyEndpointBody](docs/ApiproxyEndpointBody.md)
 - [ApiproxyEndpointBody1](docs/ApiproxyEndpointBody1.md)
 - [ApiproxyEndpointBody10](docs/ApiproxyEndpointBody10.md)
 - [ApiproxyEndpointBody11](docs/ApiproxyEndpointBody11.md)
 - [ApiproxyEndpointBody2](docs/ApiproxyEndpointBody2.md)
 - [ApiproxyEndpointBody3](docs/ApiproxyEndpointBody3.md)
 - [ApiproxyEndpointBody4](docs/ApiproxyEndpointBody4.md)
 - [ApiproxyEndpointBody5](docs/ApiproxyEndpointBody5.md)
 - [ApiproxyEndpointBody6](docs/ApiproxyEndpointBody6.md)
 - [ApiproxyEndpointBody7](docs/ApiproxyEndpointBody7.md)
 - [ApiproxyEndpointBody8](docs/ApiproxyEndpointBody8.md)
 - [ApiproxyEndpointBody9](docs/ApiproxyEndpointBody9.md)
 - [ApplicationApiKeyByApplicationIdError](docs/ApplicationApiKeyByApplicationIdError.md)
 - [ApplicationApiKeyByApplicationIdErrorCodes](docs/ApplicationApiKeyByApplicationIdErrorCodes.md)
 - [ApplicationApiKeyByApplicationIdResponse](docs/ApplicationApiKeyByApplicationIdResponse.md)
 - [ApplicationApiKeyDto](docs/ApplicationApiKeyDto.md)
 - [ApplicationCreateBody](docs/ApplicationCreateBody.md)
 - [ApplicationCreateBody1](docs/ApplicationCreateBody1.md)
 - [ApplicationCreateBody2](docs/ApplicationCreateBody2.md)
 - [ApplicationDto](docs/ApplicationDto.md)
 - [ApplicationIntegrationDto](docs/ApplicationIntegrationDto.md)
 - [ApplicationIntegrationListByApplicationIdError](docs/ApplicationIntegrationListByApplicationIdError.md)
 - [ApplicationIntegrationListByApplicationIdRequestErrorCodes](docs/ApplicationIntegrationListByApplicationIdRequestErrorCodes.md)
 - [ApplicationIntegrationListByApplicationIdResponse](docs/ApplicationIntegrationListByApplicationIdResponse.md)
 - [ApplicationIntegrationListDto](docs/ApplicationIntegrationListDto.md)
 - [ApplicationListDto](docs/ApplicationListDto.md)
 - [ApplicationUpdateBody](docs/ApplicationUpdateBody.md)
 - [ApplicationUpdateBody1](docs/ApplicationUpdateBody1.md)
 - [ApplicationUpdateBody2](docs/ApplicationUpdateBody2.md)
 - [ApplicationapikeyGenerateBody](docs/ApplicationapikeyGenerateBody.md)
 - [ApplicationapikeyGenerateBody1](docs/ApplicationapikeyGenerateBody1.md)
 - [ApplicationapikeyGenerateBody2](docs/ApplicationapikeyGenerateBody2.md)
 - [ApplicationintegrationAddintegrationtoapplicationBody](docs/ApplicationintegrationAddintegrationtoapplicationBody.md)
 - [ApplicationintegrationAddintegrationtoapplicationBody1](docs/ApplicationintegrationAddintegrationtoapplicationBody1.md)
 - [ApplicationintegrationAddintegrationtoapplicationBody2](docs/ApplicationintegrationAddintegrationtoapplicationBody2.md)
 - [CallActionBodyDto](docs/CallActionBodyDto.md)
 - [CallActionDto](docs/CallActionDto.md)
 - [CallActionError](docs/CallActionError.md)
 - [CallActionErrorCodes](docs/CallActionErrorCodes.md)
 - [CallActionResponse](docs/CallActionResponse.md)
 - [CategoryDto](docs/CategoryDto.md)
 - [CategoryListDto](docs/CategoryListDto.md)
 - [CategoryListError](docs/CategoryListError.md)
 - [CategoryListRequestErrorCodes](docs/CategoryListRequestErrorCodes.md)
 - [CategoryListResponse](docs/CategoryListResponse.md)
 - [ClientErrorType](docs/ClientErrorType.md)
 - [CreateApplicationApiKeyError](docs/CreateApplicationApiKeyError.md)
 - [CreateApplicationApiKeyRequest](docs/CreateApplicationApiKeyRequest.md)
 - [CreateApplicationApiKeyRequestDto](docs/CreateApplicationApiKeyRequestDto.md)
 - [CreateApplicationApiKeyRequestErrorCodes](docs/CreateApplicationApiKeyRequestErrorCodes.md)
 - [CreateApplicationApiKeyResponse](docs/CreateApplicationApiKeyResponse.md)
 - [CreateApplicationError](docs/CreateApplicationError.md)
 - [CreateApplicationRequest](docs/CreateApplicationRequest.md)
 - [CreateApplicationRequestDto](docs/CreateApplicationRequestDto.md)
 - [CreateApplicationRequestErrorCodes](docs/CreateApplicationRequestErrorCodes.md)
 - [CreateApplicationResponse](docs/CreateApplicationResponse.md)
 - [CreateUserProfileError](docs/CreateUserProfileError.md)
 - [CreateUserProfileRequestErrorCodes](docs/CreateUserProfileRequestErrorCodes.md)
 - [CreateUserProfileResponse](docs/CreateUserProfileResponse.md)
 - [DeclineTenantInvitationError](docs/DeclineTenantInvitationError.md)
 - [DeclineTenantInvitationRequestErrorCodes](docs/DeclineTenantInvitationRequestErrorCodes.md)
 - [DeclineTenantInvitationResponse](docs/DeclineTenantInvitationResponse.md)
 - [DeleteApplicationError](docs/DeleteApplicationError.md)
 - [DeleteApplicationRequestErrorCodes](docs/DeleteApplicationRequestErrorCodes.md)
 - [DeleteApplicationResponse](docs/DeleteApplicationResponse.md)
 - [GetAuthenticationTokenDto](docs/GetAuthenticationTokenDto.md)
 - [GetAuthenticationTokenError](docs/GetAuthenticationTokenError.md)
 - [GetAuthenticationTokenRequestErrorCodes](docs/GetAuthenticationTokenRequestErrorCodes.md)
 - [GetAuthenticationTokenResponse](docs/GetAuthenticationTokenResponse.md)
 - [GetOneIntegrationError](docs/GetOneIntegrationError.md)
 - [GetOneIntegrationRequestErrorCodes](docs/GetOneIntegrationRequestErrorCodes.md)
 - [GetOneIntegrationResponse](docs/GetOneIntegrationResponse.md)
 - [IntegrationDto](docs/IntegrationDto.md)
 - [IntegrationEndpointDto](docs/IntegrationEndpointDto.md)
 - [IntegrationEventDto](docs/IntegrationEventDto.md)
 - [IntegrationFieldDto](docs/IntegrationFieldDto.md)
 - [IntegrationListDto](docs/IntegrationListDto.md)
 - [IntegrationListError](docs/IntegrationListError.md)
 - [IntegrationListRequestErrorCodes](docs/IntegrationListRequestErrorCodes.md)
 - [IntegrationListResponse](docs/IntegrationListResponse.md)
 - [InvitationStatus](docs/InvitationStatus.md)
 - [InvitationStatus1](docs/InvitationStatus1.md)
 - [InviteTenantByApplicationIntegrationIdError](docs/InviteTenantByApplicationIntegrationIdError.md)
 - [InviteTenantByApplicationIntegrationIdRequest](docs/InviteTenantByApplicationIntegrationIdRequest.md)
 - [InviteTenantByApplicationIntegrationIdRequestDto](docs/InviteTenantByApplicationIntegrationIdRequestDto.md)
 - [InviteTenantByApplicationIntegrationIdResponse](docs/InviteTenantByApplicationIntegrationIdResponse.md)
 - [InviteTenantError](docs/InviteTenantError.md)
 - [InviteTenantRequest](docs/InviteTenantRequest.md)
 - [InviteTenantRequestDto](docs/InviteTenantRequestDto.md)
 - [InviteTenantRequestErrorCodes](docs/InviteTenantRequestErrorCodes.md)
 - [InviteTenantResponse](docs/InviteTenantResponse.md)
 - [ModelError](docs/ModelError.md)
 - [OAuth2ChallengeDto](docs/OAuth2ChallengeDto.md)
 - [OAuth2ChallengeDtoResponse](docs/OAuth2ChallengeDtoResponse.md)
 - [OAuth2ChallengeTokenDto](docs/OAuth2ChallengeTokenDto.md)
 - [RefreshAuthenticationTokenDto](docs/RefreshAuthenticationTokenDto.md)
 - [RefreshAuthenticationTokenError](docs/RefreshAuthenticationTokenError.md)
 - [RefreshAuthenticationTokenRequestErrorCodes](docs/RefreshAuthenticationTokenRequestErrorCodes.md)
 - [RefreshAuthenticationTokenResponse](docs/RefreshAuthenticationTokenResponse.md)
 - [ReleaseStatus](docs/ReleaseStatus.md)
 - [Response](docs/Response.md)
 - [RevokeApplicationApiKeyError](docs/RevokeApplicationApiKeyError.md)
 - [RevokeApplicationApiKeyRequestErrorCodes](docs/RevokeApplicationApiKeyRequestErrorCodes.md)
 - [RevokeApplicationApiKeyResponse](docs/RevokeApplicationApiKeyResponse.md)
 - [TenantDto](docs/TenantDto.md)
 - [TenantInvitetenantBody](docs/TenantInvitetenantBody.md)
 - [TenantInvitetenantBody1](docs/TenantInvitetenantBody1.md)
 - [TenantInvitetenantBody2](docs/TenantInvitetenantBody2.md)
 - [TenantInvitetenantbyapplicationintegrationidBody](docs/TenantInvitetenantbyapplicationintegrationidBody.md)
 - [TenantInvitetenantbyapplicationintegrationidBody1](docs/TenantInvitetenantbyapplicationintegrationidBody1.md)
 - [TenantInvitetenantbyapplicationintegrationidBody2](docs/TenantInvitetenantbyapplicationintegrationidBody2.md)
 - [TenantListDto](docs/TenantListDto.md)
 - [TenantUpdatetenantconnectionBody](docs/TenantUpdatetenantconnectionBody.md)
 - [TenantUpdatetenantconnectionBody1](docs/TenantUpdatetenantconnectionBody1.md)
 - [TenantUpdatetenantconnectionBody2](docs/TenantUpdatetenantconnectionBody2.md)
 - [UnifiedContact](docs/UnifiedContact.md)
 - [UnifiedContactResponse](docs/UnifiedContactResponse.md)
 - [UnifiedContactsList](docs/UnifiedContactsList.md)
 - [UnifiedContactsListResponse](docs/UnifiedContactsListResponse.md)
 - [UpdateApplicationError](docs/UpdateApplicationError.md)
 - [UpdateApplicationRequest](docs/UpdateApplicationRequest.md)
 - [UpdateApplicationRequestDto](docs/UpdateApplicationRequestDto.md)
 - [UpdateApplicationRequestErrorCodes](docs/UpdateApplicationRequestErrorCodes.md)
 - [UpdateApplicationResponse](docs/UpdateApplicationResponse.md)
 - [UpdateTenantConnectionError](docs/UpdateTenantConnectionError.md)
 - [UpdateTenantConnectionRequest](docs/UpdateTenantConnectionRequest.md)
 - [UpdateTenantConnectionRequestDto](docs/UpdateTenantConnectionRequestDto.md)
 - [UpdateTenantConnectionRequestErrorCodes](docs/UpdateTenantConnectionRequestErrorCodes.md)
 - [UpdateTenantConnectionResponse](docs/UpdateTenantConnectionResponse.md)
 - [UserApplicationListError](docs/UserApplicationListError.md)
 - [UserApplicationListRequestErrorCodes](docs/UserApplicationListRequestErrorCodes.md)
 - [UserApplicationListResponse](docs/UserApplicationListResponse.md)
 - [UserDto](docs/UserDto.md)
 - [UserProfileStatisticsDto](docs/UserProfileStatisticsDto.md)
 - [UserProfileStatisticsError](docs/UserProfileStatisticsError.md)
 - [UserProfileStatisticsErrorCodes](docs/UserProfileStatisticsErrorCodes.md)
 - [UserProfileStatisticsResponse](docs/UserProfileStatisticsResponse.md)
 - [UserTenantListError](docs/UserTenantListError.md)
 - [UserTenantListRequestErrorCodes](docs/UserTenantListRequestErrorCodes.md)
 - [UserTenantListResponse](docs/UserTenantListResponse.md)

## Documentation For Authorization

## Bearer

## Author


