package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverRuleAssociation.
// Experimental.
type IResolverRuleAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResolverRuleAssociation resource.
	// Experimental.
	ResolverRuleAssociationRef() *ResolverRuleAssociationReference
}

// The jsii proxy for IResolverRuleAssociationRef
type jsiiProxy_IResolverRuleAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IResolverRuleAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolverRuleAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

