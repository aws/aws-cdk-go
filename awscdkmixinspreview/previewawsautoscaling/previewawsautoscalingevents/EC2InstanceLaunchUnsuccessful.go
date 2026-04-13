package previewawsautoscalingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.autoscaling@EC2InstanceLaunchUnsuccessful.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceLaunchUnsuccessful := awscdkmixinspreview.Events.NewEC2InstanceLaunchUnsuccessful()
//
// Experimental.
type EC2InstanceLaunchUnsuccessful interface {
}

// The jsii proxy struct for EC2InstanceLaunchUnsuccessful
type jsiiProxy_EC2InstanceLaunchUnsuccessful struct {
	_ byte // padding
}

// Experimental.
func NewEC2InstanceLaunchUnsuccessful() EC2InstanceLaunchUnsuccessful {
	_init_.Initialize()

	j := jsiiProxy_EC2InstanceLaunchUnsuccessful{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchUnsuccessful",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2InstanceLaunchUnsuccessful_Override(e EC2InstanceLaunchUnsuccessful) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchUnsuccessful",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Instance Launch Unsuccessful.
// Experimental.
func EC2InstanceLaunchUnsuccessful_EventPattern(options *EC2InstanceLaunchUnsuccessful_EC2InstanceLaunchUnsuccessfulProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2InstanceLaunchUnsuccessful_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchUnsuccessful",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

