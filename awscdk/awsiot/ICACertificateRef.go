package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CACertificate.
// Experimental.
type ICACertificateRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICACertificateRef
type jsiiProxy_ICACertificateRef struct {
	internal.Type__constructsIConstruct
}

