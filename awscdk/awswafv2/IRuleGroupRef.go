package awswafv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RuleGroup.
// Experimental.
type IRuleGroupRef interface {
	constructs.IConstruct
	// A reference to a RuleGroup resource.
	// Experimental.
	RuleGroupRef() *RuleGroupReference
}

// The jsii proxy for IRuleGroupRef
type jsiiProxy_IRuleGroupRef struct {
	internal.Type__constructsIConstruct
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

