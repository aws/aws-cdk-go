package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverQueryLoggingConfig.
// Experimental.
type IResolverQueryLoggingConfigRef interface {
	constructs.IConstruct
	// A reference to a ResolverQueryLoggingConfig resource.
	// Experimental.
	ResolverQueryLoggingConfigRef() *ResolverQueryLoggingConfigReference
}

// The jsii proxy for IResolverQueryLoggingConfigRef
type jsiiProxy_IResolverQueryLoggingConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResolverQueryLoggingConfigRef) ResolverQueryLoggingConfigRef() *ResolverQueryLoggingConfigReference {
	var returns *ResolverQueryLoggingConfigReference
	_jsii_.Get(
		j,
		"resolverQueryLoggingConfigRef",
		&returns,
	)
	return returns
}

