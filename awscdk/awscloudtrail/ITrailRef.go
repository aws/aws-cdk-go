package awscloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Trail.
// Experimental.
type ITrailRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Trail resource.
	// Experimental.
	TrailRef() *TrailReference
}

// The jsii proxy for ITrailRef
type jsiiProxy_ITrailRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITrailRef) TrailRef() *TrailReference {
	var returns *TrailReference
	_jsii_.Get(
		j,
		"trailRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrailRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrailRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

