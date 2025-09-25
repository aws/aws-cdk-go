package awsec2


// A reference to a LocalGatewayRoute resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localGatewayRouteReference := &LocalGatewayRouteReference{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	LocalGatewayRouteTableId: jsii.String("localGatewayRouteTableId"),
//   }
//
type LocalGatewayRouteReference struct {
	// The DestinationCidrBlock of the LocalGatewayRoute resource.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The LocalGatewayRouteTableId of the LocalGatewayRoute resource.
	LocalGatewayRouteTableId *string `field:"required" json:"localGatewayRouteTableId" yaml:"localGatewayRouteTableId"`
}

