package interfacesawsnetworkmanager


// A reference to a CoreNetwork resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreNetworkReference := &CoreNetworkReference{
//   	CoreNetworkArn: jsii.String("coreNetworkArn"),
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   }
//
type CoreNetworkReference struct {
	// The ARN of the CoreNetwork resource.
	CoreNetworkArn *string `field:"required" json:"coreNetworkArn" yaml:"coreNetworkArn"`
	// The CoreNetworkId of the CoreNetwork resource.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
}

