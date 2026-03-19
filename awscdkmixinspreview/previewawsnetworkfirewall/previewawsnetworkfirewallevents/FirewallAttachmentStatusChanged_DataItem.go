package previewawsnetworkfirewallevents


// Type definition for DataItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataItem := &DataItem{
//   	AvailabilityZone: []*string{
//   		jsii.String("availabilityZone"),
//   	},
//   	CurrentAttachmentStatus: []*string{
//   		jsii.String("currentAttachmentStatus"),
//   	},
//   	EndpointId: []*string{
//   		jsii.String("endpointId"),
//   	},
//   	PreviousAttachmentStatus: []*string{
//   		jsii.String("previousAttachmentStatus"),
//   	},
//   }
//
// Experimental.
type FirewallAttachmentStatusChanged_DataItem struct {
	// Availability Zone property.
	//
	// Specify an array of string values to match this event if the actual value of Availability Zone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZone *[]*string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Current Attachment Status property.
	//
	// Specify an array of string values to match this event if the actual value of Current Attachment Status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentAttachmentStatus *[]*string `field:"optional" json:"currentAttachmentStatus" yaml:"currentAttachmentStatus"`
	// Endpoint ID property.
	//
	// Specify an array of string values to match this event if the actual value of Endpoint ID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndpointId *[]*string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// Previous Attachment Status property.
	//
	// Specify an array of string values to match this event if the actual value of Previous Attachment Status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousAttachmentStatus *[]*string `field:"optional" json:"previousAttachmentStatus" yaml:"previousAttachmentStatus"`
}

