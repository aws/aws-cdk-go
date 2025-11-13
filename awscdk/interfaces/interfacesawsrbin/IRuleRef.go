package interfacesawsrbin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrbin/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Rule.
// Experimental.
type IRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Rule resource.
	// Experimental.
	RuleRef() *RuleReference
}

// The jsii proxy for IRuleRef
type jsiiProxy_IRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRuleRef) RuleRef() *RuleReference {
	var returns *RuleReference
	_jsii_.Get(
		j,
		"ruleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

