package awsautoscaling


// Properties for defining a `CfnLifecycleHook`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLifecycleHookProps := &CfnLifecycleHookProps{
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	LifecycleTransition: jsii.String("lifecycleTransition"),
//
//   	// the properties below are optional
//   	DefaultResult: jsii.String("defaultResult"),
//   	HeartbeatTimeout: jsii.Number(123),
//   	LifecycleHookName: jsii.String("lifecycleHookName"),
//   	NotificationMetadata: jsii.String("notificationMetadata"),
//   	NotificationTargetArn: jsii.String("notificationTargetArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html
//
type CfnLifecycleHookProps struct {
	// The name of the Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-autoscalinggroupname
	//
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// The lifecycle transition. For Auto Scaling groups, there are two major lifecycle transitions.
	//
	// - To create a lifecycle hook for scale-out events, specify `autoscaling:EC2_INSTANCE_LAUNCHING` .
	// - To create a lifecycle hook for scale-in events, specify `autoscaling:EC2_INSTANCE_TERMINATING` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-lifecycletransition
	//
	LifecycleTransition *string `field:"required" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	//
	// The default value is `ABANDON` .
	//
	// Valid values: `CONTINUE` | `ABANDON`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-defaultresult
	//
	DefaultResult *string `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// The maximum time, in seconds, that can elapse before the lifecycle hook times out.
	//
	// The range is from `30` to `7200` seconds. The default value is `3600` seconds (1 hour).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-heartbeattimeout
	//
	HeartbeatTimeout *float64 `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// The name of the lifecycle hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-lifecyclehookname
	//
	LifecycleHookName *string `field:"optional" json:"lifecycleHookName" yaml:"lifecycleHookName"`
	// Additional information that you want to include any time Amazon EC2 Auto Scaling sends a message to the notification target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-notificationmetadata
	//
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling sends notifications to when an instance is in a wait state for the lifecycle hook.
	//
	// You can specify an Amazon SNS topic or an Amazon SQS queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-notificationtargetarn
	//
	NotificationTargetArn *string `field:"optional" json:"notificationTargetArn" yaml:"notificationTargetArn"`
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	//
	// For information about creating this role, see [Configure a notification target for a lifecycle hook](https://docs.aws.amazon.com/autoscaling/ec2/userguide/prepare-for-lifecycle-notifications.html#lifecycle-hook-notification-target) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Valid only if the notification target is an Amazon SNS topic or an Amazon SQS queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

