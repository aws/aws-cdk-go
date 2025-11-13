package interfacesawsmediatailor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VodSource.
// Experimental.
type IVodSourceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VodSource resource.
	// Experimental.
	VodSourceRef() *VodSourceReference
}

// The jsii proxy for IVodSourceRef
type jsiiProxy_IVodSourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVodSourceRef) VodSourceRef() *VodSourceReference {
	var returns *VodSourceReference
	_jsii_.Get(
		j,
		"vodSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVodSourceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVodSourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

