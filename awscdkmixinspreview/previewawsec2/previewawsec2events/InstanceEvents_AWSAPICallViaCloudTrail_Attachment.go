package previewawsec2events


// Type definition for Attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attachment := &Attachment{
//   	AttachmentId: []*string{
//   		jsii.String("attachmentId"),
//   	},
//   	AttachTime: []*string{
//   		jsii.String("attachTime"),
//   	},
//   	DeleteOnTermination: []*string{
//   		jsii.String("deleteOnTermination"),
//   	},
//   	DeviceIndex: []*string{
//   		jsii.String("deviceIndex"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_Attachment struct {
	// attachmentId property.
	//
	// Specify an array of string values to match this event if the actual value of attachmentId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AttachmentId *[]*string `field:"optional" json:"attachmentId" yaml:"attachmentId"`
	// attachTime property.
	//
	// Specify an array of string values to match this event if the actual value of attachTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AttachTime *[]*string `field:"optional" json:"attachTime" yaml:"attachTime"`
	// deleteOnTermination property.
	//
	// Specify an array of string values to match this event if the actual value of deleteOnTermination is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeleteOnTermination *[]*string `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// deviceIndex property.
	//
	// Specify an array of string values to match this event if the actual value of deviceIndex is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeviceIndex *[]*string `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}

