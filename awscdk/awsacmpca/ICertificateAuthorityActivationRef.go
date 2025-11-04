package awsacmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CertificateAuthorityActivation.
// Experimental.
type ICertificateAuthorityActivationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CertificateAuthorityActivation resource.
	// Experimental.
	CertificateAuthorityActivationRef() *CertificateAuthorityActivationReference
}

// The jsii proxy for ICertificateAuthorityActivationRef
type jsiiProxy_ICertificateAuthorityActivationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICertificateAuthorityActivationRef) CertificateAuthorityActivationRef() *CertificateAuthorityActivationReference {
	var returns *CertificateAuthorityActivationReference
	_jsii_.Get(
		j,
		"certificateAuthorityActivationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAuthorityActivationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAuthorityActivationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

