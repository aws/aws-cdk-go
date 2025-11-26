package previewawsecsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ecs@ECSTaskStateChange event types for Cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSTaskStateChange := #error#.NewECSTaskStateChange()
//
// Experimental.
type ClusterEvents_ECSTaskStateChange interface {
}

// The jsii proxy struct for ClusterEvents_ECSTaskStateChange
type jsiiProxy_ClusterEvents_ECSTaskStateChange struct {
	_ byte // padding
}

// Experimental.
func NewClusterEvents_ECSTaskStateChange() ClusterEvents_ECSTaskStateChange {
	_init_.Initialize()

	j := jsiiProxy_ClusterEvents_ECSTaskStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewClusterEvents_ECSTaskStateChange_Override(c ClusterEvents_ECSTaskStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSTaskStateChange",
		nil, // no parameters
		c,
	)
}

