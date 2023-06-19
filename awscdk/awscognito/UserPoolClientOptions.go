package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options to create a UserPoolClient.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   pool.addClient(jsii.String("app-client"), &UserPoolClientOptions{
//   	OAuth: &OAuthSettings{
//   		Flows: &OAuthFlows{
//   			AuthorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		Scopes: []oAuthScope{
//   			cognito.*oAuthScope_OPENID(),
//   		},
//   		CallbackUrls: []*string{
//   			jsii.String("https://my-app-domain.com/welcome"),
//   		},
//   		LogoutUrls: []*string{
//   			jsii.String("https://my-app-domain.com/signin"),
//   		},
//   	},
//   })
//
// Experimental.
type UserPoolClientOptions struct {
	// Validity of the access token.
	//
	// Values between 5 minutes and 1 day are valid. The duration can not be longer than the refresh token validity.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-access-token
	//
	// Experimental.
	AccessTokenValidity awscdk.Duration `field:"optional" json:"accessTokenValidity" yaml:"accessTokenValidity"`
	// The set of OAuth authentication flows to enable on the client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow.html
	//
	// Experimental.
	AuthFlows *AuthFlow `field:"optional" json:"authFlows" yaml:"authFlows"`
	// Turns off all OAuth interactions for this client.
	// Experimental.
	DisableOAuth *bool `field:"optional" json:"disableOAuth" yaml:"disableOAuth"`
	// Enable token revocation for this client.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/token-revocation.html#enable-token-revocation
	//
	// Experimental.
	EnableTokenRevocation *bool `field:"optional" json:"enableTokenRevocation" yaml:"enableTokenRevocation"`
	// Whether to generate a client secret.
	// Experimental.
	GenerateSecret *bool `field:"optional" json:"generateSecret" yaml:"generateSecret"`
	// Validity of the ID token.
	//
	// Values between 5 minutes and 1 day are valid. The duration can not be longer than the refresh token validity.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-id-token
	//
	// Experimental.
	IdTokenValidity awscdk.Duration `field:"optional" json:"idTokenValidity" yaml:"idTokenValidity"`
	// OAuth settings for this client to interact with the app.
	//
	// An error is thrown when this is specified and `disableOAuth` is set.
	// Experimental.
	OAuth *OAuthSettings `field:"optional" json:"oAuth" yaml:"oAuth"`
	// Whether Cognito returns a UserNotFoundException exception when the user does not exist in the user pool (false), or whether it returns another type of error that doesn't reveal the user's absence.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-managing-errors.html
	//
	// Experimental.
	PreventUserExistenceErrors *bool `field:"optional" json:"preventUserExistenceErrors" yaml:"preventUserExistenceErrors"`
	// The set of attributes this client will be able to read.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-attribute-permissions-and-scopes
	//
	// Experimental.
	ReadAttributes ClientAttributes `field:"optional" json:"readAttributes" yaml:"readAttributes"`
	// Validity of the refresh token.
	//
	// Values between 60 minutes and 10 years are valid.
	// See: https://docs.aws.amazon.com/en_us/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html#amazon-cognito-user-pools-using-the-refresh-token
	//
	// Experimental.
	RefreshTokenValidity awscdk.Duration `field:"optional" json:"refreshTokenValidity" yaml:"refreshTokenValidity"`
	// The list of identity providers that users should be able to use to sign in using this client.
	// Experimental.
	SupportedIdentityProviders *[]UserPoolClientIdentityProvider `field:"optional" json:"supportedIdentityProviders" yaml:"supportedIdentityProviders"`
	// Name of the application client.
	// Experimental.
	UserPoolClientName *string `field:"optional" json:"userPoolClientName" yaml:"userPoolClientName"`
	// The set of attributes this client will be able to write.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-attribute-permissions-and-scopes
	//
	// Experimental.
	WriteAttributes ClientAttributes `field:"optional" json:"writeAttributes" yaml:"writeAttributes"`
}

