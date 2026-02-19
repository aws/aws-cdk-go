package interfacesawsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverDNSSECConfig.
// Experimental.
type IResolverDNSSECConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ResolverDNSSECConfig resource.
	// Experimental.
	ResolverDnssecConfigRef() *ResolverDNSSECConfigReference
}

// The jsii proxy for IResolverDNSSECConfigRef
type jsiiProxy_IResolverDNSSECConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IResolverDNSSECConfigRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IResolverDNSSECConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolverDNSSECConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

