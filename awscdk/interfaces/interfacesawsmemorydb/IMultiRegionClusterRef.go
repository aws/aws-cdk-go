package interfacesawsmemorydb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmemorydb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MultiRegionCluster.
// Experimental.
type IMultiRegionClusterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MultiRegionCluster resource.
	// Experimental.
	MultiRegionClusterRef() *MultiRegionClusterReference
}

// The jsii proxy for IMultiRegionClusterRef
type jsiiProxy_IMultiRegionClusterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMultiRegionClusterRef) MultiRegionClusterRef() *MultiRegionClusterReference {
	var returns *MultiRegionClusterReference
	_jsii_.Get(
		j,
		"multiRegionClusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiRegionClusterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiRegionClusterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

