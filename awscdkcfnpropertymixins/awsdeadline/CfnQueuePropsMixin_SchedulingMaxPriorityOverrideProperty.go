package awsdeadline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var alwaysScheduleFirst interface{}
//
//   schedulingMaxPriorityOverrideProperty := &SchedulingMaxPriorityOverrideProperty{
//   	AlwaysScheduleFirst: alwaysScheduleFirst,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-schedulingmaxpriorityoverride.html
//
type CfnQueuePropsMixin_SchedulingMaxPriorityOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-schedulingmaxpriorityoverride.html#cfn-deadline-queue-schedulingmaxpriorityoverride-alwaysschedulefirst
	//
	AlwaysScheduleFirst interface{} `field:"optional" json:"alwaysScheduleFirst" yaml:"alwaysScheduleFirst"`
}

