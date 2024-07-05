package awsapplicationsignals


// If the interval for this service level objective is a calendar interval, this structure contains the interval specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calendarIntervalProperty := &CalendarIntervalProperty{
//   	Duration: jsii.Number(123),
//   	DurationUnit: jsii.String("durationUnit"),
//   	StartTime: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-calendarinterval.html
//
type CfnServiceLevelObjective_CalendarIntervalProperty struct {
	// Specifies the duration of each calendar interval.
	//
	// For example, if `Duration` is `1` and `DurationUnit` is `MONTH` , each interval is one month, aligned with the calendar.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-calendarinterval.html#cfn-applicationsignals-servicelevelobjective-calendarinterval-duration
	//
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// Specifies the calendar interval unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-calendarinterval.html#cfn-applicationsignals-servicelevelobjective-calendarinterval-durationunit
	//
	DurationUnit *string `field:"required" json:"durationUnit" yaml:"durationUnit"`
	// The date and time when you want the first interval to start.
	//
	// Be sure to choose a time that configures the intervals the way that you want. For example, if you want weekly intervals starting on Mondays at 6 a.m., be sure to specify a start time that is a Monday at 6 a.m.
	//
	// When used in a raw HTTP Query API, it is formatted as be epoch time in seconds. For example: `1698778057`
	//
	// As soon as one calendar interval ends, another automatically begins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-calendarinterval.html#cfn-applicationsignals-servicelevelobjective-calendarinterval-starttime
	//
	StartTime *float64 `field:"required" json:"startTime" yaml:"startTime"`
}

