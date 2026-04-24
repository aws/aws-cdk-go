package awsdeadline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var alwaysScheduleFirst interface{}
//   var alwaysScheduleLast interface{}
//   var priorityFifo interface{}
//
//   schedulingConfigurationProperty := &SchedulingConfigurationProperty{
//   	PriorityBalanced: &PriorityBalancedSchedulingConfigurationProperty{
//   		RenderingTaskBuffer: jsii.Number(123),
//   	},
//   	PriorityFifo: priorityFifo,
//   	WeightedBalanced: &WeightedBalancedSchedulingConfigurationProperty{
//   		ErrorWeight: jsii.Number(123),
//   		MaxPriorityOverride: &SchedulingMaxPriorityOverrideProperty{
//   			AlwaysScheduleFirst: alwaysScheduleFirst,
//   		},
//   		MinPriorityOverride: &SchedulingMinPriorityOverrideProperty{
//   			AlwaysScheduleLast: alwaysScheduleLast,
//   		},
//   		PriorityWeight: jsii.Number(123),
//   		RenderingTaskBuffer: jsii.Number(123),
//   		RenderingTaskWeight: jsii.Number(123),
//   		SubmissionTimeWeight: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-schedulingconfiguration.html
//
type CfnQueuePropsMixin_SchedulingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-schedulingconfiguration.html#cfn-deadline-queue-schedulingconfiguration-prioritybalanced
	//
	PriorityBalanced interface{} `field:"optional" json:"priorityBalanced" yaml:"priorityBalanced"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-schedulingconfiguration.html#cfn-deadline-queue-schedulingconfiguration-priorityfifo
	//
	PriorityFifo interface{} `field:"optional" json:"priorityFifo" yaml:"priorityFifo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-schedulingconfiguration.html#cfn-deadline-queue-schedulingconfiguration-weightedbalanced
	//
	WeightedBalanced interface{} `field:"optional" json:"weightedBalanced" yaml:"weightedBalanced"`
}

