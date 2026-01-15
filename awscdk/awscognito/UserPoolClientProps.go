package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito"
)

// Properties for the UserPoolClient construct.
//
// Example:
//   import pinpoint "github.com/aws/aws-cdk-go/awscdk"
//
//   var userPool UserPool
//   var pinpointApp CfnApp
//   var pinpointRole Role
//
//
//   cognito.NewUserPoolClient(this, jsii.String("Client"), &UserPoolClientProps{
//   	UserPool: UserPool,
//   	Analytics: &AnalyticsConfiguration{
//   		// Your Pinpoint project ID
//   		ApplicationId: pinpointApp.ref,
//
//   		// External ID for the IAM role
//   		ExternalId: jsii.String("sample-external-id"),
//
//   		// IAM role that Cognito can assume to publish to Pinpoint
//   		Role: pinpointRole,
//
//   		// Whether to include user data in analytics events
//   		ShareUserData: jsii.Boolean(true),
//   	},
//   })
//
type UserPoolClientProps struct {
	// Validity of the access token.
	//
	// Values between 5 minutes and 1 day are valid. The duration can not be longer than the refresh token validity.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-access-token
	//
	// Default: Duration.minutes(60)
	//
	AccessTokenValidity awscdk.Duration `field:"optional" json:"accessTokenValidity" yaml:"accessTokenValidity"`
	// The analytics configuration for this client.
	// Default: - no analytics configuration.
	//
	Analytics *AnalyticsConfiguration `field:"optional" json:"analytics" yaml:"analytics"`
	// The set of OAuth authentication flows to enable on the client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow.html
	//
	// Default: - If you don't specify a value, your user client supports ALLOW_REFRESH_TOKEN_AUTH, ALLOW_USER_SRP_AUTH, and ALLOW_CUSTOM_AUTH.
	//
	AuthFlows *AuthFlow `field:"optional" json:"authFlows" yaml:"authFlows"`
	// Cognito creates a session token for each API request in an authentication flow.
	//
	// AuthSessionValidity is the duration, in minutes, of that session token.
	// see defaults in `AuthSessionValidity`. Valid duration is from 3 to 15 minutes.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html#cfn-cognito-userpoolclient-authsessionvalidity
	//
	// Default: - Duration.minutes(3)
	//
	AuthSessionValidity awscdk.Duration `field:"optional" json:"authSessionValidity" yaml:"authSessionValidity"`
	// Turns off all OAuth interactions for this client.
	// Default: false.
	//
	DisableOAuth *bool `field:"optional" json:"disableOAuth" yaml:"disableOAuth"`
	// Enable the propagation of additional user context data.
	//
	// You can only activate enablePropagateAdditionalUserContextData in an app client that has a client secret.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-adaptive-authentication.html#user-pool-settings-adaptive-authentication-device-fingerprint
	//
	// Default: false for new user pool clients.
	//
	EnablePropagateAdditionalUserContextData *bool `field:"optional" json:"enablePropagateAdditionalUserContextData" yaml:"enablePropagateAdditionalUserContextData"`
	// Enable token revocation for this client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/token-revocation.html#enable-token-revocation
	//
	// Default: true for new user pool clients.
	//
	EnableTokenRevocation *bool `field:"optional" json:"enableTokenRevocation" yaml:"enableTokenRevocation"`
	// Whether to generate a client secret.
	// Default: false.
	//
	GenerateSecret *bool `field:"optional" json:"generateSecret" yaml:"generateSecret"`
	// Validity of the ID token.
	//
	// Values between 5 minutes and 1 day are valid. The duration can not be longer than the refresh token validity.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-id-token
	//
	// Default: Duration.minutes(60)
	//
	IdTokenValidity awscdk.Duration `field:"optional" json:"idTokenValidity" yaml:"idTokenValidity"`
	// OAuth settings for this client to interact with the app.
	//
	// An error is thrown when this is specified and `disableOAuth` is set.
	// Default: - see defaults in `OAuthSettings`. meaningless if `disableOAuth` is set.
	//
	OAuth *OAuthSettings `field:"optional" json:"oAuth" yaml:"oAuth"`
	// Whether Cognito returns a UserNotFoundException exception when the user does not exist in the user pool (false), or whether it returns another type of error that doesn't reveal the user's absence.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-managing-errors.html
	//
	// Default: false.
	//
	PreventUserExistenceErrors *bool `field:"optional" json:"preventUserExistenceErrors" yaml:"preventUserExistenceErrors"`
	// The set of attributes this client will be able to read.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-attribute-permissions-and-scopes
	//
	// Default: - all standard and custom attributes.
	//
	ReadAttributes ClientAttributes `field:"optional" json:"readAttributes" yaml:"readAttributes"`
	// Enables refresh token rotation when set.
	//
	// Defines the grace period for the original refresh token (0-60 seconds).
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-the-refresh-token.html#using-the-refresh-token-rotation
	//
	// Default: - undefined (refresh token rotation is disabled).
	//
	RefreshTokenRotationGracePeriod awscdk.Duration `field:"optional" json:"refreshTokenRotationGracePeriod" yaml:"refreshTokenRotationGracePeriod"`
	// Validity of the refresh token.
	//
	// Values between 60 minutes and 10 years are valid.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-refresh-token
	//
	// Default: Duration.days(30)
	//
	RefreshTokenValidity awscdk.Duration `field:"optional" json:"refreshTokenValidity" yaml:"refreshTokenValidity"`
	// The list of identity providers that users should be able to use to sign in using this client.
	// Default: - supports all identity providers that are registered with the user pool. If the user pool and/or
	// identity providers are imported, either specify this option explicitly or ensure that the identity providers are
	// registered with the user pool using the `UserPool.registerIdentityProvider()` API.
	//
	SupportedIdentityProviders *[]UserPoolClientIdentityProvider `field:"optional" json:"supportedIdentityProviders" yaml:"supportedIdentityProviders"`
	// Name of the application client.
	// Default: - cloudformation generated name.
	//
	UserPoolClientName *string `field:"optional" json:"userPoolClientName" yaml:"userPoolClientName"`
	// The set of attributes this client will be able to write.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-attribute-permissions-and-scopes
	//
	// Default: - all standard and custom attributes.
	//
	WriteAttributes ClientAttributes `field:"optional" json:"writeAttributes" yaml:"writeAttributes"`
	// The UserPool resource this client will have access to.
	UserPool interfacesawscognito.IUserPoolRef `field:"required" json:"userPool" yaml:"userPool"`
}

