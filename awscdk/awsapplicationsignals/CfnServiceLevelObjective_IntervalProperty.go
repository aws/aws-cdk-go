package awsapplicationsignals


// The time period used to evaluate the SLO.
//
// It can be either a calendar interval or rolling interval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intervalProperty := &IntervalProperty{
//   	CalendarInterval: &CalendarIntervalProperty{
//   		Duration: jsii.Number(123),
//   		DurationUnit: jsii.String("durationUnit"),
//   		StartTime: jsii.Number(123),
//   	},
//   	RollingInterval: &RollingIntervalProperty{
//   		Duration: jsii.Number(123),
//   		DurationUnit: jsii.String("durationUnit"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-interval.html
//
type CfnServiceLevelObjective_IntervalProperty struct {
	// If the interval is a calendar interval, this structure contains the interval specifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-interval.html#cfn-applicationsignals-servicelevelobjective-interval-calendarinterval
	//
	CalendarInterval interface{} `field:"optional" json:"calendarInterval" yaml:"calendarInterval"`
	// If the interval is a rolling interval, this structure contains the interval specifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-interval.html#cfn-applicationsignals-servicelevelobjective-interval-rollinginterval
	//
	RollingInterval interface{} `field:"optional" json:"rollingInterval" yaml:"rollingInterval"`
}

