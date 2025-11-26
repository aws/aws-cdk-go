package previewawsecsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ecs@ECSServiceAction event types for Cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSServiceAction := #error#.NewECSServiceAction()
//
// Experimental.
type ClusterEvents_ECSServiceAction interface {
}

// The jsii proxy struct for ClusterEvents_ECSServiceAction
type jsiiProxy_ClusterEvents_ECSServiceAction struct {
	_ byte // padding
}

// Experimental.
func NewClusterEvents_ECSServiceAction() ClusterEvents_ECSServiceAction {
	_init_.Initialize()

	j := jsiiProxy_ClusterEvents_ECSServiceAction{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSServiceAction",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewClusterEvents_ECSServiceAction_Override(c ClusterEvents_ECSServiceAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSServiceAction",
		nil, // no parameters
		c,
	)
}

