package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.opsworks@OpsWorksInstanceStateChange event types for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksInstanceStateChange := #error#.NewOpsWorksInstanceStateChange()
//
// Experimental.
type InstanceEvents_OpsWorksInstanceStateChange interface {
}

// The jsii proxy struct for InstanceEvents_OpsWorksInstanceStateChange
type jsiiProxy_InstanceEvents_OpsWorksInstanceStateChange struct {
	_ byte // padding
}

// Experimental.
func NewInstanceEvents_OpsWorksInstanceStateChange() InstanceEvents_OpsWorksInstanceStateChange {
	_init_.Initialize()

	j := jsiiProxy_InstanceEvents_OpsWorksInstanceStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksInstanceStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEvents_OpsWorksInstanceStateChange_Override(i InstanceEvents_OpsWorksInstanceStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.InstanceEvents.OpsWorksInstanceStateChange",
		nil, // no parameters
		i,
	)
}

