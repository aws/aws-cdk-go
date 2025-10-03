package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ListenerCertificate.
// Experimental.
type IListenerCertificateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IListenerCertificateRef
type jsiiProxy_IListenerCertificateRef struct {
	internal.Type__constructsIConstruct
}

