package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ec2@EC2SpotInstanceInterruptionWarning event types for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2SpotInstanceInterruptionWarning := #error#.NewEC2SpotInstanceInterruptionWarning()
//
// Experimental.
type InstanceEvents_EC2SpotInstanceInterruptionWarning interface {
}

// The jsii proxy struct for InstanceEvents_EC2SpotInstanceInterruptionWarning
type jsiiProxy_InstanceEvents_EC2SpotInstanceInterruptionWarning struct {
	_ byte // padding
}

// Experimental.
func NewInstanceEvents_EC2SpotInstanceInterruptionWarning() InstanceEvents_EC2SpotInstanceInterruptionWarning {
	_init_.Initialize()

	j := jsiiProxy_InstanceEvents_EC2SpotInstanceInterruptionWarning{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.InstanceEvents.EC2SpotInstanceInterruptionWarning",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEvents_EC2SpotInstanceInterruptionWarning_Override(i InstanceEvents_EC2SpotInstanceInterruptionWarning) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.InstanceEvents.EC2SpotInstanceInterruptionWarning",
		nil, // no parameters
		i,
	)
}

