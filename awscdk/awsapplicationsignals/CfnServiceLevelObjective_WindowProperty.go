package awsapplicationsignals


// The object that defines the time length of an exclusion window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   windowProperty := &WindowProperty{
//   	Duration: jsii.Number(123),
//   	DurationUnit: jsii.String("durationUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-window.html
//
type CfnServiceLevelObjective_WindowProperty struct {
	// The number of time units for the exclusion window length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-window.html#cfn-applicationsignals-servicelevelobjective-window-duration
	//
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// The unit of time for the exclusion window duration.
	//
	// Valid values: MINUTE, HOUR, DAY, MONTH.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-window.html#cfn-applicationsignals-servicelevelobjective-window-durationunit
	//
	DurationUnit *string `field:"required" json:"durationUnit" yaml:"durationUnit"`
}

