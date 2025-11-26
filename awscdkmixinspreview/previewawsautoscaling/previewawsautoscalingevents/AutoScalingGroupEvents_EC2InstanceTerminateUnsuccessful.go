package previewawsautoscalingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.autoscaling@EC2InstanceTerminateUnsuccessful event types for AutoScalingGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceTerminateUnsuccessful := #error#.NewEC2InstanceTerminateUnsuccessful()
//
// Experimental.
type AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful interface {
}

// The jsii proxy struct for AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful
type jsiiProxy_AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful struct {
	_ byte // padding
}

// Experimental.
func NewAutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful() AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful {
	_init_.Initialize()

	j := jsiiProxy_AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateUnsuccessful",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful_Override(a AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateUnsuccessful",
		nil, // no parameters
		a,
	)
}

