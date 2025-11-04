package awssupportapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssupportapp/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccountAlias.
// Experimental.
type IAccountAliasRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AccountAlias resource.
	// Experimental.
	AccountAliasRef() *AccountAliasReference
}

// The jsii proxy for IAccountAliasRef
type jsiiProxy_IAccountAliasRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAccountAliasRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

