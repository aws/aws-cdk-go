package interfacesawsacmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CertificateAuthority.
// Experimental.
type ICertificateAuthorityRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CertificateAuthority resource.
	// Experimental.
	CertificateAuthorityRef() *CertificateAuthorityReference
}

// The jsii proxy for ICertificateAuthorityRef
type jsiiProxy_ICertificateAuthorityRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICertificateAuthorityRef) CertificateAuthorityRef() *CertificateAuthorityReference {
	var returns *CertificateAuthorityReference
	_jsii_.Get(
		j,
		"certificateAuthorityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAuthorityRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAuthorityRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

