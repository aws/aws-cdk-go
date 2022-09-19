// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Authentication providers for using in identity pool.
//
// Example:
//   var openIdConnectProvider openIdConnectProvider
//
//   awscdkcognitoidentitypoolalpha.NewIdentityPool(this, jsii.String("myidentitypool"), &identityPoolProps{
//   	identityPoolName: jsii.String("myidentitypool"),
//   	authenticationProviders: &identityPoolAuthenticationProviders{
//   		google: &identityPoolGoogleLoginProvider{
//   			clientId: jsii.String("12345678012.apps.googleusercontent.com"),
//   		},
//   		openIdConnectProviders: []iOpenIdConnectProvider{
//   			openIdConnectProvider,
//   		},
//   		customProvider: jsii.String("my-custom-provider.example.com"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/external-identity-providers.html
//
// Experimental.
type IdentityPoolAuthenticationProviders struct {
	// App Id for Amazon Identity Federation.
	// Experimental.
	Amazon *IdentityPoolAmazonLoginProvider `field:"optional" json:"amazon" yaml:"amazon"`
	// Services Id for Apple Identity Federation.
	// Experimental.
	Apple *IdentityPoolAppleLoginProvider `field:"optional" json:"apple" yaml:"apple"`
	// Consumer Key and Secret for Digits Identity Federation.
	// Experimental.
	Digits *IdentityPoolDigitsLoginProvider `field:"optional" json:"digits" yaml:"digits"`
	// App Id for Facebook Identity Federation.
	// Experimental.
	Facebook *IdentityPoolFacebookLoginProvider `field:"optional" json:"facebook" yaml:"facebook"`
	// Client Id for Google Identity Federation.
	// Experimental.
	Google *IdentityPoolGoogleLoginProvider `field:"optional" json:"google" yaml:"google"`
	// Consumer Key and Secret for Twitter Identity Federation.
	// Experimental.
	Twitter *IdentityPoolTwitterLoginProvider `field:"optional" json:"twitter" yaml:"twitter"`
	// The Developer Provider Name to associate with this Identity Pool.
	// Experimental.
	CustomProvider *string `field:"optional" json:"customProvider" yaml:"customProvider"`
	// The OpenIdConnect Provider associated with this Identity Pool.
	// Experimental.
	OpenIdConnectProviders *[]awsiam.IOpenIdConnectProvider `field:"optional" json:"openIdConnectProviders" yaml:"openIdConnectProviders"`
	// The Security Assertion Markup Language Provider associated with this Identity Pool.
	// Experimental.
	SamlProviders *[]awsiam.ISamlProvider `field:"optional" json:"samlProviders" yaml:"samlProviders"`
	// The User Pool Authentication Providers associated with this Identity Pool.
	// Experimental.
	UserPools *[]IUserPoolAuthenticationProvider `field:"optional" json:"userPools" yaml:"userPools"`
}

