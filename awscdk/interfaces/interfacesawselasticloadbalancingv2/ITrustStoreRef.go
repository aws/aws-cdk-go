package interfacesawselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustStore.
// Experimental.
type ITrustStoreRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TrustStore resource.
	// Experimental.
	TrustStoreRef() *TrustStoreReference
}

// The jsii proxy for ITrustStoreRef
type jsiiProxy_ITrustStoreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITrustStoreRef) TrustStoreRef() *TrustStoreReference {
	var returns *TrustStoreReference
	_jsii_.Get(
		j,
		"trustStoreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStoreRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStoreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

