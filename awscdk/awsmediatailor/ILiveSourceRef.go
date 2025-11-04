package awsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LiveSource.
// Experimental.
type ILiveSourceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LiveSource resource.
	// Experimental.
	LiveSourceRef() *LiveSourceReference
}

// The jsii proxy for ILiveSourceRef
type jsiiProxy_ILiveSourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILiveSourceRef) LiveSourceRef() *LiveSourceReference {
	var returns *LiveSourceReference
	_jsii_.Get(
		j,
		"liveSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILiveSourceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILiveSourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

