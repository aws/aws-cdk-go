package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverRuleAssociation.
// Experimental.
type IResolverRuleAssociationRef interface {
	constructs.IConstruct
	// A reference to a ResolverRuleAssociation resource.
	// Experimental.
	ResolverRuleAssociationRef() *ResolverRuleAssociationReference
}

// The jsii proxy for IResolverRuleAssociationRef
type jsiiProxy_IResolverRuleAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResolverRuleAssociationRef) ResolverRuleAssociationRef() *ResolverRuleAssociationReference {
	var returns *ResolverRuleAssociationReference
	_jsii_.Get(
		j,
		"resolverRuleAssociationRef",
		&returns,
	)
	return returns
}

