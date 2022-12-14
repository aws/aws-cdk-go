package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   peeringAttachmentStatusProperty := &peeringAttachmentStatusProperty{
//   	code: jsii.String("code"),
//   	message: jsii.String("message"),
//   }
//
type CfnTransitGatewayPeeringAttachment_PeeringAttachmentStatusProperty struct {
	// `CfnTransitGatewayPeeringAttachment.PeeringAttachmentStatusProperty.Code`.
	Code *string `field:"optional" json:"code" yaml:"code"`
	// `CfnTransitGatewayPeeringAttachment.PeeringAttachmentStatusProperty.Message`.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

