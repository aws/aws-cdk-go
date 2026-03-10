package previewawsappmeshevents


// Type definition for ServiceEventDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceEventDetails := &ServiceEventDetails{
//   	ConnectionId: []*string{
//   		jsii.String("connectionId"),
//   	},
//   	EventStatus: []*string{
//   		jsii.String("eventStatus"),
//   	},
//   	FailureReason: []*string{
//   		jsii.String("failureReason"),
//   	},
//   	NodeArn: []*string{
//   		jsii.String("nodeArn"),
//   	},
//   }
//
// Experimental.
type AWSServiceEventViaCloudTrail_ServiceEventDetails struct {
	// connectionId property.
	//
	// Specify an array of string values to match this event if the actual value of connectionId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ConnectionId *[]*string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// eventStatus property.
	//
	// Specify an array of string values to match this event if the actual value of eventStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventStatus *[]*string `field:"optional" json:"eventStatus" yaml:"eventStatus"`
	// failureReason property.
	//
	// Specify an array of string values to match this event if the actual value of failureReason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailureReason *[]*string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// nodeArn property.
	//
	// Specify an array of string values to match this event if the actual value of nodeArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NodeArn *[]*string `field:"optional" json:"nodeArn" yaml:"nodeArn"`
}

