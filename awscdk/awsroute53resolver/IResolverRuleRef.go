package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverRule.
// Experimental.
type IResolverRuleRef interface {
	constructs.IConstruct
	// A reference to a ResolverRule resource.
	// Experimental.
	ResolverRuleRef() *ResolverRuleReference
}

// The jsii proxy for IResolverRuleRef
type jsiiProxy_IResolverRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResolverRuleRef) ResolverRuleRef() *ResolverRuleReference {
	var returns *ResolverRuleReference
	_jsii_.Get(
		j,
		"resolverRuleRef",
		&returns,
	)
	return returns
}

