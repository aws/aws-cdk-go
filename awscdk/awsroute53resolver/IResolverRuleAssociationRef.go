package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverRuleAssociation.
// Experimental.
type IResolverRuleAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResolverRuleAssociationRef
type jsiiProxy_IResolverRuleAssociationRef struct {
	internal.Type__constructsIConstruct
}

