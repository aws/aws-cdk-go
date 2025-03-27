package awsapplicationsignals


// The core SLO time window exclusion object that includes Window, StartTime, RecurrenceRule, and Reason.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exclusionWindowProperty := &ExclusionWindowProperty{
//   	Window: &WindowProperty{
//   		Duration: jsii.Number(123),
//   		DurationUnit: jsii.String("durationUnit"),
//   	},
//
//   	// the properties below are optional
//   	Reason: jsii.String("reason"),
//   	RecurrenceRule: &RecurrenceRuleProperty{
//   		Expression: jsii.String("expression"),
//   	},
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html
//
type CfnServiceLevelObjective_ExclusionWindowProperty struct {
	// The SLO time window exclusion .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html#cfn-applicationsignals-servicelevelobjective-exclusionwindow-window
	//
	Window interface{} `field:"required" json:"window" yaml:"window"`
	// A description explaining why this time period should be excluded from SLO calculations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html#cfn-applicationsignals-servicelevelobjective-exclusionwindow-reason
	//
	// Default: - "No reason".
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// The recurrence rule for the SLO time window exclusion.
	//
	// Supports both cron and rate expressions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html#cfn-applicationsignals-servicelevelobjective-exclusionwindow-recurrencerule
	//
	RecurrenceRule interface{} `field:"optional" json:"recurrenceRule" yaml:"recurrenceRule"`
	// The start of the SLO time window exclusion.
	//
	// Defaults to current time if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html#cfn-applicationsignals-servicelevelobjective-exclusionwindow-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

