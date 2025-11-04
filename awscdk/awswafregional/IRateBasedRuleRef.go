package awswafregional

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RateBasedRule.
// Experimental.
type IRateBasedRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RateBasedRule resource.
	// Experimental.
	RateBasedRuleRef() *RateBasedRuleReference
}

// The jsii proxy for IRateBasedRuleRef
type jsiiProxy_IRateBasedRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IRateBasedRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRateBasedRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

