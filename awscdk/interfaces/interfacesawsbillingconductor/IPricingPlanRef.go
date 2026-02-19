package interfacesawsbillingconductor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PricingPlan.
// Experimental.
type IPricingPlanRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PricingPlan resource.
	// Experimental.
	PricingPlanRef() *PricingPlanReference
}

// The jsii proxy for IPricingPlanRef
type jsiiProxy_IPricingPlanRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPricingPlanRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IPricingPlanRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPricingPlanRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

