package mixinsawsapplicationsignals


// The time window to be excluded from the SLO performance metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   exclusionWindowProperty := &ExclusionWindowProperty{
//   	Reason: jsii.String("reason"),
//   	RecurrenceRule: &RecurrenceRuleProperty{
//   		Expression: jsii.String("expression"),
//   	},
//   	StartTime: jsii.String("startTime"),
//   	Window: &WindowProperty{
//   		Duration: jsii.Number(123),
//   		DurationUnit: jsii.String("durationUnit"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html
//
type CfnServiceLevelObjectivePropsMixin_ExclusionWindowProperty struct {
	// The reason for the time exclusion windows.
	//
	// For example, maintenance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html#cfn-applicationsignals-servicelevelobjective-exclusionwindow-reason
	//
	// Default: - "No reason".
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// The recurrence rule for the time exclusion window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html#cfn-applicationsignals-servicelevelobjective-exclusionwindow-recurrencerule
	//
	RecurrenceRule interface{} `field:"optional" json:"recurrenceRule" yaml:"recurrenceRule"`
	// The start time of the time exclusion window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html#cfn-applicationsignals-servicelevelobjective-exclusionwindow-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The time exclusion window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.html#cfn-applicationsignals-servicelevelobjective-exclusionwindow-window
	//
	Window interface{} `field:"optional" json:"window" yaml:"window"`
}

