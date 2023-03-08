package awsec2


// Properties for defining a `CfnTransitGatewayRouteTablePropagation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayRouteTablePropagationProps := &cfnTransitGatewayRouteTablePropagationProps{
//   	transitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	transitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//   }
//
type CfnTransitGatewayRouteTablePropagationProps struct {
	// The ID of the attachment.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The ID of the propagation route table.
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
}

