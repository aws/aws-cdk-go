package previewawssyntheticsevents


// Type definition for Canary-run-timeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   canaryRunTimeline := &CanaryRunTimeline{
//   	Completed: []*string{
//   		jsii.String("completed"),
//   	},
//   	Started: []*string{
//   		jsii.String("started"),
//   	},
//   }
//
// Experimental.
type SyntheticsCanaryTestRunSuccessful_CanaryRunTimeline struct {
	// completed property.
	//
	// Specify an array of string values to match this event if the actual value of completed is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Completed *[]*string `field:"optional" json:"completed" yaml:"completed"`
	// started property.
	//
	// Specify an array of string values to match this event if the actual value of started is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Started *[]*string `field:"optional" json:"started" yaml:"started"`
}

