package awscognito


// Properties for defining a `CfnUserPoolClient`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolClientProps := &cfnUserPoolClientProps{
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	accessTokenValidity: jsii.Number(123),
//   	allowedOAuthFlows: []*string{
//   		jsii.String("allowedOAuthFlows"),
//   	},
//   	allowedOAuthFlowsUserPoolClient: jsii.Boolean(false),
//   	allowedOAuthScopes: []*string{
//   		jsii.String("allowedOAuthScopes"),
//   	},
//   	analyticsConfiguration: &analyticsConfigurationProperty{
//   		applicationArn: jsii.String("applicationArn"),
//   		applicationId: jsii.String("applicationId"),
//   		externalId: jsii.String("externalId"),
//   		roleArn: jsii.String("roleArn"),
//   		userDataShared: jsii.Boolean(false),
//   	},
//   	authSessionValidity: jsii.Number(123),
//   	callbackUrLs: []*string{
//   		jsii.String("callbackUrLs"),
//   	},
//   	clientName: jsii.String("clientName"),
//   	defaultRedirectUri: jsii.String("defaultRedirectUri"),
//   	enablePropagateAdditionalUserContextData: jsii.Boolean(false),
//   	enableTokenRevocation: jsii.Boolean(false),
//   	explicitAuthFlows: []*string{
//   		jsii.String("explicitAuthFlows"),
//   	},
//   	generateSecret: jsii.Boolean(false),
//   	idTokenValidity: jsii.Number(123),
//   	logoutUrLs: []*string{
//   		jsii.String("logoutUrLs"),
//   	},
//   	preventUserExistenceErrors: jsii.String("preventUserExistenceErrors"),
//   	readAttributes: []*string{
//   		jsii.String("readAttributes"),
//   	},
//   	refreshTokenValidity: jsii.Number(123),
//   	supportedIdentityProviders: []*string{
//   		jsii.String("supportedIdentityProviders"),
//   	},
//   	tokenValidityUnits: &tokenValidityUnitsProperty{
//   		accessToken: jsii.String("accessToken"),
//   		idToken: jsii.String("idToken"),
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	writeAttributes: []*string{
//   		jsii.String("writeAttributes"),
//   	},
//   }
//
type CfnUserPoolClientProps struct {
	// The user pool ID for the user pool where you want to create a user pool client.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// The access token time limit.
	//
	// After this limit expires, your user can't use their access token. To specify the time unit for `AccessTokenValidity` as `seconds` , `minutes` , `hours` , or `days` , set a `TokenValidityUnits` value in your API request.
	//
	// For example, when you set `AccessTokenValidity` to `10` and `TokenValidityUnits` to `hours` , your user can authorize access with their access token for 10 hours.
	//
	// The default time unit for `AccessTokenValidity` in an API request is hours.
	AccessTokenValidity *float64 `field:"optional" json:"accessTokenValidity" yaml:"accessTokenValidity"`
	// The allowed OAuth flows.
	//
	// - **code** - Use a code grant flow, which provides an authorization code as the response. This code can be exchanged for access tokens with the `/oauth2/token` endpoint.
	// - **implicit** - Issue the access token (and, optionally, ID token, based on scopes) directly to your user.
	// - **client_credentials** - Issue the access token from the `/oauth2/token` endpoint directly to a non-person user using a combination of the client ID and client secret.
	AllowedOAuthFlows *[]*string `field:"optional" json:"allowedOAuthFlows" yaml:"allowedOAuthFlows"`
	// Set to true if the client is allowed to follow the OAuth protocol when interacting with Amazon Cognito user pools.
	AllowedOAuthFlowsUserPoolClient interface{} `field:"optional" json:"allowedOAuthFlowsUserPoolClient" yaml:"allowedOAuthFlowsUserPoolClient"`
	// The allowed OAuth scopes.
	//
	// Possible values provided by OAuth are `phone` , `email` , `openid` , and `profile` . Possible values provided by AWS are `aws.cognito.signin.user.admin` . Custom scopes created in Resource Servers are also supported.
	AllowedOAuthScopes *[]*string `field:"optional" json:"allowedOAuthScopes" yaml:"allowedOAuthScopes"`
	// The user pool analytics configuration for collecting metrics and sending them to your Amazon Pinpoint campaign.
	//
	// > In AWS Regions where Amazon Pinpoint isn't available, user pools only support sending events to Amazon Pinpoint projects in AWS Region us-east-1. In Regions where Amazon Pinpoint is available, user pools support sending events to Amazon Pinpoint projects within that same Region.
	AnalyticsConfiguration interface{} `field:"optional" json:"analyticsConfiguration" yaml:"analyticsConfiguration"`
	// `AWS::Cognito::UserPoolClient.AuthSessionValidity`.
	AuthSessionValidity *float64 `field:"optional" json:"authSessionValidity" yaml:"authSessionValidity"`
	// A list of allowed redirect (callback) URLs for the IdPs.
	//
	// A redirect URI must:
	//
	// - Be an absolute URI.
	// - Be registered with the authorization server.
	// - Not include a fragment component.
	//
	// See [OAuth 2.0 - Redirection Endpoint](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6749#section-3.1.2) .
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	CallbackUrLs *[]*string `field:"optional" json:"callbackUrLs" yaml:"callbackUrLs"`
	// The client name for the user pool client you would like to create.
	ClientName *string `field:"optional" json:"clientName" yaml:"clientName"`
	// The default redirect URI. Must be in the `CallbackURLs` list.
	//
	// A redirect URI must:
	//
	// - Be an absolute URI.
	// - Be registered with the authorization server.
	// - Not include a fragment component.
	//
	// See [OAuth 2.0 - Redirection Endpoint](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6749#section-3.1.2) .
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectUri *string `field:"optional" json:"defaultRedirectUri" yaml:"defaultRedirectUri"`
	// `AWS::Cognito::UserPoolClient.EnablePropagateAdditionalUserContextData`.
	EnablePropagateAdditionalUserContextData interface{} `field:"optional" json:"enablePropagateAdditionalUserContextData" yaml:"enablePropagateAdditionalUserContextData"`
	// Activates or deactivates token revocation. For more information about revoking tokens, see [RevokeToken](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_RevokeToken.html) .
	//
	// If you don't include this parameter, token revocation is automatically activated for the new user pool client.
	EnableTokenRevocation interface{} `field:"optional" json:"enableTokenRevocation" yaml:"enableTokenRevocation"`
	// The authentication flows that are supported by the user pool clients.
	//
	// Flow names without the `ALLOW_` prefix are no longer supported, in favor of new names with the `ALLOW_` prefix.
	//
	// > Values with `ALLOW_` prefix must be used only along with the `ALLOW_` prefix.
	//
	// Valid values include:
	//
	// - `ALLOW_ADMIN_USER_PASSWORD_AUTH` : Enable admin based user password authentication flow `ADMIN_USER_PASSWORD_AUTH` . This setting replaces the `ADMIN_NO_SRP_AUTH` setting. With this authentication flow, Amazon Cognito receives the password in the request instead of using the Secure Remote Password (SRP) protocol to verify passwords.
	// - `ALLOW_CUSTOM_AUTH` : Enable AWS Lambda trigger based authentication.
	// - `ALLOW_USER_PASSWORD_AUTH` : Enable user password-based authentication. In this flow, Amazon Cognito receives the password in the request instead of using the SRP protocol to verify passwords.
	// - `ALLOW_USER_SRP_AUTH` : Enable SRP-based authentication.
	// - `ALLOW_REFRESH_TOKEN_AUTH` : Enable authflow to refresh tokens.
	//
	// If you don't specify a value for `ExplicitAuthFlows` , your app client activates the `ALLOW_USER_SRP_AUTH` and `ALLOW_CUSTOM_AUTH` authentication flows.
	ExplicitAuthFlows *[]*string `field:"optional" json:"explicitAuthFlows" yaml:"explicitAuthFlows"`
	// Boolean to specify whether you want to generate a secret for the user pool client being created.
	GenerateSecret interface{} `field:"optional" json:"generateSecret" yaml:"generateSecret"`
	// The ID token time limit.
	//
	// After this limit expires, your user can't use their ID token. To specify the time unit for `IdTokenValidity` as `seconds` , `minutes` , `hours` , or `days` , set a `TokenValidityUnits` value in your API request.
	//
	// For example, when you set `IdTokenValidity` as `10` and `TokenValidityUnits` as `hours` , your user can authenticate their session with their ID token for 10 hours.
	//
	// The default time unit for `AccessTokenValidity` in an API request is hours.
	IdTokenValidity *float64 `field:"optional" json:"idTokenValidity" yaml:"idTokenValidity"`
	// A list of allowed logout URLs for the IdPs.
	LogoutUrLs *[]*string `field:"optional" json:"logoutUrLs" yaml:"logoutUrLs"`
	// Use this setting to choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool.
	//
	// When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY` , those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors *string `field:"optional" json:"preventUserExistenceErrors" yaml:"preventUserExistenceErrors"`
	// The read attributes.
	ReadAttributes *[]*string `field:"optional" json:"readAttributes" yaml:"readAttributes"`
	// The refresh token time limit.
	//
	// After this limit expires, your user can't use their refresh token. To specify the time unit for `RefreshTokenValidity` as `seconds` , `minutes` , `hours` , or `days` , set a `TokenValidityUnits` value in your API request.
	//
	// For example, when you set `RefreshTokenValidity` as `10` and `TokenValidityUnits` as `days` , your user can refresh their session and retrieve new access and ID tokens for 10 days.
	//
	// The default time unit for `RefreshTokenValidity` in an API request is days. You can't set `RefreshTokenValidity` to 0. If you do, Amazon Cognito overrides the value with the default value of 30 days.
	RefreshTokenValidity *float64 `field:"optional" json:"refreshTokenValidity" yaml:"refreshTokenValidity"`
	// A list of provider names for the IdPs that this client supports.
	//
	// The following are supported: `COGNITO` , `Facebook` , `Google` `LoginWithAmazon` , and the names of your own SAML and OIDC providers.
	SupportedIdentityProviders *[]*string `field:"optional" json:"supportedIdentityProviders" yaml:"supportedIdentityProviders"`
	// The units in which the validity times are represented.
	//
	// The default unit for RefreshToken is days, and default for ID and access tokens are hours.
	TokenValidityUnits interface{} `field:"optional" json:"tokenValidityUnits" yaml:"tokenValidityUnits"`
	// The user pool attributes that the app client can write to.
	//
	// If your app client allows users to sign in through an IdP, this array must include all attributes that you have mapped to IdP attributes. Amazon Cognito updates mapped attributes when users sign in to your application through an IdP. If your app client does not have write access to a mapped attribute, Amazon Cognito throws an error when it tries to update the attribute. For more information, see [Specifying IdP Attribute Mappings for Your user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html) .
	WriteAttributes *[]*string `field:"optional" json:"writeAttributes" yaml:"writeAttributes"`
}

