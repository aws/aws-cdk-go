package interfacesawspcs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspcs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputeNodeGroup.
// Experimental.
type IComputeNodeGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ComputeNodeGroup resource.
	// Experimental.
	ComputeNodeGroupRef() *ComputeNodeGroupReference
}

// The jsii proxy for IComputeNodeGroupRef
type jsiiProxy_IComputeNodeGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IComputeNodeGroupRef) ComputeNodeGroupRef() *ComputeNodeGroupReference {
	var returns *ComputeNodeGroupReference
	_jsii_.Get(
		j,
		"computeNodeGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeNodeGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeNodeGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

