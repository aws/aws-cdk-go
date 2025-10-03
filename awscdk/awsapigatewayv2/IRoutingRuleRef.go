package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RoutingRule.
// Experimental.
type IRoutingRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRoutingRuleRef
type jsiiProxy_IRoutingRuleRef struct {
	internal.Type__constructsIConstruct
}

