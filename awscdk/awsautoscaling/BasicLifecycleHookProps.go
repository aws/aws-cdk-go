package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Basic properties for a lifecycle hook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var lifecycleHookTarget iLifecycleHookTarget
//   var role role
//
//   basicLifecycleHookProps := &BasicLifecycleHookProps{
//   	LifecycleTransition: awscdk.Aws_autoscaling.LifecycleTransition_INSTANCE_LAUNCHING,
//
//   	// the properties below are optional
//   	DefaultResult: awscdk.*Aws_autoscaling.DefaultResult_CONTINUE,
//   	HeartbeatTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	LifecycleHookName: jsii.String("lifecycleHookName"),
//   	NotificationMetadata: jsii.String("notificationMetadata"),
//   	NotificationTarget: lifecycleHookTarget,
//   	Role: role,
//   }
//
type BasicLifecycleHookProps struct {
	// The state of the Amazon EC2 instance to which you want to attach the lifecycle hook.
	LifecycleTransition LifecycleTransition `field:"required" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	// Default: Continue.
	//
	DefaultResult DefaultResult `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// Maximum time between calls to RecordLifecycleActionHeartbeat for the hook.
	//
	// If the lifecycle hook times out, perform the action in DefaultResult.
	// Default: - No heartbeat timeout.
	//
	HeartbeatTimeout awscdk.Duration `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// Name of the lifecycle hook.
	// Default: - Automatically generated name.
	//
	LifecycleHookName *string `field:"optional" json:"lifecycleHookName" yaml:"lifecycleHookName"`
	// Additional data to pass to the lifecycle hook target.
	// Default: - No metadata.
	//
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// The target of the lifecycle hook.
	// Default: - No target.
	//
	NotificationTarget ILifecycleHookTarget `field:"optional" json:"notificationTarget" yaml:"notificationTarget"`
	// The role that allows publishing to the notification target.
	// Default: - A role will be created if a target is provided. Otherwise, no role is created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

