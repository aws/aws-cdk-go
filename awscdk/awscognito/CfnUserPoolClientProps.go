package awscognito


// Properties for defining a `CfnUserPoolClient`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolClientProps := &CfnUserPoolClientProps{
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	AccessTokenValidity: jsii.Number(123),
//   	AllowedOAuthFlows: []*string{
//   		jsii.String("allowedOAuthFlows"),
//   	},
//   	AllowedOAuthFlowsUserPoolClient: jsii.Boolean(false),
//   	AllowedOAuthScopes: []*string{
//   		jsii.String("allowedOAuthScopes"),
//   	},
//   	AnalyticsConfiguration: &AnalyticsConfigurationProperty{
//   		ApplicationArn: jsii.String("applicationArn"),
//   		ApplicationId: jsii.String("applicationId"),
//   		ExternalId: jsii.String("externalId"),
//   		RoleArn: jsii.String("roleArn"),
//   		UserDataShared: jsii.Boolean(false),
//   	},
//   	AuthSessionValidity: jsii.Number(123),
//   	CallbackUrLs: []*string{
//   		jsii.String("callbackUrLs"),
//   	},
//   	ClientName: jsii.String("clientName"),
//   	DefaultRedirectUri: jsii.String("defaultRedirectUri"),
//   	EnablePropagateAdditionalUserContextData: jsii.Boolean(false),
//   	EnableTokenRevocation: jsii.Boolean(false),
//   	ExplicitAuthFlows: []*string{
//   		jsii.String("explicitAuthFlows"),
//   	},
//   	GenerateSecret: jsii.Boolean(false),
//   	IdTokenValidity: jsii.Number(123),
//   	LogoutUrLs: []*string{
//   		jsii.String("logoutUrLs"),
//   	},
//   	PreventUserExistenceErrors: jsii.String("preventUserExistenceErrors"),
//   	ReadAttributes: []*string{
//   		jsii.String("readAttributes"),
//   	},
//   	RefreshTokenValidity: jsii.Number(123),
//   	SupportedIdentityProviders: []*string{
//   		jsii.String("supportedIdentityProviders"),
//   	},
//   	TokenValidityUnits: &TokenValidityUnitsProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		IdToken: jsii.String("idToken"),
//   		RefreshToken: jsii.String("refreshToken"),
//   	},
//   	WriteAttributes: []*string{
//   		jsii.String("writeAttributes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html
//
type CfnUserPoolClientProps struct {
	// The ID of the user pool where you want to create an app client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The access token time limit.
	//
	// After this limit expires, your user can't use their access token. To specify the time unit for `AccessTokenValidity` as `seconds` , `minutes` , `hours` , or `days` , set a `TokenValidityUnits` value in your API request.
	//
	// For example, when you set `AccessTokenValidity` to `10` and `TokenValidityUnits` to `hours` , your user can authorize access with
	// their access token for 10 hours.
	//
	// The default time unit for `AccessTokenValidity` in an API request is hours. *Valid range* is displayed below in seconds.
	//
	// If you don't specify otherwise in the configuration of your app client, your access
	// tokens are valid for one hour.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-accesstokenvalidity
	//
	AccessTokenValidity *float64 `field:"optional" json:"accessTokenValidity" yaml:"accessTokenValidity"`
	// The OAuth grant types that you want your app client to generate for clients in managed login authentication.
	//
	// To create an app client that generates client credentials grants, you must add `client_credentials` as the only allowed OAuth flow.
	//
	// - **code** - Use a code grant flow, which provides an authorization code as the response. This code can be exchanged for access tokens with the `/oauth2/token` endpoint.
	// - **implicit** - Issue the access token, and the ID token when scopes like `openid` and `profile` are requested, directly to your user.
	// - **client_credentials** - Issue the access token from the `/oauth2/token` endpoint directly to a non-person user, authorized by a combination of the client ID and client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-allowedoauthflows
	//
	AllowedOAuthFlows *[]*string `field:"optional" json:"allowedOAuthFlows" yaml:"allowedOAuthFlows"`
	// Set to `true` to use OAuth 2.0 authorization server features in your app client.
	//
	// This parameter must have a value of `true` before you can configure the following features in your app client.
	//
	// - `CallBackURLs` : Callback URLs.
	// - `LogoutURLs` : Sign-out redirect URLs.
	// - `AllowedOAuthScopes` : OAuth 2.0 scopes.
	// - `AllowedOAuthFlows` : Support for authorization code, implicit, and client credentials OAuth 2.0 grants.
	//
	// To use authorization server features, configure one of these features in the Amazon Cognito console or set `AllowedOAuthFlowsUserPoolClient` to `true` in a `CreateUserPoolClient` or `UpdateUserPoolClient` API request. If you don't set a value for `AllowedOAuthFlowsUserPoolClient` in a request with the AWS CLI or SDKs, it defaults to `false` . When `false` , only SDK-based API sign-in is permitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-allowedoauthflowsuserpoolclient
	//
	AllowedOAuthFlowsUserPoolClient interface{} `field:"optional" json:"allowedOAuthFlowsUserPoolClient" yaml:"allowedOAuthFlowsUserPoolClient"`
	// The OAuth, OpenID Connect (OIDC), and custom scopes that you want to permit your app client to authorize access with.
	//
	// Scopes govern access control to user pool self-service API operations, user data from the `userInfo` endpoint, and third-party APIs. Scope values include `phone` , `email` , `openid` , and `profile` . The `aws.cognito.signin.user.admin` scope authorizes user self-service operations. Custom scopes with resource servers authorize access to external APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-allowedoauthscopes
	//
	AllowedOAuthScopes *[]*string `field:"optional" json:"allowedOAuthScopes" yaml:"allowedOAuthScopes"`
	// The user pool analytics configuration for collecting metrics and sending them to your Amazon Pinpoint campaign.
	//
	// In AWS Regions where Amazon Pinpoint isn't available, user pools might not have access to analytics or might be configurable with campaigns in the US East (N. Virginia) Region. For more information, see [Using Amazon Pinpoint analytics](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-pinpoint-integration.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-analyticsconfiguration
	//
	AnalyticsConfiguration interface{} `field:"optional" json:"analyticsConfiguration" yaml:"analyticsConfiguration"`
	// Amazon Cognito creates a session token for each API request in an authentication flow.
	//
	// `AuthSessionValidity` is the duration, in minutes, of that session token. Your user pool native user must respond to each authentication challenge before the session expires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-authsessionvalidity
	//
	AuthSessionValidity *float64 `field:"optional" json:"authSessionValidity" yaml:"authSessionValidity"`
	// A list of allowed redirect, or callback, URLs for managed login authentication.
	//
	// These URLs are the paths where you want to send your users' browsers after they complete authentication with managed login or a third-party IdP. Typically, callback URLs are the home of an application that uses OAuth or OIDC libraries to process authentication outcomes.
	//
	// A redirect URI must meet the following requirements:
	//
	// - Be an absolute URI.
	// - Be registered with the authorization server. Amazon Cognito doesn't accept authorization requests with `redirect_uri` values that aren't in the list of `CallbackURLs` that you provide in this parameter.
	// - Not include a fragment component.
	//
	// See [OAuth 2.0 - Redirection Endpoint](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6749#section-3.1.2) .
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-callbackurls
	//
	CallbackUrLs *[]*string `field:"optional" json:"callbackUrLs" yaml:"callbackUrLs"`
	// A friendly name for the app client that you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-clientname
	//
	ClientName *string `field:"optional" json:"clientName" yaml:"clientName"`
	// The default redirect URI.
	//
	// In app clients with one assigned IdP, replaces `redirect_uri` in authentication requests. Must be in the `CallbackURLs` list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-defaultredirecturi
	//
	DefaultRedirectUri *string `field:"optional" json:"defaultRedirectUri" yaml:"defaultRedirectUri"`
	// When `true` , your application can include additional `UserContextData` in authentication requests.
	//
	// This data includes the IP address, and contributes to analysis by threat protection features. For more information about propagation of user context data, see [Adding session data to API requests](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-adaptive-authentication.html#user-pool-settings-adaptive-authentication-device-fingerprint) . If you don’t include this parameter, you can't send the source IP address to Amazon Cognito threat protection features. You can only activate `EnablePropagateAdditionalUserContextData` in an app client that has a client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-enablepropagateadditionalusercontextdata
	//
	EnablePropagateAdditionalUserContextData interface{} `field:"optional" json:"enablePropagateAdditionalUserContextData" yaml:"enablePropagateAdditionalUserContextData"`
	// Activates or deactivates [token revocation](https://docs.aws.amazon.com/cognito/latest/developerguide/token-revocation.html) in the target app client.
	//
	// Revoke tokens with `API_RevokeToken` .
	//
	// If you don't include this parameter, token revocation is automatically activated for the new user pool client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-enabletokenrevocation
	//
	EnableTokenRevocation interface{} `field:"optional" json:"enableTokenRevocation" yaml:"enableTokenRevocation"`
	// The [authentication flows](https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow-methods.html) that you want your user pool client to support. For each app client in your user pool, you can sign in your users with any combination of one or more flows, including with a user name and Secure Remote Password (SRP), a user name and password, or a custom authentication process that you define with Lambda functions.
	//
	// > If you don't specify a value for `ExplicitAuthFlows` , your app client supports `ALLOW_REFRESH_TOKEN_AUTH` , `ALLOW_USER_SRP_AUTH` , and `ALLOW_CUSTOM_AUTH` .
	//
	// The values for authentication flow options include the following.
	//
	// - `ALLOW_USER_AUTH` : Enable selection-based sign-in with `USER_AUTH` . This setting covers username-password, secure remote password (SRP), passwordless, and passkey authentication. This authentiation flow can do username-password and SRP authentication without other `ExplicitAuthFlows` permitting them. For example users can complete an SRP challenge through `USER_AUTH` without the flow `USER_SRP_AUTH` being active for the app client. This flow doesn't include `CUSTOM_AUTH` .
	//
	// To activate this setting, your user pool must be in the [Essentials tier](https://docs.aws.amazon.com/cognito/latest/developerguide/feature-plans-features-essentials.html) or higher.
	// - `ALLOW_ADMIN_USER_PASSWORD_AUTH` : Enable admin based user password authentication flow `ADMIN_USER_PASSWORD_AUTH` . This setting replaces the `ADMIN_NO_SRP_AUTH` setting. With this authentication flow, your app passes a user name and password to Amazon Cognito in the request, instead of using the Secure Remote Password (SRP) protocol to securely transmit the password.
	// - `ALLOW_CUSTOM_AUTH` : Enable Lambda trigger based authentication.
	// - `ALLOW_USER_PASSWORD_AUTH` : Enable user password-based authentication. In this flow, Amazon Cognito receives the password in the request instead of using the SRP protocol to verify passwords.
	// - `ALLOW_USER_SRP_AUTH` : Enable SRP-based authentication.
	// - `ALLOW_REFRESH_TOKEN_AUTH` : Enable authflow to refresh tokens.
	//
	// In some environments, you will see the values `ADMIN_NO_SRP_AUTH` , `CUSTOM_AUTH_FLOW_ONLY` , or `USER_PASSWORD_AUTH` . You can't assign these legacy `ExplicitAuthFlows` values to user pool clients at the same time as values that begin with `ALLOW_` ,
	// like `ALLOW_USER_SRP_AUTH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-explicitauthflows
	//
	ExplicitAuthFlows *[]*string `field:"optional" json:"explicitAuthFlows" yaml:"explicitAuthFlows"`
	// When `true` , generates a client secret for the app client.
	//
	// Client secrets are used with server-side and machine-to-machine applications. Client secrets are automatically generated; you can't specify a secret value. For more information, see [App client types](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-client-apps.html#user-pool-settings-client-app-client-types) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-generatesecret
	//
	GenerateSecret interface{} `field:"optional" json:"generateSecret" yaml:"generateSecret"`
	// The ID token time limit.
	//
	// After this limit expires, your user can't use their ID token. To specify the time unit for `IdTokenValidity` as `seconds` , `minutes` , `hours` , or `days` , set a `TokenValidityUnits` value in your API request.
	//
	// For example, when you set `IdTokenValidity` as `10` and `TokenValidityUnits` as `hours` , your user can authenticate their session with their ID token for 10 hours.
	//
	// The default time unit for `IdTokenValidity` in an API request is hours. *Valid range* is displayed below in seconds.
	//
	// If you don't specify otherwise in the configuration of your app client, your ID
	// tokens are valid for one hour.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-idtokenvalidity
	//
	IdTokenValidity *float64 `field:"optional" json:"idTokenValidity" yaml:"idTokenValidity"`
	// A list of allowed logout URLs for managed login authentication.
	//
	// When you pass `logout_uri` and `client_id` parameters to `/logout` , Amazon Cognito signs out your user and redirects them to the logout URL. This parameter describes the URLs that you want to be the permitted targets of `logout_uri` . A typical use of these URLs is when a user selects "Sign out" and you redirect them to your public homepage. For more information, see [Logout endpoint](https://docs.aws.amazon.com/cognito/latest/developerguide/logout-endpoint.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-logouturls
	//
	LogoutUrLs *[]*string `field:"optional" json:"logoutUrLs" yaml:"logoutUrLs"`
	// Errors and responses that you want Amazon Cognito APIs to return during authentication, account confirmation, and password recovery when the user doesn't exist in the user pool.
	//
	// When set to `ENABLED` and the user doesn't exist, authentication returns an error indicating either the username or password was incorrect. Account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY` , those APIs return a `UserNotFoundException` exception if the user doesn't exist in the user pool.
	//
	// Valid values include:
	//
	// - `ENABLED` - This prevents user existence-related errors.
	// - `LEGACY` - This represents the early behavior of Amazon Cognito where user existence related errors aren't prevented.
	//
	// Defaults to `LEGACY` when you don't provide a value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-preventuserexistenceerrors
	//
	PreventUserExistenceErrors *string `field:"optional" json:"preventUserExistenceErrors" yaml:"preventUserExistenceErrors"`
	// The list of user attributes that you want your app client to have read access to.
	//
	// After your user authenticates in your app, their access token authorizes them to read their own attribute value for any attribute in this list.
	//
	// An example of this kind of activity is when your user selects a link to view their profile information. Your app makes a `API_GetUser` API request to retrieve and display your user's profile data.
	//
	// When you don't specify the `ReadAttributes` for your app client, your app can read the values of `email_verified` , `phone_number_verified` , and the standard attributes of your user pool. When your user pool app client has read access to these default attributes, `ReadAttributes` doesn't return any information. Amazon Cognito only populates `ReadAttributes` in the API response if you have specified your own custom set of read attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-readattributes
	//
	ReadAttributes *[]*string `field:"optional" json:"readAttributes" yaml:"readAttributes"`
	// The refresh token time limit.
	//
	// After this limit expires, your user can't use their refresh token. To specify the time unit for `RefreshTokenValidity` as `seconds` , `minutes` , `hours` , or `days` , set a `TokenValidityUnits` value in your API request.
	//
	// For example, when you set `RefreshTokenValidity` as `10` and `TokenValidityUnits` as `days` , your user can refresh their session
	// and retrieve new access and ID tokens for 10 days.
	//
	// The default time unit for `RefreshTokenValidity` in an API request is days. You can't set `RefreshTokenValidity` to 0. If you do, Amazon Cognito overrides the value with the default value of 30 days. *Valid range* is displayed below in seconds.
	//
	// If you don't specify otherwise in the configuration of your app client, your refresh
	// tokens are valid for 30 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-refreshtokenvalidity
	//
	RefreshTokenValidity *float64 `field:"optional" json:"refreshTokenValidity" yaml:"refreshTokenValidity"`
	// A list of provider names for the identity providers (IdPs) that are supported on this client.
	//
	// The following are supported: `COGNITO` , `Facebook` , `Google` , `SignInWithApple` , and `LoginWithAmazon` . You can also specify the names that you configured for the SAML and OIDC IdPs in your user pool, for example `MySAMLIdP` or `MyOIDCIdP` .
	//
	// This parameter sets the IdPs that [managed login](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html) will display on the login page for your app client. The removal of `COGNITO` from this list doesn't prevent authentication operations for local users with the user pools API in an AWS SDK. The only way to prevent SDK-based authentication is to block access with a [AWS WAF rule](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-waf.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-supportedidentityproviders
	//
	SupportedIdentityProviders *[]*string `field:"optional" json:"supportedIdentityProviders" yaml:"supportedIdentityProviders"`
	// The units that validity times are represented in.
	//
	// The default unit for refresh tokens is days, and the default for ID and access tokens are hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-tokenvalidityunits
	//
	TokenValidityUnits interface{} `field:"optional" json:"tokenValidityUnits" yaml:"tokenValidityUnits"`
	// The list of user attributes that you want your app client to have write access to.
	//
	// After your user authenticates in your app, their access token authorizes them to set or modify their own attribute value for any attribute in this list.
	//
	// An example of this kind of activity is when you present your user with a form to update their profile information and they change their last name. Your app then makes an `API_UpdateUserAttributes` API request and sets `family_name` to the new value.
	//
	// When you don't specify the `WriteAttributes` for your app client, your app can write the values of the Standard attributes of your user pool. When your user pool has write access to these default attributes, `WriteAttributes` doesn't return any information. Amazon Cognito only populates `WriteAttributes` in the API response if you have specified your own custom set of write attributes.
	//
	// If your app client allows users to sign in through an IdP, this array must include all attributes that you have mapped to IdP attributes. Amazon Cognito updates mapped attributes when users sign in to your application through an IdP. If your app client does not have write access to a mapped attribute, Amazon Cognito throws an error when it tries to update the attribute. For more information, see [Specifying IdP Attribute Mappings for Your user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-writeattributes
	//
	WriteAttributes *[]*string `field:"optional" json:"writeAttributes" yaml:"writeAttributes"`
}

