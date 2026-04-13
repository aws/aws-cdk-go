package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.ec2@EC2SpotInstanceInterruptionWarning.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2SpotInstanceInterruptionWarning := awscdkmixinspreview.Events.NewEC2SpotInstanceInterruptionWarning()
//
// Experimental.
type EC2SpotInstanceInterruptionWarning interface {
}

// The jsii proxy struct for EC2SpotInstanceInterruptionWarning
type jsiiProxy_EC2SpotInstanceInterruptionWarning struct {
	_ byte // padding
}

// Experimental.
func NewEC2SpotInstanceInterruptionWarning() EC2SpotInstanceInterruptionWarning {
	_init_.Initialize()

	j := jsiiProxy_EC2SpotInstanceInterruptionWarning{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EC2SpotInstanceInterruptionWarning",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewEC2SpotInstanceInterruptionWarning_Override(e EC2SpotInstanceInterruptionWarning) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.EC2SpotInstanceInterruptionWarning",
		nil, // no parameters
		e,
	)
}

// EventBridge event pattern for EC2 Spot Instance Interruption Warning.
// Experimental.
func EC2SpotInstanceInterruptionWarning_EventPattern(options *EC2SpotInstanceInterruptionWarning_EC2SpotInstanceInterruptionWarningProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateEC2SpotInstanceInterruptionWarning_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.events.EC2SpotInstanceInterruptionWarning",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

