package interfacesawsbackupgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbackupgateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Hypervisor.
// Experimental.
type IHypervisorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Hypervisor resource.
	// Experimental.
	HypervisorRef() *HypervisorReference
}

// The jsii proxy for IHypervisorRef
type jsiiProxy_IHypervisorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IHypervisorRef) HypervisorRef() *HypervisorReference {
	var returns *HypervisorReference
	_jsii_.Get(
		j,
		"hypervisorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHypervisorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHypervisorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

