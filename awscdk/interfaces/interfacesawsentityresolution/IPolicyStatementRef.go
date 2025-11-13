package interfacesawsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyStatement.
// Experimental.
type IPolicyStatementRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PolicyStatement resource.
	// Experimental.
	PolicyStatementRef() *PolicyStatementReference
}

// The jsii proxy for IPolicyStatementRef
type jsiiProxy_IPolicyStatementRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPolicyStatementRef) PolicyStatementRef() *PolicyStatementReference {
	var returns *PolicyStatementReference
	_jsii_.Get(
		j,
		"policyStatementRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyStatementRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyStatementRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

