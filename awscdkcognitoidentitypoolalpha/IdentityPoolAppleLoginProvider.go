// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha


// Login Provider for Identity Federation using Apple Credentials.
//
// Example:
//   awscdkcognitoidentitypoolalpha.NewIdentityPool(this, jsii.String("myidentitypool"), &IdentityPoolProps{
//   	IdentityPoolName: jsii.String("myidentitypool"),
//   	AuthenticationProviders: &IdentityPoolAuthenticationProviders{
//   		Amazon: &IdentityPoolAmazonLoginProvider{
//   			AppId: jsii.String("amzn1.application.12312k3j234j13rjiwuenf"),
//   		},
//   		Facebook: &IdentityPoolFacebookLoginProvider{
//   			AppId: jsii.String("1234567890123"),
//   		},
//   		Google: &IdentityPoolGoogleLoginProvider{
//   			ClientId: jsii.String("12345678012.apps.googleusercontent.com"),
//   		},
//   		Apple: &IdentityPoolAppleLoginProvider{
//   			ServicesId: jsii.String("com.myappleapp.auth"),
//   		},
//   		Twitter: &IdentityPoolTwitterLoginProvider{
//   			ConsumerKey: jsii.String("my-twitter-id"),
//   			ConsumerSecret: jsii.String("my-twitter-secret"),
//   		},
//   	},
//   })
//
// Experimental.
type IdentityPoolAppleLoginProvider struct {
	// App Id for Apple Identity Federation.
	// Experimental.
	ServicesId *string `field:"required" json:"servicesId" yaml:"servicesId"`
}

