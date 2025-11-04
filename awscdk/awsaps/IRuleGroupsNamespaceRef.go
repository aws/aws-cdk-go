package awsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsaps/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RuleGroupsNamespace.
// Experimental.
type IRuleGroupsNamespaceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RuleGroupsNamespace resource.
	// Experimental.
	RuleGroupsNamespaceRef() *RuleGroupsNamespaceReference
}

// The jsii proxy for IRuleGroupsNamespaceRef
type jsiiProxy_IRuleGroupsNamespaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRuleGroupsNamespaceRef) RuleGroupsNamespaceRef() *RuleGroupsNamespaceReference {
	var returns *RuleGroupsNamespaceReference
	_jsii_.Get(
		j,
		"ruleGroupsNamespaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupsNamespaceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupsNamespaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

