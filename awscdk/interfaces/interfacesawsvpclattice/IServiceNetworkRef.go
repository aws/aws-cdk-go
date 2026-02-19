package interfacesawsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetwork.
// Experimental.
type IServiceNetworkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServiceNetwork resource.
	// Experimental.
	ServiceNetworkRef() *ServiceNetworkReference
}

// The jsii proxy for IServiceNetworkRef
type jsiiProxy_IServiceNetworkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IServiceNetworkRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IServiceNetworkRef) ServiceNetworkRef() *ServiceNetworkReference {
	var returns *ServiceNetworkReference
	_jsii_.Get(
		j,
		"serviceNetworkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceNetworkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceNetworkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

