package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LoadBalancerTlsCertificate.
// Experimental.
type ILoadBalancerTlsCertificateRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILoadBalancerTlsCertificateRef
type jsiiProxy_ILoadBalancerTlsCertificateRef struct {
	internal.Type__constructsIConstruct
}

