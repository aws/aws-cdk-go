package awsec2alpha


// Common properties for a Transit Gateway Route Table Association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//
//   var transitGatewayAttachment iTransitGatewayAttachment
//   var transitGatewayRouteTable transitGatewayRouteTable
//
//   transitGatewayRouteTableAssociationProps := &TransitGatewayRouteTableAssociationProps{
//   	TransitGatewayRouteTable: transitGatewayRouteTable,
//   	TransitGatewayVpcAttachment: transitGatewayAttachment,
//
//   	// the properties below are optional
//   	TransitGatewayRouteTableAssociationName: jsii.String("transitGatewayRouteTableAssociationName"),
//   }
//
// Experimental.
type TransitGatewayRouteTableAssociationProps struct {
	// The ID of the transit gateway route table association.
	// Experimental.
	TransitGatewayRouteTable ITransitGatewayRouteTable `field:"required" json:"transitGatewayRouteTable" yaml:"transitGatewayRouteTable"`
	// The ID of the transit gateway route table association.
	// Experimental.
	TransitGatewayVpcAttachment ITransitGatewayAttachment `field:"required" json:"transitGatewayVpcAttachment" yaml:"transitGatewayVpcAttachment"`
	// Physical name of this association.
	// Default: - Assigned by CloudFormation.
	//
	// Experimental.
	TransitGatewayRouteTableAssociationName *string `field:"optional" json:"transitGatewayRouteTableAssociationName" yaml:"transitGatewayRouteTableAssociationName"`
}

