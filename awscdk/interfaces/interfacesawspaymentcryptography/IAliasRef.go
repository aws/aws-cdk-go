package interfacesawspaymentcryptography

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspaymentcryptography/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Alias.
// Experimental.
type IAliasRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Alias resource.
	// Experimental.
	AliasRef() *AliasReference
}

// The jsii proxy for IAliasRef
type jsiiProxy_IAliasRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAliasRef) AliasRef() *AliasReference {
	var returns *AliasReference
	_jsii_.Get(
		j,
		"aliasRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAliasRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAliasRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

