package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverDNSSECConfig.
// Experimental.
type IResolverDNSSECConfigRef interface {
	constructs.IConstruct
	// A reference to a ResolverDNSSECConfig resource.
	// Experimental.
	ResolverDnssecConfigRef() *ResolverDNSSECConfigReference
}

// The jsii proxy for IResolverDNSSECConfigRef
type jsiiProxy_IResolverDNSSECConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResolverDNSSECConfigRef) ResolverDnssecConfigRef() *ResolverDNSSECConfigReference {
	var returns *ResolverDNSSECConfigReference
	_jsii_.Get(
		j,
		"resolverDnssecConfigRef",
		&returns,
	)
	return returns
}

