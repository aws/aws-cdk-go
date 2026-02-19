package interfacesawsverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsverifiedpermissions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyStore.
// Experimental.
type IPolicyStoreRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PolicyStore resource.
	// Experimental.
	PolicyStoreRef() *PolicyStoreReference
}

// The jsii proxy for IPolicyStoreRef
type jsiiProxy_IPolicyStoreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPolicyStoreRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPolicyStoreRef) PolicyStoreRef() *PolicyStoreReference {
	var returns *PolicyStoreReference
	_jsii_.Get(
		j,
		"policyStoreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyStoreRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyStoreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

