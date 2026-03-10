package previewawsautoscalingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.autoscaling@EC2InstanceLaunchLifecycleAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceLaunchLifecycleAction := awscdkmixinspreview.Events.NewEC2InstanceLaunchLifecycleAction()
//
// Experimental.
type EC2InstanceLaunchLifecycleAction interface {
}

// The jsii proxy struct for EC2InstanceLaunchLifecycleAction
type jsiiProxy_EC2InstanceLaunchLifecycleAction struct {
	_ byte // padding
}

// Experimental.
func NewEC2InstanceLaunchLifecycleAction() EC2InstanceLaunchLifecycleAction {
	_init_.Initialize()

	j := jsiiProxy_EC2InstanceLaunchLifecycleAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchLifecycleAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2InstanceLaunchLifecycleAction_Override(e EC2InstanceLaunchLifecycleAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchLifecycleAction",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Instance-launch Lifecycle Action.
// Experimental.
func EC2InstanceLaunchLifecycleAction_Ec2InstanceLaunchLifecycleActionPattern(options *EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2InstanceLaunchLifecycleAction_Ec2InstanceLaunchLifecycleActionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchLifecycleAction",
		"ec2InstanceLaunchLifecycleActionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

