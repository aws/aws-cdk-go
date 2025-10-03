package awsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsaps/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RuleGroupsNamespace.
// Experimental.
type IRuleGroupsNamespaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRuleGroupsNamespaceRef
type jsiiProxy_IRuleGroupsNamespaceRef struct {
	internal.Type__constructsIConstruct
}

