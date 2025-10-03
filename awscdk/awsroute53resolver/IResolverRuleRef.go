package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverRule.
// Experimental.
type IResolverRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResolverRuleRef
type jsiiProxy_IResolverRuleRef struct {
	internal.Type__constructsIConstruct
}

