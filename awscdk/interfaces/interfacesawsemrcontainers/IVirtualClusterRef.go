package interfacesawsemrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsemrcontainers/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualCluster.
// Experimental.
type IVirtualClusterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VirtualCluster resource.
	// Experimental.
	VirtualClusterRef() *VirtualClusterReference
}

// The jsii proxy for IVirtualClusterRef
type jsiiProxy_IVirtualClusterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVirtualClusterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVirtualClusterRef) VirtualClusterRef() *VirtualClusterReference {
	var returns *VirtualClusterReference
	_jsii_.Get(
		j,
		"virtualClusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualClusterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualClusterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

