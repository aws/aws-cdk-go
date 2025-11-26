package previewawsautoscalingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.autoscaling@EC2InstanceTerminateLifecycleAction event types for AutoScalingGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceTerminateLifecycleAction := #error#.NewEC2InstanceTerminateLifecycleAction()
//
// Experimental.
type AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction interface {
}

// The jsii proxy struct for AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction
type jsiiProxy_AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction struct {
	_ byte // padding
}

// Experimental.
func NewAutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction() AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction {
	_init_.Initialize()

	j := jsiiProxy_AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateLifecycleAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction_Override(a AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateLifecycleAction",
		nil, // no parameters
		a,
	)
}

