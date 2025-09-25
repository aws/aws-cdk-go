package awsnetworkmanager


// A reference to a TransitGatewayRouteTableAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayRouteTableAttachmentReference := &TransitGatewayRouteTableAttachmentReference{
//   	AttachmentId: jsii.String("attachmentId"),
//   }
//
type TransitGatewayRouteTableAttachmentReference struct {
	// The AttachmentId of the TransitGatewayRouteTableAttachment resource.
	AttachmentId *string `field:"required" json:"attachmentId" yaml:"attachmentId"`
}

