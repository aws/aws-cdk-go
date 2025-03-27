package awscdkcognitoidentitypoolalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// External Authentication Providers for usage in Identity Pool.
//
// Example:
//   var openIdConnectProvider openIdConnectProvider
//
//   awscdkcognitoidentitypoolalpha.NewIdentityPool(this, jsii.String("myidentitypool"), &IdentityPoolProps{
//   	IdentityPoolName: jsii.String("myidentitypool"),
//   	AuthenticationProviders: &IdentityPoolAuthenticationProviders{
//   		Google: &IdentityPoolGoogleLoginProvider{
//   			ClientId: jsii.String("12345678012.apps.googleusercontent.com"),
//   		},
//   		OpenIdConnectProviders: []iOpenIdConnectProvider{
//   			openIdConnectProvider,
//   		},
//   		CustomProvider: jsii.String("my-custom-provider.example.com"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/external-identity-providers.html
//
// Deprecated.
type IdentityPoolAuthenticationProviders struct {
	// The Amazon Authentication Provider associated with this Identity Pool.
	// Default: -  No Amazon Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Deprecated.
	Amazon *IdentityPoolAmazonLoginProvider `field:"optional" json:"amazon" yaml:"amazon"`
	// The Apple Authentication Provider associated with this Identity Pool.
	// Default: - No Apple Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Deprecated.
	Apple *IdentityPoolAppleLoginProvider `field:"optional" json:"apple" yaml:"apple"`
	// The developer provider name to associate with this Identity Pool.
	// Default: - no custom provider.
	//
	// Deprecated.
	CustomProvider *string `field:"optional" json:"customProvider" yaml:"customProvider"`
	// The Facebook Authentication Provider associated with this Identity Pool.
	// Default: - No Facebook Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Deprecated.
	Facebook *IdentityPoolFacebookLoginProvider `field:"optional" json:"facebook" yaml:"facebook"`
	// The Google Authentication Provider associated with this Identity Pool.
	// Default: - No Google Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Deprecated.
	Google *IdentityPoolGoogleLoginProvider `field:"optional" json:"google" yaml:"google"`
	// The OpenIdConnect Provider associated with this Identity Pool.
	// Default: - no OpenIdConnectProvider.
	//
	// Deprecated.
	OpenIdConnectProviders *[]awsiam.IOpenIdConnectProvider `field:"optional" json:"openIdConnectProviders" yaml:"openIdConnectProviders"`
	// The Security Assertion Markup Language provider associated with this Identity Pool.
	// Default: - no SamlProvider.
	//
	// Deprecated.
	SamlProviders *[]awsiam.ISamlProvider `field:"optional" json:"samlProviders" yaml:"samlProviders"`
	// The Twitter Authentication Provider associated with this Identity Pool.
	// Default: - No Twitter Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Deprecated.
	Twitter *IdentityPoolTwitterLoginProvider `field:"optional" json:"twitter" yaml:"twitter"`
	// The User Pool Authentication Providers associated with this Identity Pool.
	// Default: - no User Pools associated.
	//
	// Deprecated.
	UserPools *[]IUserPoolAuthenticationProvider `field:"optional" json:"userPools" yaml:"userPools"`
}

