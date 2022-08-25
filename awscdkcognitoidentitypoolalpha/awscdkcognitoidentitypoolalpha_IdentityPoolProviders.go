// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha


// External Identity Providers To Connect to User Pools and Identity Pools.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cognito_identitypool_alpha "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//
//   identityPoolProviders := &identityPoolProviders{
//   	amazon: &identityPoolAmazonLoginProvider{
//   		appId: jsii.String("appId"),
//   	},
//   	apple: &identityPoolAppleLoginProvider{
//   		servicesId: jsii.String("servicesId"),
//   	},
//   	digits: &identityPoolDigitsLoginProvider{
//   		consumerKey: jsii.String("consumerKey"),
//   		consumerSecret: jsii.String("consumerSecret"),
//   	},
//   	facebook: &identityPoolFacebookLoginProvider{
//   		appId: jsii.String("appId"),
//   	},
//   	google: &identityPoolGoogleLoginProvider{
//   		clientId: jsii.String("clientId"),
//   	},
//   	twitter: &identityPoolTwitterLoginProvider{
//   		consumerKey: jsii.String("consumerKey"),
//   		consumerSecret: jsii.String("consumerSecret"),
//   	},
//   }
//
// Experimental.
type IdentityPoolProviders struct {
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
}

