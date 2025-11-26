package previewawslogsevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	LogGroupName: []*string{
//   		jsii.String("logGroupName"),
//   	},
//   	LogStreamName: []*string{
//   		jsii.String("logStreamName"),
//   	},
//   }
//
// Experimental.
type LogGroupEvents_AWSAPICallViaCloudTrail_RequestParameters struct {
	// logGroupName property.
	//
	// Specify an array of string values to match this event if the actual value of logGroupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the LogGroup reference.
	//
	// Experimental.
	LogGroupName *[]*string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// logStreamName property.
	//
	// Specify an array of string values to match this event if the actual value of logStreamName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LogStreamName *[]*string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

