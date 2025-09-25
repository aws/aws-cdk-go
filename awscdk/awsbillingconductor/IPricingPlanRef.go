package awsbillingconductor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PricingPlan.
// Experimental.
type IPricingPlanRef interface {
	constructs.IConstruct
	// A reference to a PricingPlan resource.
	// Experimental.
	PricingPlanRef() *PricingPlanReference
}

// The jsii proxy for IPricingPlanRef
type jsiiProxy_IPricingPlanRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPricingPlanRef) PricingPlanRef() *PricingPlanReference {
	var returns *PricingPlanReference
	_jsii_.Get(
		j,
		"pricingPlanRef",
		&returns,
	)
	return returns
}

