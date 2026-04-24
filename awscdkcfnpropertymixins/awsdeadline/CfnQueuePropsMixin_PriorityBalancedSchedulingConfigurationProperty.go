package awsdeadline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   priorityBalancedSchedulingConfigurationProperty := &PriorityBalancedSchedulingConfigurationProperty{
//   	RenderingTaskBuffer: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-prioritybalancedschedulingconfiguration.html
//
type CfnQueuePropsMixin_PriorityBalancedSchedulingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-prioritybalancedschedulingconfiguration.html#cfn-deadline-queue-prioritybalancedschedulingconfiguration-renderingtaskbuffer
	//
	// Default: - 1.
	//
	RenderingTaskBuffer *float64 `field:"optional" json:"renderingTaskBuffer" yaml:"renderingTaskBuffer"`
}

