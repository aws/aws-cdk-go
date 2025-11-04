package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverQueryLoggingConfigAssociation.
// Experimental.
type IResolverQueryLoggingConfigAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResolverQueryLoggingConfigAssociation resource.
	// Experimental.
	ResolverQueryLoggingConfigAssociationRef() *ResolverQueryLoggingConfigAssociationReference
}

// The jsii proxy for IResolverQueryLoggingConfigAssociationRef
type jsiiProxy_IResolverQueryLoggingConfigAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResolverQueryLoggingConfigAssociationRef) ResolverQueryLoggingConfigAssociationRef() *ResolverQueryLoggingConfigAssociationReference {
	var returns *ResolverQueryLoggingConfigAssociationReference
	_jsii_.Get(
		j,
		"resolverQueryLoggingConfigAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolverQueryLoggingConfigAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolverQueryLoggingConfigAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

