package previewawsglueevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	AllocatedCapacity: []*string{
//   		jsii.String("allocatedCapacity"),
//   	},
//   	JobName: []*string{
//   		jsii.String("jobName"),
//   	},
//   }
//
// Experimental.
type JobEvents_AWSAPICallViaCloudTrail_RequestParameters struct {
	// allocatedCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of allocatedCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AllocatedCapacity *[]*string `field:"optional" json:"allocatedCapacity" yaml:"allocatedCapacity"`
	// jobName property.
	//
	// Specify an array of string values to match this event if the actual value of jobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Job reference.
	//
	// Experimental.
	JobName *[]*string `field:"optional" json:"jobName" yaml:"jobName"`
}

