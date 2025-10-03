package awstransfer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Certificate.
// Experimental.
type ICertificateRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICertificateRef
type jsiiProxy_ICertificateRef struct {
	internal.Type__constructsIConstruct
}

