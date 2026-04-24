package awsdeadline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var alwaysScheduleFirst interface{}
//   var alwaysScheduleLast interface{}
//
//   weightedBalancedSchedulingConfigurationProperty := &WeightedBalancedSchedulingConfigurationProperty{
//   	ErrorWeight: jsii.Number(123),
//   	MaxPriorityOverride: &SchedulingMaxPriorityOverrideProperty{
//   		AlwaysScheduleFirst: alwaysScheduleFirst,
//   	},
//   	MinPriorityOverride: &SchedulingMinPriorityOverrideProperty{
//   		AlwaysScheduleLast: alwaysScheduleLast,
//   	},
//   	PriorityWeight: jsii.Number(123),
//   	RenderingTaskBuffer: jsii.Number(123),
//   	RenderingTaskWeight: jsii.Number(123),
//   	SubmissionTimeWeight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-weightedbalancedschedulingconfiguration.html
//
type CfnQueuePropsMixin_WeightedBalancedSchedulingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-weightedbalancedschedulingconfiguration.html#cfn-deadline-queue-weightedbalancedschedulingconfiguration-errorweight
	//
	// Default: - -10.
	//
	ErrorWeight *float64 `field:"optional" json:"errorWeight" yaml:"errorWeight"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-weightedbalancedschedulingconfiguration.html#cfn-deadline-queue-weightedbalancedschedulingconfiguration-maxpriorityoverride
	//
	MaxPriorityOverride interface{} `field:"optional" json:"maxPriorityOverride" yaml:"maxPriorityOverride"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-weightedbalancedschedulingconfiguration.html#cfn-deadline-queue-weightedbalancedschedulingconfiguration-minpriorityoverride
	//
	MinPriorityOverride interface{} `field:"optional" json:"minPriorityOverride" yaml:"minPriorityOverride"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-weightedbalancedschedulingconfiguration.html#cfn-deadline-queue-weightedbalancedschedulingconfiguration-priorityweight
	//
	// Default: - 100.
	//
	PriorityWeight *float64 `field:"optional" json:"priorityWeight" yaml:"priorityWeight"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-weightedbalancedschedulingconfiguration.html#cfn-deadline-queue-weightedbalancedschedulingconfiguration-renderingtaskbuffer
	//
	// Default: - 1.
	//
	RenderingTaskBuffer *float64 `field:"optional" json:"renderingTaskBuffer" yaml:"renderingTaskBuffer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-weightedbalancedschedulingconfiguration.html#cfn-deadline-queue-weightedbalancedschedulingconfiguration-renderingtaskweight
	//
	// Default: - -100.
	//
	RenderingTaskWeight *float64 `field:"optional" json:"renderingTaskWeight" yaml:"renderingTaskWeight"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-weightedbalancedschedulingconfiguration.html#cfn-deadline-queue-weightedbalancedschedulingconfiguration-submissiontimeweight
	//
	// Default: - 3.
	//
	SubmissionTimeWeight *float64 `field:"optional" json:"submissionTimeWeight" yaml:"submissionTimeWeight"`
}

