package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ListenerCertificate.
// Experimental.
type IListenerCertificateRef interface {
	constructs.IConstruct
	// A reference to a ListenerCertificate resource.
	// Experimental.
	ListenerCertificateRef() *ListenerCertificateReference
}

// The jsii proxy for IListenerCertificateRef
type jsiiProxy_IListenerCertificateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IListenerCertificateRef) ListenerCertificateRef() *ListenerCertificateReference {
	var returns *ListenerCertificateReference
	_jsii_.Get(
		j,
		"listenerCertificateRef",
		&returns,
	)
	return returns
}

