package interfacesawslex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslex/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BotAlias.
// Experimental.
type IBotAliasRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BotAlias resource.
	// Experimental.
	BotAliasRef() *BotAliasReference
}

// The jsii proxy for IBotAliasRef
type jsiiProxy_IBotAliasRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IBotAliasRef) BotAliasRef() *BotAliasReference {
	var returns *BotAliasReference
	_jsii_.Get(
		j,
		"botAliasRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBotAliasRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBotAliasRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

