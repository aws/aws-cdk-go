package awsbillingconductor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PricingRule.
// Experimental.
type IPricingRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PricingRule resource.
	// Experimental.
	PricingRuleRef() *PricingRuleReference
}

// The jsii proxy for IPricingRuleRef
type jsiiProxy_IPricingRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IPricingRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPricingRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

