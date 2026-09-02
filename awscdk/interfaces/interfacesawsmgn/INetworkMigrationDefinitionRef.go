package interfacesawsmgn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmgn/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkMigrationDefinition.
// Experimental.
type INetworkMigrationDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NetworkMigrationDefinition resource.
	// Experimental.
	NetworkMigrationDefinitionRef() *NetworkMigrationDefinitionReference
}

// The jsii proxy for INetworkMigrationDefinitionRef
type jsiiProxy_INetworkMigrationDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_INetworkMigrationDefinitionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INetworkMigrationDefinitionRef) NetworkMigrationDefinitionRef() *NetworkMigrationDefinitionReference {
	var returns *NetworkMigrationDefinitionReference
	_jsii_.Get(
		j,
		"networkMigrationDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkMigrationDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkMigrationDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

