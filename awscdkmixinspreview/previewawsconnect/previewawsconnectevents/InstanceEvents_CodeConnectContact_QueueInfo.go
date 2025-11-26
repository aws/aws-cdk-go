package previewawsconnectevents


// Type definition for QueueInfo.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   queueInfo := &QueueInfo{
//   	QueueArn: []*string{
//   		jsii.String("queueArn"),
//   	},
//   	QueueType: []*string{
//   		jsii.String("queueType"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_CodeConnectContact_QueueInfo struct {
	// queueArn property.
	//
	// Specify an array of string values to match this event if the actual value of queueArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueueArn *[]*string `field:"optional" json:"queueArn" yaml:"queueArn"`
	// queueType property.
	//
	// Specify an array of string values to match this event if the actual value of queueType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	QueueType *[]*string `field:"optional" json:"queueType" yaml:"queueType"`
}

