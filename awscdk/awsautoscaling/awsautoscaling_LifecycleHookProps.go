package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a Lifecycle hook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var lifecycleHookTarget iLifecycleHookTarget
//   var role role
//
//   lifecycleHookProps := &lifecycleHookProps{
//   	autoScalingGroup: autoScalingGroup,
//   	lifecycleTransition: awscdk.Aws_autoscaling.lifecycleTransition_INSTANCE_LAUNCHING,
//
//   	// the properties below are optional
//   	defaultResult: awscdk.*Aws_autoscaling.defaultResult_CONTINUE,
//   	heartbeatTimeout: cdk.duration.minutes(jsii.Number(30)),
//   	lifecycleHookName: jsii.String("lifecycleHookName"),
//   	notificationMetadata: jsii.String("notificationMetadata"),
//   	notificationTarget: lifecycleHookTarget,
//   	role: role,
//   }
//
type LifecycleHookProps struct {
	// The state of the Amazon EC2 instance to which you want to attach the lifecycle hook.
	LifecycleTransition LifecycleTransition `field:"required" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	DefaultResult DefaultResult `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// Maximum time between calls to RecordLifecycleActionHeartbeat for the hook.
	//
	// If the lifecycle hook times out, perform the action in DefaultResult.
	HeartbeatTimeout awscdk.Duration `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// Name of the lifecycle hook.
	LifecycleHookName *string `field:"optional" json:"lifecycleHookName" yaml:"lifecycleHookName"`
	// Additional data to pass to the lifecycle hook target.
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// The target of the lifecycle hook.
	NotificationTarget ILifecycleHookTarget `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// The role that allows publishing to the notification target.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The AutoScalingGroup to add the lifecycle hook to.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

