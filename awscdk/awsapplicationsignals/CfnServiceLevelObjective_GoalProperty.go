package awsapplicationsignals


// This structure contains the attributes that determine the goal of an SLO.
//
// This includes the time period for evaluation and the attainment threshold.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   goalProperty := &GoalProperty{
//   	AttainmentGoal: jsii.Number(123),
//   	Interval: &IntervalProperty{
//   		CalendarInterval: &CalendarIntervalProperty{
//   			Duration: jsii.Number(123),
//   			DurationUnit: jsii.String("durationUnit"),
//   			StartTime: jsii.Number(123),
//   		},
//   		RollingInterval: &RollingIntervalProperty{
//   			Duration: jsii.Number(123),
//   			DurationUnit: jsii.String("durationUnit"),
//   		},
//   	},
//   	WarningThreshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-goal.html
//
type CfnServiceLevelObjective_GoalProperty struct {
	// The threshold that determines if the goal is being met.
	//
	// If this is a period-based SLO, the attainment goal is the percentage of good periods that meet the threshold requirements to the total periods within the interval. For example, an attainment goal of 99.9% means that within your interval, you are targeting 99.9% of the periods to be in healthy state.
	//
	// If this is a request-based SLO, the attainment goal is the percentage of requests that must be successful to meet the attainment goal.
	//
	// If you omit this parameter, 99 is used to represent 99% as the attainment goal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-goal.html#cfn-applicationsignals-servicelevelobjective-goal-attainmentgoal
	//
	AttainmentGoal *float64 `field:"optional" json:"attainmentGoal" yaml:"attainmentGoal"`
	// The time period used to evaluate the SLO. It can be either a calendar interval or rolling interval.
	//
	// If you omit this parameter, a rolling interval of 7 days is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-goal.html#cfn-applicationsignals-servicelevelobjective-goal-interval
	//
	Interval interface{} `field:"optional" json:"interval" yaml:"interval"`
	// The percentage of remaining budget over total budget that you want to get warnings for.
	//
	// If you omit this parameter, the default of 50.0 is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-goal.html#cfn-applicationsignals-servicelevelobjective-goal-warningthreshold
	//
	WarningThreshold *float64 `field:"optional" json:"warningThreshold" yaml:"warningThreshold"`
}

