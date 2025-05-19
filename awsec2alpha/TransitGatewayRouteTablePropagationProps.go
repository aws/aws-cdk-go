package awsec2alpha


// Common properties for a Transit Gateway Route Table Propagation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//
//   var transitGatewayAttachment iTransitGatewayAttachment
//   var transitGatewayRouteTable transitGatewayRouteTable
//
//   transitGatewayRouteTablePropagationProps := &TransitGatewayRouteTablePropagationProps{
//   	TransitGatewayRouteTable: transitGatewayRouteTable,
//   	TransitGatewayVpcAttachment: transitGatewayAttachment,
//
//   	// the properties below are optional
//   	TransitGatewayRouteTablePropagationName: jsii.String("transitGatewayRouteTablePropagationName"),
//   }
//
// Experimental.
type TransitGatewayRouteTablePropagationProps struct {
	// The ID of the transit gateway route table propagation.
	// Experimental.
	TransitGatewayRouteTable ITransitGatewayRouteTable `field:"required" json:"transitGatewayRouteTable" yaml:"transitGatewayRouteTable"`
	// The ID of the transit gateway route table propagation.
	// Experimental.
	TransitGatewayVpcAttachment ITransitGatewayAttachment `field:"required" json:"transitGatewayVpcAttachment" yaml:"transitGatewayVpcAttachment"`
	// Physical name of this propagation.
	// Default: - Assigned by CloudFormation.
	//
	// Experimental.
	TransitGatewayRouteTablePropagationName *string `field:"optional" json:"transitGatewayRouteTablePropagationName" yaml:"transitGatewayRouteTablePropagationName"`
}

