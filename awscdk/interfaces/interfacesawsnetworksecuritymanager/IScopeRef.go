package interfacesawsnetworksecuritymanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworksecuritymanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Scope.
// Experimental.
type IScopeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Scope resource.
	// Experimental.
	ScopeRef() *ScopeReference
}

// The jsii proxy for IScopeRef
type jsiiProxy_IScopeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IScopeRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IScopeRef) ScopeRef() *ScopeReference {
	var returns *ScopeReference
	_jsii_.Get(
		j,
		"scopeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScopeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScopeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

