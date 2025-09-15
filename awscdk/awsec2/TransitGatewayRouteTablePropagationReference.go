package awsec2


// A reference to a TransitGatewayRouteTablePropagation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayRouteTablePropagationReference := &TransitGatewayRouteTablePropagationReference{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	TransitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//   }
//
type TransitGatewayRouteTablePropagationReference struct {
	// The TransitGatewayAttachmentId of the TransitGatewayRouteTablePropagation resource.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The TransitGatewayRouteTableId of the TransitGatewayRouteTablePropagation resource.
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
}

