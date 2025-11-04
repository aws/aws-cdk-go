package awsservicediscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublicDnsNamespace.
// Experimental.
type IPublicDnsNamespaceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PublicDnsNamespace resource.
	// Experimental.
	PublicDnsNamespaceRef() *PublicDnsNamespaceReference
}

// The jsii proxy for IPublicDnsNamespaceRef
type jsiiProxy_IPublicDnsNamespaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPublicDnsNamespaceRef) PublicDnsNamespaceRef() *PublicDnsNamespaceReference {
	var returns *PublicDnsNamespaceReference
	_jsii_.Get(
		j,
		"publicDnsNamespaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicDnsNamespaceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicDnsNamespaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

