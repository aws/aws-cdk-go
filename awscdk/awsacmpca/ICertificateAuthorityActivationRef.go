package awsacmpca

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CertificateAuthorityActivation.
// Experimental.
type ICertificateAuthorityActivationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICertificateAuthorityActivationRef
type jsiiProxy_ICertificateAuthorityActivationRef struct {
	internal.Type__constructsIConstruct
}

