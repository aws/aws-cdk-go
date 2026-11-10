package awslambda


// Attributes of an existing network connector, for importing it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConnectorAttributes := &NetworkConnectorAttributes{
//   	NetworkConnectorArn: jsii.String("networkConnectorArn"),
//
//   	// the properties below are optional
//   	LatestVersionArn: jsii.String("latestVersionArn"),
//   }
//
type NetworkConnectorAttributes struct {
	// The unqualified ARN of the connector.
	NetworkConnectorArn *string `field:"required" json:"networkConnectorArn" yaml:"networkConnectorArn"`
	// The qualified ARN of the version a WebFunction should pin to.
	// Default: - not known; reading `latestVersionArn` on the imported connector throws.
	//
	LatestVersionArn *string `field:"optional" json:"latestVersionArn" yaml:"latestVersionArn"`
}

