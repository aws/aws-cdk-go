package interfacesawsverifiedpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsverifiedpermissions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyStoreAlias.
// Experimental.
type IPolicyStoreAliasRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PolicyStoreAlias resource.
	// Experimental.
	PolicyStoreAliasRef() *PolicyStoreAliasReference
}

// The jsii proxy for IPolicyStoreAliasRef
type jsiiProxy_IPolicyStoreAliasRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPolicyStoreAliasRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPolicyStoreAliasRef) PolicyStoreAliasRef() *PolicyStoreAliasReference {
	var returns *PolicyStoreAliasReference
	_jsii_.Get(
		j,
		"policyStoreAliasRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyStoreAliasRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyStoreAliasRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

