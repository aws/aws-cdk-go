package awsec2


// A reference to a TransitGatewayRouteTableAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayRouteTableAssociationReference := &TransitGatewayRouteTableAssociationReference{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	TransitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//   }
//
type TransitGatewayRouteTableAssociationReference struct {
	// The TransitGatewayAttachmentId of the TransitGatewayRouteTableAssociation resource.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The TransitGatewayRouteTableId of the TransitGatewayRouteTableAssociation resource.
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
}

