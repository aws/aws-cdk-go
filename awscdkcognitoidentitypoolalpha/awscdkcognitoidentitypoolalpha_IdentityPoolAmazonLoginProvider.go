// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha


// Login Provider for Identity Federation using Amazon Credentials.
//
// Example:
//   awscdkcognitoidentitypoolalpha.NewIdentityPool(this, jsii.String("myidentitypool"), &identityPoolProps{
//   	identityPoolName: jsii.String("myidentitypool"),
//   	authenticationProviders: &identityPoolAuthenticationProviders{
//   		amazon: &identityPoolAmazonLoginProvider{
//   			appId: jsii.String("amzn1.application.12312k3j234j13rjiwuenf"),
//   		},
//   		facebook: &identityPoolFacebookLoginProvider{
//   			appId: jsii.String("1234567890123"),
//   		},
//   		google: &identityPoolGoogleLoginProvider{
//   			clientId: jsii.String("12345678012.apps.googleusercontent.com"),
//   		},
//   		apple: &identityPoolAppleLoginProvider{
//   			servicesId: jsii.String("com.myappleapp.auth"),
//   		},
//   		twitter: &identityPoolTwitterLoginProvider{
//   			consumerKey: jsii.String("my-twitter-id"),
//   			consumerSecret: jsii.String("my-twitter-secret"),
//   		},
//   	},
//   })
//
// Experimental.
type IdentityPoolAmazonLoginProvider struct {
	// App Id for Amazon Identity Federation.
	// Experimental.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
}

