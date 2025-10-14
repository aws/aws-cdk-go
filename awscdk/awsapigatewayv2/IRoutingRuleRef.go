package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RoutingRule.
// Experimental.
type IRoutingRuleRef interface {
	constructs.IConstruct
	// A reference to a RoutingRule resource.
	// Experimental.
	RoutingRuleRef() *RoutingRuleReference
}

// The jsii proxy for IRoutingRuleRef
type jsiiProxy_IRoutingRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRoutingRuleRef) RoutingRuleRef() *RoutingRuleReference {
	var returns *RoutingRuleReference
	_jsii_.Get(
		j,
		"routingRuleRef",
		&returns,
	)
	return returns
}

