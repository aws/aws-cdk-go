package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PricingRule.
// Experimental.
type IPricingRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPricingRuleRef
type jsiiProxy_IPricingRuleRef struct {
	internal.Type__constructsIConstruct
}

