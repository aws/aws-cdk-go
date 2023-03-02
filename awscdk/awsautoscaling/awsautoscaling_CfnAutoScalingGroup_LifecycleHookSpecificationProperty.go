package awsautoscaling


// `LifecycleHookSpecification` specifies a lifecycle hook for the `LifecycleHookSpecificationList` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource. A lifecycle hook specifies actions to perform when Amazon EC2 Auto Scaling launches or terminates instances.
//
// For more information, see [Amazon EC2 Auto Scaling lifecycle hooks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html) in the *Amazon EC2 Auto Scaling User Guide* . You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#aws-resource-as-lifecyclehook--examples) section of the `AWS::AutoScaling::LifecycleHook` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleHookSpecificationProperty := &lifecycleHookSpecificationProperty{
//   	lifecycleHookName: jsii.String("lifecycleHookName"),
//   	lifecycleTransition: jsii.String("lifecycleTransition"),
//
//   	// the properties below are optional
//   	defaultResult: jsii.String("defaultResult"),
//   	heartbeatTimeout: jsii.Number(123),
//   	notificationMetadata: jsii.String("notificationMetadata"),
//   	notificationTargetArn: jsii.String("notificationTargetArn"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnAutoScalingGroup_LifecycleHookSpecificationProperty struct {
	// The name of the lifecycle hook.
	LifecycleHookName *string `field:"required" json:"lifecycleHookName" yaml:"lifecycleHookName"`
	// The lifecycle transition. For Auto Scaling groups, there are two major lifecycle transitions.
	//
	// - To create a lifecycle hook for scale-out events, specify `autoscaling:EC2_INSTANCE_LAUNCHING` .
	// - To create a lifecycle hook for scale-in events, specify `autoscaling:EC2_INSTANCE_TERMINATING` .
	LifecycleTransition *string `field:"required" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	//
	// The default value is `ABANDON` .
	//
	// Valid values: `CONTINUE` | `ABANDON`.
	DefaultResult *string `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// The maximum time, in seconds, that can elapse before the lifecycle hook times out.
	//
	// The range is from `30` to `7200` seconds. The default value is `3600` seconds (1 hour).
	HeartbeatTimeout *float64 `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// Additional information that you want to include any time Amazon EC2 Auto Scaling sends a message to the notification target.
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling sends notifications to when an instance is in a wait state for the lifecycle hook.
	//
	// You can specify an Amazon SNS topic or an Amazon SQS queue.
	NotificationTargetArn *string `field:"optional" json:"notificationTargetArn" yaml:"notificationTargetArn"`
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	//
	// For information about creating this role, see [Configure a notification target for a lifecycle hook](https://docs.aws.amazon.com/autoscaling/ec2/userguide/prepare-for-lifecycle-notifications.html#lifecycle-hook-notification-target) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Valid only if the notification target is an Amazon SNS topic or an Amazon SQS queue.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

