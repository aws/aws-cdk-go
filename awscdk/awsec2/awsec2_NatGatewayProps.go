package awsec2


// Properties for a NAT gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   natGatewayProps := &natGatewayProps{
//   	eipAllocationIds: []*string{
//   		jsii.String("eipAllocationIds"),
//   	},
//   }
//
type NatGatewayProps struct {
	// EIP allocation IDs for the NAT gateways.
	EipAllocationIds *[]*string `field:"optional" json:"eipAllocationIds" yaml:"eipAllocationIds"`
}

