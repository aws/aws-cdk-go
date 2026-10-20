package awsappfabric


// Contains API key credential information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeyCredentialProperty := &ApiKeyCredentialProperty{
//   	ApiKey: jsii.String("apiKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-apikeycredential.html
//
type CfnAppAuthorization_ApiKeyCredentialProperty struct {
	// An API key for an application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-appauthorization-apikeycredential.html#cfn-appfabric-appauthorization-apikeycredential-apikey
	//
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
}

