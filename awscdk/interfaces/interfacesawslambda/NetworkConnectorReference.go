package interfacesawslambda


// A reference to a NetworkConnector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConnectorReference := &NetworkConnectorReference{
//   	NetworkConnectorArn: jsii.String("networkConnectorArn"),
//   }
//
type NetworkConnectorReference struct {
	// The Arn of the NetworkConnector resource.
	NetworkConnectorArn *string `field:"required" json:"networkConnectorArn" yaml:"networkConnectorArn"`
}

