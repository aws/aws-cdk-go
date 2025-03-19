package awsapplicationsignals


// If the interval for this SLO is a rolling interval, this structure contains the interval specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rollingIntervalProperty := &RollingIntervalProperty{
//   	Duration: jsii.Number(123),
//   	DurationUnit: jsii.String("durationUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-rollinginterval.html
//
type CfnServiceLevelObjective_RollingIntervalProperty struct {
	// Specifies the duration of each rolling interval.
	//
	// For example, if `Duration` is `7` and `DurationUnit` is `DAY` , each rolling interval is seven days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-rollinginterval.html#cfn-applicationsignals-servicelevelobjective-rollinginterval-duration
	//
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Specifies the rolling interval unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-rollinginterval.html#cfn-applicationsignals-servicelevelobjective-rollinginterval-durationunit
	//
	DurationUnit *string `field:"required" json:"durationUnit" yaml:"durationUnit"`
}

