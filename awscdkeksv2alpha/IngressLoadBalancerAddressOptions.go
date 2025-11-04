package awscdkeksv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for fetching an IngressLoadBalancerAddress.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressLoadBalancerAddressOptions := &IngressLoadBalancerAddressOptions{
//   	Namespace: jsii.String("namespace"),
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type IngressLoadBalancerAddressOptions struct {
	// The namespace the service belongs to.
	// Default: 'default'.
	//
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Timeout for waiting on the load balancer address.
	// Default: Duration.minutes(5)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

