package previewawsbatchevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   responseElements := &ResponseElements{
//   	JobId: []*string{
//   		jsii.String("jobId"),
//   	},
//   	JobName: []*string{
//   		jsii.String("jobName"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// jobId property.
	//
	// Specify an array of string values to match this event if the actual value of jobId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobId *[]*string `field:"optional" json:"jobId" yaml:"jobId"`
	// jobName property.
	//
	// Specify an array of string values to match this event if the actual value of jobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobName *[]*string `field:"optional" json:"jobName" yaml:"jobName"`
}

