package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RuleGroup.
// Experimental.
type IRuleGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RuleGroup resource.
	// Experimental.
	RuleGroupRef() *RuleGroupReference
}

// The jsii proxy for IRuleGroupRef
type jsiiProxy_IRuleGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRuleGroupRef) RuleGroupRef() *RuleGroupReference {
	var returns *RuleGroupReference
	_jsii_.Get(
		j,
		"ruleGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

