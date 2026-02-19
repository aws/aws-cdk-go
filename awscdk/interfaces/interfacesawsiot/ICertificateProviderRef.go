package interfacesawsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CertificateProvider.
// Experimental.
type ICertificateProviderRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CertificateProvider resource.
	// Experimental.
	CertificateProviderRef() *CertificateProviderReference
}

// The jsii proxy for ICertificateProviderRef
type jsiiProxy_ICertificateProviderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICertificateProviderRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICertificateProviderRef) CertificateProviderRef() *CertificateProviderReference {
	var returns *CertificateProviderReference
	_jsii_.Get(
		j,
		"certificateProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateProviderRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateProviderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

