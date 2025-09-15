package awsec2


// A reference to a NatGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   natGatewayReference := &NatGatewayReference{
//   	NatGatewayId: jsii.String("natGatewayId"),
//   }
//
type NatGatewayReference struct {
	// The NatGatewayId of the NatGateway resource.
	NatGatewayId *string `field:"required" json:"natGatewayId" yaml:"natGatewayId"`
}

