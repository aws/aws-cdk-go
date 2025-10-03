package awsacmpca

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CertificateAuthority.
// Experimental.
type ICertificateAuthorityRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICertificateAuthorityRef
type jsiiProxy_ICertificateAuthorityRef struct {
	internal.Type__constructsIConstruct
}

