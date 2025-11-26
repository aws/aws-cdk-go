package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.opsworks@OpsWorksAlert event types for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksAlert := #error#.NewOpsWorksAlert()
//
// Experimental.
type InstanceEvents_OpsWorksAlert interface {
}

// The jsii proxy struct for InstanceEvents_OpsWorksAlert
type jsiiProxy_InstanceEvents_OpsWorksAlert struct {
	_ byte // padding
}

// Experimental.
func NewInstanceEvents_OpsWorksAlert() InstanceEvents_OpsWorksAlert {
	_init_.Initialize()

	j := jsiiProxy_InstanceEvents_OpsWorksAlert{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksAlert",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEvents_OpsWorksAlert_Override(i InstanceEvents_OpsWorksAlert) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksAlert",
		nil, // no parameters
		i,
	)
}

