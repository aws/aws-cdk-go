package interfacesawselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustStoreRevocation.
// Experimental.
type ITrustStoreRevocationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TrustStoreRevocation resource.
	// Experimental.
	TrustStoreRevocationRef() *TrustStoreRevocationReference
}

// The jsii proxy for ITrustStoreRevocationRef
type jsiiProxy_ITrustStoreRevocationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITrustStoreRevocationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITrustStoreRevocationRef) TrustStoreRevocationRef() *TrustStoreRevocationReference {
	var returns *TrustStoreRevocationReference
	_jsii_.Get(
		j,
		"trustStoreRevocationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStoreRevocationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrustStoreRevocationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

