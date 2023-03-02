package awsacmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsacmpca/internal"
)

// Interface which all CertificateAuthority based class must implement.
// Experimental.
type ICertificateAuthority interface {
	awscdk.IResource
	// The Amazon Resource Name of the Certificate.
	// Experimental.
	CertificateAuthorityArn() *string
}

// The jsii proxy for ICertificateAuthority
type jsiiProxy_ICertificateAuthority struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ICertificateAuthority) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

