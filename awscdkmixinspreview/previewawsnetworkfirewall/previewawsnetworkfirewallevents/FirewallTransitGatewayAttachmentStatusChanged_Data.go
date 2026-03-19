package previewawsnetworkfirewallevents


// Type definition for Data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   data := &Data{
//   	AttachmentId: []*string{
//   		jsii.String("attachmentId"),
//   	},
//   	CurrentTransitGatewayAttachmentStatus: []*string{
//   		jsii.String("currentTransitGatewayAttachmentStatus"),
//   	},
//   	PreviousTransitGatewayAttachmentStatus: []*string{
//   		jsii.String("previousTransitGatewayAttachmentStatus"),
//   	},
//   }
//
// Experimental.
type FirewallTransitGatewayAttachmentStatusChanged_Data struct {
	// Attachment ID property.
	//
	// Specify an array of string values to match this event if the actual value of Attachment ID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AttachmentId *[]*string `field:"optional" json:"attachmentId" yaml:"attachmentId"`
	// Current Transit Gateway Attachment Status property.
	//
	// Specify an array of string values to match this event if the actual value of Current Transit Gateway Attachment Status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentTransitGatewayAttachmentStatus *[]*string `field:"optional" json:"currentTransitGatewayAttachmentStatus" yaml:"currentTransitGatewayAttachmentStatus"`
	// Previous Transit Gateway Attachment Status property.
	//
	// Specify an array of string values to match this event if the actual value of Previous Transit Gateway Attachment Status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousTransitGatewayAttachmentStatus *[]*string `field:"optional" json:"previousTransitGatewayAttachmentStatus" yaml:"previousTransitGatewayAttachmentStatus"`
}

