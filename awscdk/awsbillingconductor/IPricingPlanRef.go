package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PricingPlan.
// Experimental.
type IPricingPlanRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPricingPlanRef
type jsiiProxy_IPricingPlanRef struct {
	internal.Type__constructsIConstruct
}

