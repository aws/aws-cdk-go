package awsec2


// Properties for defining a `CfnTransitGatewayRouteTableAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayRouteTableAssociationProps := &cfnTransitGatewayRouteTableAssociationProps{
//   	transitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	transitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//   }
//
type CfnTransitGatewayRouteTableAssociationProps struct {
	// The ID of the attachment.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The ID of the route table for the transit gateway.
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
}

