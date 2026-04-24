package awsdeadline


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alwaysScheduleLast interface{}
//
//   schedulingMinPriorityOverrideProperty := &SchedulingMinPriorityOverrideProperty{
//   	AlwaysScheduleLast: alwaysScheduleLast,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-schedulingminpriorityoverride.html
//
type CfnQueue_SchedulingMinPriorityOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-schedulingminpriorityoverride.html#cfn-deadline-queue-schedulingminpriorityoverride-alwaysschedulelast
	//
	AlwaysScheduleLast interface{} `field:"required" json:"alwaysScheduleLast" yaml:"alwaysScheduleLast"`
}

