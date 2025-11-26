package previewawsec2events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ec2@EC2InstanceStateChangeNotification event types for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceStateChangeNotification := #error#.NewEC2InstanceStateChangeNotification()
//
// Experimental.
type InstanceEvents_EC2InstanceStateChangeNotification interface {
}

// The jsii proxy struct for InstanceEvents_EC2InstanceStateChangeNotification
type jsiiProxy_InstanceEvents_EC2InstanceStateChangeNotification struct {
	_ byte // padding
}

// Experimental.
func NewInstanceEvents_EC2InstanceStateChangeNotification() InstanceEvents_EC2InstanceStateChangeNotification {
	_init_.Initialize()

	j := jsiiProxy_InstanceEvents_EC2InstanceStateChangeNotification{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.InstanceEvents.EC2InstanceStateChangeNotification",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEvents_EC2InstanceStateChangeNotification_Override(i InstanceEvents_EC2InstanceStateChangeNotification) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.events.InstanceEvents.EC2InstanceStateChangeNotification",
		nil, // no parameters
		i,
	)
}

