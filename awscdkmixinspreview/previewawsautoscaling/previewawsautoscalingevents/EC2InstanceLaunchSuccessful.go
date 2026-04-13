package previewawsautoscalingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.autoscaling@EC2InstanceLaunchSuccessful.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceLaunchSuccessful := awscdkmixinspreview.Events.NewEC2InstanceLaunchSuccessful()
//
// Experimental.
type EC2InstanceLaunchSuccessful interface {
}

// The jsii proxy struct for EC2InstanceLaunchSuccessful
type jsiiProxy_EC2InstanceLaunchSuccessful struct {
	_ byte // padding
}

// Experimental.
func NewEC2InstanceLaunchSuccessful() EC2InstanceLaunchSuccessful {
	_init_.Initialize()

	j := jsiiProxy_EC2InstanceLaunchSuccessful{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchSuccessful",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2InstanceLaunchSuccessful_Override(e EC2InstanceLaunchSuccessful) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchSuccessful",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Instance Launch Successful.
// Experimental.
func EC2InstanceLaunchSuccessful_EventPattern(options *EC2InstanceLaunchSuccessful_EC2InstanceLaunchSuccessfulProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2InstanceLaunchSuccessful_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchSuccessful",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

