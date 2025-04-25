package awsec2alpha


// Common properties for a Transit Gateway Route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//
//   var transitGatewayRouteTable transitGatewayRouteTable
//
//   baseTransitGatewayRouteProps := &BaseTransitGatewayRouteProps{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	TransitGatewayRouteTable: transitGatewayRouteTable,
//
//   	// the properties below are optional
//   	TransitGatewayRouteName: jsii.String("transitGatewayRouteName"),
//   }
//
// Experimental.
type BaseTransitGatewayRouteProps struct {
	// The destination CIDR block for this route.
	//
	// Destination Cidr cannot overlap for static routes but is allowed for propagated routes.
	// When overlapping occurs, static routes take precedence over propagated routes.
	// Experimental.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The transit gateway route table you want to install this route into.
	// Experimental.
	TransitGatewayRouteTable ITransitGatewayRouteTable `field:"required" json:"transitGatewayRouteTable" yaml:"transitGatewayRouteTable"`
	// Physical name of this Transit Gateway Route.
	// Default: - Assigned by CloudFormation.
	//
	// Experimental.
	TransitGatewayRouteName *string `field:"optional" json:"transitGatewayRouteName" yaml:"transitGatewayRouteName"`
}

