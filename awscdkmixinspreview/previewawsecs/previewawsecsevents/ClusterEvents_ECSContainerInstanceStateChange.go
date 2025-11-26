package previewawsecsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ecs@ECSContainerInstanceStateChange event types for Cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSContainerInstanceStateChange := #error#.NewECSContainerInstanceStateChange()
//
// Experimental.
type ClusterEvents_ECSContainerInstanceStateChange interface {
}

// The jsii proxy struct for ClusterEvents_ECSContainerInstanceStateChange
type jsiiProxy_ClusterEvents_ECSContainerInstanceStateChange struct {
	_ byte // padding
}

// Experimental.
func NewClusterEvents_ECSContainerInstanceStateChange() ClusterEvents_ECSContainerInstanceStateChange {
	_init_.Initialize()

	j := jsiiProxy_ClusterEvents_ECSContainerInstanceStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSContainerInstanceStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewClusterEvents_ECSContainerInstanceStateChange_Override(c ClusterEvents_ECSContainerInstanceStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.ECSContainerInstanceStateChange",
		nil, // no parameters
		c,
	)
}

