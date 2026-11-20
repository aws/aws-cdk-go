package interfacesawsec2


// A reference to a TransitGatewayNetworkFunctionAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayNetworkFunctionAttachmentReference := &TransitGatewayNetworkFunctionAttachmentReference{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }
//
type TransitGatewayNetworkFunctionAttachmentReference struct {
	// The TransitGatewayAttachmentId of the TransitGatewayNetworkFunctionAttachment resource.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
}

