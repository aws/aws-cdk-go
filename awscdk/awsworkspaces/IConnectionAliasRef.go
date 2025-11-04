package awsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectionAlias.
// Experimental.
type IConnectionAliasRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ConnectionAlias resource.
	// Experimental.
	ConnectionAliasRef() *ConnectionAliasReference
}

// The jsii proxy for IConnectionAliasRef
type jsiiProxy_IConnectionAliasRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IConnectionAliasRef) ConnectionAliasRef() *ConnectionAliasReference {
	var returns *ConnectionAliasReference
	_jsii_.Get(
		j,
		"connectionAliasRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionAliasRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionAliasRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

