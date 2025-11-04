package awsopsworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Volume.
// Experimental.
type IVolumeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Volume resource.
	// Experimental.
	VolumeRef() *VolumeReference
}

// The jsii proxy for IVolumeRef
type jsiiProxy_IVolumeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVolumeRef) VolumeRef() *VolumeReference {
	var returns *VolumeReference
	_jsii_.Get(
		j,
		"volumeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVolumeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVolumeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

