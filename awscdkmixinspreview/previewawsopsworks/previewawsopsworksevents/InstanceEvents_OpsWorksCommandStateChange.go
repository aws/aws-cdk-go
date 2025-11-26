package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.opsworks@OpsWorksCommandStateChange event types for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksCommandStateChange := #error#.NewOpsWorksCommandStateChange()
//
// Experimental.
type InstanceEvents_OpsWorksCommandStateChange interface {
}

// The jsii proxy struct for InstanceEvents_OpsWorksCommandStateChange
type jsiiProxy_InstanceEvents_OpsWorksCommandStateChange struct {
	_ byte // padding
}

// Experimental.
func NewInstanceEvents_OpsWorksCommandStateChange() InstanceEvents_OpsWorksCommandStateChange {
	_init_.Initialize()

	j := jsiiProxy_InstanceEvents_OpsWorksCommandStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksCommandStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEvents_OpsWorksCommandStateChange_Override(i InstanceEvents_OpsWorksCommandStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksCommandStateChange",
		nil, // no parameters
		i,
	)
}

