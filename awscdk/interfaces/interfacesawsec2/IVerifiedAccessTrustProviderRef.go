package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VerifiedAccessTrustProvider.
// Experimental.
type IVerifiedAccessTrustProviderRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VerifiedAccessTrustProvider resource.
	// Experimental.
	VerifiedAccessTrustProviderRef() *VerifiedAccessTrustProviderReference
}

// The jsii proxy for IVerifiedAccessTrustProviderRef
type jsiiProxy_IVerifiedAccessTrustProviderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVerifiedAccessTrustProviderRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVerifiedAccessTrustProviderRef) VerifiedAccessTrustProviderRef() *VerifiedAccessTrustProviderReference {
	var returns *VerifiedAccessTrustProviderReference
	_jsii_.Get(
		j,
		"verifiedAccessTrustProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedAccessTrustProviderRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVerifiedAccessTrustProviderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

