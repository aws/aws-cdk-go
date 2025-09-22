package awsnetworkmanager


// A reference to a GlobalNetwork resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalNetworkReference := &GlobalNetworkReference{
//   	GlobalNetworkArn: jsii.String("globalNetworkArn"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   }
//
type GlobalNetworkReference struct {
	// The ARN of the GlobalNetwork resource.
	GlobalNetworkArn *string `field:"required" json:"globalNetworkArn" yaml:"globalNetworkArn"`
	// The Id of the GlobalNetwork resource.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
}

