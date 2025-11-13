package interfacesawsec2


// A reference to a LocalGatewayRouteTable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localGatewayRouteTableReference := &LocalGatewayRouteTableReference{
//   	LocalGatewayRouteTableArn: jsii.String("localGatewayRouteTableArn"),
//   	LocalGatewayRouteTableId: jsii.String("localGatewayRouteTableId"),
//   }
//
type LocalGatewayRouteTableReference struct {
	// The ARN of the LocalGatewayRouteTable resource.
	LocalGatewayRouteTableArn *string `field:"required" json:"localGatewayRouteTableArn" yaml:"localGatewayRouteTableArn"`
	// The LocalGatewayRouteTableId of the LocalGatewayRouteTable resource.
	LocalGatewayRouteTableId *string `field:"required" json:"localGatewayRouteTableId" yaml:"localGatewayRouteTableId"`
}

