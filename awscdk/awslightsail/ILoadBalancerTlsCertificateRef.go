package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoadBalancerTlsCertificate.
// Experimental.
type ILoadBalancerTlsCertificateRef interface {
	constructs.IConstruct
	// A reference to a LoadBalancerTlsCertificate resource.
	// Experimental.
	LoadBalancerTlsCertificateRef() *LoadBalancerTlsCertificateReference
}

// The jsii proxy for ILoadBalancerTlsCertificateRef
type jsiiProxy_ILoadBalancerTlsCertificateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILoadBalancerTlsCertificateRef) LoadBalancerTlsCertificateRef() *LoadBalancerTlsCertificateReference {
	var returns *LoadBalancerTlsCertificateReference
	_jsii_.Get(
		j,
		"loadBalancerTlsCertificateRef",
		&returns,
	)
	return returns
}

