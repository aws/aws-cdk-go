package previewawsecsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.ecs@AWSAPICallViaCloudTrail event types for Cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrail := #error#.NewAWSAPICallViaCloudTrail()
//
// Experimental.
type ClusterEvents_AWSAPICallViaCloudTrail interface {
}

// The jsii proxy struct for ClusterEvents_AWSAPICallViaCloudTrail
type jsiiProxy_ClusterEvents_AWSAPICallViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewClusterEvents_AWSAPICallViaCloudTrail() ClusterEvents_AWSAPICallViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_ClusterEvents_AWSAPICallViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewClusterEvents_AWSAPICallViaCloudTrail_Override(c ClusterEvents_AWSAPICallViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.events.ClusterEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		c,
	)
}

