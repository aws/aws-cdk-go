package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CertificateProvider.
// Experimental.
type ICertificateProviderRef interface {
	constructs.IConstruct
	// A reference to a CertificateProvider resource.
	// Experimental.
	CertificateProviderRef() *CertificateProviderReference
}

// The jsii proxy for ICertificateProviderRef
type jsiiProxy_ICertificateProviderRef struct {
	internal.Type__constructsIConstruct
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

