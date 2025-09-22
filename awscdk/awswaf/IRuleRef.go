package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswaf/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Rule.
// Experimental.
type IRuleRef interface {
	constructs.IConstruct
	// A reference to a Rule resource.
	// Experimental.
	RuleRef() *RuleReference
}

// The jsii proxy for IRuleRef
type jsiiProxy_IRuleRef struct {
	internal.Type__constructsIConstruct
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

