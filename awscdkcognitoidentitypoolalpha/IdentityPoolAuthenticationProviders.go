package awscdkcognitoidentitypoolalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Authentication providers for using in identity pool.
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
// Experimental.
type IdentityPoolAuthenticationProviders struct {
	// App Id for Amazon Identity Federation.
	// Default: -  No Amazon Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Amazon *IdentityPoolAmazonLoginProvider `field:"optional" json:"amazon" yaml:"amazon"`
	// Services Id for Apple Identity Federation.
	// Default: - No Apple Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Apple *IdentityPoolAppleLoginProvider `field:"optional" json:"apple" yaml:"apple"`
	// Consumer Key and Secret for Digits Identity Federation.
	// Default: - No Digits Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Digits *IdentityPoolDigitsLoginProvider `field:"optional" json:"digits" yaml:"digits"`
	// App Id for Facebook Identity Federation.
	// Default: - No Facebook Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Facebook *IdentityPoolFacebookLoginProvider `field:"optional" json:"facebook" yaml:"facebook"`
	// Client Id for Google Identity Federation.
	// Default: - No Google Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Google *IdentityPoolGoogleLoginProvider `field:"optional" json:"google" yaml:"google"`
	// Consumer Key and Secret for Twitter Identity Federation.
	// Default: - No Twitter Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Twitter *IdentityPoolTwitterLoginProvider `field:"optional" json:"twitter" yaml:"twitter"`
	// The Developer Provider Name to associate with this Identity Pool.
	// Default: - no Custom Provider.
	//
	// Experimental.
	CustomProvider *string `field:"optional" json:"customProvider" yaml:"customProvider"`
	// The OpenIdConnect Provider associated with this Identity Pool.
	// Default: - no OpenIdConnectProvider.
	//
	// Experimental.
	OpenIdConnectProviders *[]awsiam.IOpenIdConnectProvider `field:"optional" json:"openIdConnectProviders" yaml:"openIdConnectProviders"`
	// The Security Assertion Markup Language Provider associated with this Identity Pool.
	// Default: - no SamlProvider.
	//
	// Experimental.
	SamlProviders *[]awsiam.ISamlProvider `field:"optional" json:"samlProviders" yaml:"samlProviders"`
	// The User Pool Authentication Providers associated with this Identity Pool.
	// Default: - no User Pools Associated.
	//
	// Experimental.
	UserPools *[]IUserPoolAuthenticationProvider `field:"optional" json:"userPools" yaml:"userPools"`
}

