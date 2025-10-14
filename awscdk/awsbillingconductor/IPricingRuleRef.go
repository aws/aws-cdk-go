package awsbillingconductor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PricingRule.
// Experimental.
type IPricingRuleRef interface {
	constructs.IConstruct
	// A reference to a PricingRule resource.
	// Experimental.
	PricingRuleRef() *PricingRuleReference
}

// The jsii proxy for IPricingRuleRef
type jsiiProxy_IPricingRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPricingRuleRef) PricingRuleRef() *PricingRuleReference {
	var returns *PricingRuleReference
	_jsii_.Get(
		j,
		"pricingRuleRef",
		&returns,
	)
	return returns
}

