package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for a Lifecycle hook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var duration duration
//   var lifecycleHookTarget iLifecycleHookTarget
//   var role role
//
//   lifecycleHookProps := &LifecycleHookProps{
//   	AutoScalingGroup: autoScalingGroup,
//   	LifecycleTransition: awscdk.Aws_autoscaling.LifecycleTransition_INSTANCE_LAUNCHING,
//
//   	// the properties below are optional
//   	DefaultResult: awscdk.*Aws_autoscaling.DefaultResult_CONTINUE,
//   	HeartbeatTimeout: duration,
//   	LifecycleHookName: jsii.String("lifecycleHookName"),
//   	NotificationMetadata: jsii.String("notificationMetadata"),
//   	NotificationTarget: lifecycleHookTarget,
//   	Role: role,
//   }
//
// Experimental.
type LifecycleHookProps struct {
	// The state of the Amazon EC2 instance to which you want to attach the lifecycle hook.
	// Experimental.
	LifecycleTransition LifecycleTransition `field:"required" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	// Experimental.
	DefaultResult DefaultResult `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// Maximum time between calls to RecordLifecycleActionHeartbeat for the hook.
	//
	// If the lifecycle hook times out, perform the action in DefaultResult.
	// Experimental.
	HeartbeatTimeout awscdk.Duration `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// Name of the lifecycle hook.
	// Experimental.
	LifecycleHookName *string `field:"optional" json:"lifecycleHookName" yaml:"lifecycleHookName"`
	// Additional data to pass to the lifecycle hook target.
	// Experimental.
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// The target of the lifecycle hook.
	// Experimental.
	NotificationTarget ILifecycleHookTarget `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// The role that allows publishing to the notification target.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The AutoScalingGroup to add the lifecycle hook to.
	// Experimental.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

