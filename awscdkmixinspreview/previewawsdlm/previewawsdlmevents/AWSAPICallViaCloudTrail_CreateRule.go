package previewawsdlmevents


// Type definition for CreateRule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   createRule := &CreateRule{
//   	Interval: []*string{
//   		jsii.String("interval"),
//   	},
//   	IntervalUnit: []*string{
//   		jsii.String("intervalUnit"),
//   	},
//   	Times: []*string{
//   		jsii.String("times"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_CreateRule struct {
	// Interval property.
	//
	// Specify an array of string values to match this event if the actual value of Interval is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Interval *[]*string `field:"optional" json:"interval" yaml:"interval"`
	// IntervalUnit property.
	//
	// Specify an array of string values to match this event if the actual value of IntervalUnit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IntervalUnit *[]*string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// Times property.
	//
	// Specify an array of string values to match this event if the actual value of Times is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Times *[]*string `field:"optional" json:"times" yaml:"times"`
}

