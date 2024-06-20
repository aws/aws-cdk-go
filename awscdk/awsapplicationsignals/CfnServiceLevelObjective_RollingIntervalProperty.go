package awsapplicationsignals


// If the interval is a calendar interval, this structure contains the interval specifications.
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
	// Specifies the duration of each calendar interval.
	//
	// For example, if `Duration` is 1 and `DurationUnit` is `MONTH`, each interval is one month, aligned with the calendar.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-rollinginterval.html#cfn-applicationsignals-servicelevelobjective-rollinginterval-duration
	//
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Specifies the calendar interval unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-rollinginterval.html#cfn-applicationsignals-servicelevelobjective-rollinginterval-durationunit
	//
	DurationUnit *string `field:"required" json:"durationUnit" yaml:"durationUnit"`
}

