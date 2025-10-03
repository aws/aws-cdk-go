package awswafv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RuleGroup.
// Experimental.
type IRuleGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRuleGroupRef
type jsiiProxy_IRuleGroupRef struct {
	internal.Type__constructsIConstruct
}

