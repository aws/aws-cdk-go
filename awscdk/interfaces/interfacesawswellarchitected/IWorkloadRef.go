package interfacesawswellarchitected

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswellarchitected/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workload.
// Experimental.
type IWorkloadRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Workload resource.
	// Experimental.
	WorkloadRef() *WorkloadReference
}

// The jsii proxy for IWorkloadRef
type jsiiProxy_IWorkloadRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWorkloadRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWorkloadRef) WorkloadRef() *WorkloadReference {
	var returns *WorkloadReference
	_jsii_.Get(
		j,
		"workloadRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

