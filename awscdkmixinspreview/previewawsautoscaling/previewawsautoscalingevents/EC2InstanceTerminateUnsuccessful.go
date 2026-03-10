package previewawsautoscalingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.autoscaling@EC2InstanceTerminateUnsuccessful.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceTerminateUnsuccessful := awscdkmixinspreview.Events.NewEC2InstanceTerminateUnsuccessful()
//
// Experimental.
type EC2InstanceTerminateUnsuccessful interface {
}

// The jsii proxy struct for EC2InstanceTerminateUnsuccessful
type jsiiProxy_EC2InstanceTerminateUnsuccessful struct {
	_ byte // padding
}

// Experimental.
func NewEC2InstanceTerminateUnsuccessful() EC2InstanceTerminateUnsuccessful {
	_init_.Initialize()

	j := jsiiProxy_EC2InstanceTerminateUnsuccessful{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateUnsuccessful",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2InstanceTerminateUnsuccessful_Override(e EC2InstanceTerminateUnsuccessful) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateUnsuccessful",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Instance Terminate Unsuccessful.
// Experimental.
func EC2InstanceTerminateUnsuccessful_Ec2InstanceTerminateUnsuccessfulPattern(options *EC2InstanceTerminateUnsuccessful_EC2InstanceTerminateUnsuccessfulProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2InstanceTerminateUnsuccessful_Ec2InstanceTerminateUnsuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateUnsuccessful",
		"ec2InstanceTerminateUnsuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

