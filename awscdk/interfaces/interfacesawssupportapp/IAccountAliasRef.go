package interfacesawssupportapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssupportapp/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccountAlias.
// Experimental.
type IAccountAliasRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AccountAlias resource.
	// Experimental.
	AccountAliasRef() *AccountAliasReference
}

// The jsii proxy for IAccountAliasRef
type jsiiProxy_IAccountAliasRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAccountAliasRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAccountAliasRef) AccountAliasRef() *AccountAliasReference {
	var returns *AccountAliasReference
	_jsii_.Get(
		j,
		"accountAliasRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccountAliasRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccountAliasRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

