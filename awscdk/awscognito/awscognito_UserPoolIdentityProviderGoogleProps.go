package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties to initialize UserPoolGoogleIdentityProvider.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//   secret := secretsManager.secret_FromSecretAttributes(this, jsii.String("CognitoClientSecret"), map[string]*string{
//   	"secretCompleteArn": jsii.String("arn:aws:secretsmanager:xxx:xxx:secret:xxx-xxx"),
//   }).secretValue
//
//   provider := cognito.NewUserPoolIdentityProviderGoogle(this, jsii.String("Google"), &userPoolIdentityProviderGoogleProps{
//   	clientId: jsii.String("amzn-client-id"),
//   	clientSecretValue: secret,
//   	userPool: userpool,
//   })
//
type UserPoolIdentityProviderGoogleProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The client id recognized by Google APIs.
	// See: https://developers.google.com/identity/sign-in/web/sign-in#specify_your_apps_client_id
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret to be accompanied with clientId for Google APIs to authenticate the client.
	// See: https://developers.google.com/identity/sign-in/web/sign-in
	//
	// Deprecated: use clientSecretValue instead.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The client secret to be accompanied with clientId for Google APIs to authenticate the client as SecretValue.
	// See: https://developers.google.com/identity/sign-in/web/sign-in
	//
	ClientSecretValue awscdk.SecretValue `field:"optional" json:"clientSecretValue" yaml:"clientSecretValue"`
	// The list of google permissions to obtain for getting access to the google profile.
	// See: https://developers.google.com/identity/sign-in/web/sign-in
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

