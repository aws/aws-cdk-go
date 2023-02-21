package awsappflow


// The API key credentials required for API key authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyCredentialsProperty := &apiKeyCredentialsProperty{
//   	apiKey: jsii.String("apiKey"),
//
//   	// the properties below are optional
//   	apiSecretKey: jsii.String("apiSecretKey"),
//   }
//
type CfnConnectorProfile_ApiKeyCredentialsProperty struct {
	// The API key required for API key authentication.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The API secret key required for API key authentication.
	ApiSecretKey *string `field:"optional" json:"apiSecretKey" yaml:"apiSecretKey"`
}

