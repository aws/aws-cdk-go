// The CDK Construct Library for AWS::Cognito Identity Pools
package awscdkcognitoidentitypoolalpha


// Login Provider for Identity Federation using Digits Credentials.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cognito_identitypool_alpha "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//
//   identityPoolDigitsLoginProvider := &identityPoolDigitsLoginProvider{
//   	consumerKey: jsii.String("consumerKey"),
//   	consumerSecret: jsii.String("consumerSecret"),
//   }
//
// Experimental.
type IdentityPoolDigitsLoginProvider struct {
	// App Id for Twitter Identity Federation.
	// Experimental.
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// App Secret for Twitter Identity Federation.
	// Experimental.
	ConsumerSecret *string `field:"required" json:"consumerSecret" yaml:"consumerSecret"`
}

