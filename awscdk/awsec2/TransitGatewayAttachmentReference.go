package awsec2


// A reference to a TransitGatewayAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayAttachmentReference := &TransitGatewayAttachmentReference{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }
//
type TransitGatewayAttachmentReference struct {
	// The Id of the TransitGatewayAttachment resource.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
}

