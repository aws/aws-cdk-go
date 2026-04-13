package previewawsautoscalingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.autoscaling@EC2InstanceTerminateLifecycleAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceTerminateLifecycleAction := awscdkmixinspreview.Events.NewEC2InstanceTerminateLifecycleAction()
//
// Experimental.
type EC2InstanceTerminateLifecycleAction interface {
}

// The jsii proxy struct for EC2InstanceTerminateLifecycleAction
type jsiiProxy_EC2InstanceTerminateLifecycleAction struct {
	_ byte // padding
}

// Experimental.
func NewEC2InstanceTerminateLifecycleAction() EC2InstanceTerminateLifecycleAction {
	_init_.Initialize()

	j := jsiiProxy_EC2InstanceTerminateLifecycleAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateLifecycleAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2InstanceTerminateLifecycleAction_Override(e EC2InstanceTerminateLifecycleAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateLifecycleAction",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Instance-terminate Lifecycle Action.
// Experimental.
func EC2InstanceTerminateLifecycleAction_EventPattern(options *EC2InstanceTerminateLifecycleAction_EC2InstanceTerminateLifecycleActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2InstanceTerminateLifecycleAction_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateLifecycleAction",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

