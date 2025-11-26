package previewawsopsworksevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.opsworks@OpsWorksDeploymentStateChange event types for Stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksDeploymentStateChange := #error#.NewOpsWorksDeploymentStateChange()
//
// Experimental.
type StackEvents_OpsWorksDeploymentStateChange interface {
}

// The jsii proxy struct for StackEvents_OpsWorksDeploymentStateChange
type jsiiProxy_StackEvents_OpsWorksDeploymentStateChange struct {
	_ byte // padding
}

// Experimental.
func NewStackEvents_OpsWorksDeploymentStateChange() StackEvents_OpsWorksDeploymentStateChange {
	_init_.Initialize()

	j := jsiiProxy_StackEvents_OpsWorksDeploymentStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.StackEvents.OpsWorksDeploymentStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStackEvents_OpsWorksDeploymentStateChange_Override(s StackEvents_OpsWorksDeploymentStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.events.StackEvents.OpsWorksDeploymentStateChange",
		nil, // no parameters
		s,
	)
}

