package previewawsdlmevents


// Type definition for FastRestoreRule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fastRestoreRule := &FastRestoreRule{
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	Count: []*string{
//   		jsii.String("count"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_FastRestoreRule struct {
	// AvailabilityZones property.
	//
	// Specify an array of string values to match this event if the actual value of AvailabilityZones is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Count property.
	//
	// Specify an array of string values to match this event if the actual value of Count is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Count *[]*string `field:"optional" json:"count" yaml:"count"`
}

