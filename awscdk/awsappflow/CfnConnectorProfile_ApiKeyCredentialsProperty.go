package awsappflow


// The API key credentials required for API key authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyCredentialsProperty := &ApiKeyCredentialsProperty{
//   	ApiKey: jsii.String("apiKey"),
//
//   	// the properties below are optional
//   	ApiSecretKey: jsii.String("apiSecretKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-apikeycredentials.html
//
type CfnConnectorProfile_ApiKeyCredentialsProperty struct {
	// The API key required for API key authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-apikeycredentials.html#cfn-appflow-connectorprofile-apikeycredentials-apikey
	//
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The API secret key required for API key authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-apikeycredentials.html#cfn-appflow-connectorprofile-apikeycredentials-apisecretkey
	//
	ApiSecretKey *string `field:"optional" json:"apiSecretKey" yaml:"apiSecretKey"`
}

