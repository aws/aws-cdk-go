package previewawsautoscalingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.autoscaling@EC2InstanceTerminateSuccessful.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceTerminateSuccessful := awscdkmixinspreview.Events.NewEC2InstanceTerminateSuccessful()
//
// Experimental.
type EC2InstanceTerminateSuccessful interface {
}

// The jsii proxy struct for EC2InstanceTerminateSuccessful
type jsiiProxy_EC2InstanceTerminateSuccessful struct {
	_ byte // padding
}

// Experimental.
func NewEC2InstanceTerminateSuccessful() EC2InstanceTerminateSuccessful {
	_init_.Initialize()

	j := jsiiProxy_EC2InstanceTerminateSuccessful{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateSuccessful",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2InstanceTerminateSuccessful_Override(e EC2InstanceTerminateSuccessful) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateSuccessful",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Instance Terminate Successful.
// Experimental.
func EC2InstanceTerminateSuccessful_Ec2InstanceTerminateSuccessfulPattern(options *EC2InstanceTerminateSuccessful_EC2InstanceTerminateSuccessfulProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2InstanceTerminateSuccessful_Ec2InstanceTerminateSuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateSuccessful",
		"ec2InstanceTerminateSuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

