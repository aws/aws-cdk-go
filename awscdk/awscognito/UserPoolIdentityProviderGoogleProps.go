package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties to initialize UserPoolGoogleIdentityProvider.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//   secret := secretsmanager.Secret_FromSecretAttributes(this, jsii.String("CognitoClientSecret"), &SecretAttributes{
//   	SecretCompleteArn: jsii.String("arn:aws:secretsmanager:xxx:xxx:secret:xxx-xxx"),
//   }).SecretValue
//
//   provider := cognito.NewUserPoolIdentityProviderGoogle(this, jsii.String("Google"), &UserPoolIdentityProviderGoogleProps{
//   	ClientId: jsii.String("amzn-client-id"),
//   	ClientSecretValue: secret,
//   	UserPool: userpool,
//   })
//
type UserPoolIdentityProviderGoogleProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Default: - no attribute mapping.
	//
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by Google APIs.
	// See: https://developers.google.com/identity/sign-in/web/sign-in#specify_your_apps_client_id
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret to be accompanied with clientId for Google APIs to authenticate the client.
	// See: https://developers.google.com/identity/sign-in/web/sign-in
	//
	// Default: none.
	//
	// Deprecated: use clientSecretValue instead.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The client secret to be accompanied with clientId for Google APIs to authenticate the client as SecretValue.
	// See: https://developers.google.com/identity/sign-in/web/sign-in
	//
	// Default: none.
	//
	ClientSecretValue awscdk.SecretValue `field:"optional" json:"clientSecretValue" yaml:"clientSecretValue"`
	// The list of Google permissions to obtain for getting access to the Google profile.
	// See: https://developers.google.com/identity/sign-in/web/sign-in
	//
	// Default: [ profile ].
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

