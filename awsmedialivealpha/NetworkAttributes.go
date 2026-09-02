package awsmedialivealpha


// Attributes for importing an existing Network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   networkAttributes := &NetworkAttributes{
//   	NetworkArn: jsii.String("networkArn"),
//   	NetworkId: jsii.String("networkId"),
//   }
//
// Experimental.
type NetworkAttributes struct {
	// The ARN of the network.
	// Experimental.
	NetworkArn *string `field:"required" json:"networkArn" yaml:"networkArn"`
	// The ID of the network.
	// Experimental.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
}

