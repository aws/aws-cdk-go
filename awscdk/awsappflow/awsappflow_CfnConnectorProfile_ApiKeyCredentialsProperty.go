package awsappflow


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
	// `CfnConnectorProfile.ApiKeyCredentialsProperty.ApiKey`.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// `CfnConnectorProfile.ApiKeyCredentialsProperty.ApiSecretKey`.
	ApiSecretKey *string `field:"optional" json:"apiSecretKey" yaml:"apiSecretKey"`
}

