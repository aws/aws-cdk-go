package awsmedialive


// A reference to a Network resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkReference := &NetworkReference{
//   	NetworkArn: jsii.String("networkArn"),
//   	NetworkId: jsii.String("networkId"),
//   }
//
type NetworkReference struct {
	// The ARN of the Network resource.
	NetworkArn *string `field:"required" json:"networkArn" yaml:"networkArn"`
	// The Id of the Network resource.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
}

