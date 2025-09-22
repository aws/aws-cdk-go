package awswafregional

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RateBasedRule.
// Experimental.
type IRateBasedRuleRef interface {
	constructs.IConstruct
	// A reference to a RateBasedRule resource.
	// Experimental.
	RateBasedRuleRef() *RateBasedRuleReference
}

// The jsii proxy for IRateBasedRuleRef
type jsiiProxy_IRateBasedRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRateBasedRuleRef) RateBasedRuleRef() *RateBasedRuleReference {
	var returns *RateBasedRuleReference
	_jsii_.Get(
		j,
		"rateBasedRuleRef",
		&returns,
	)
	return returns
}

