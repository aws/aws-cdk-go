package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

// Properties to define an application load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var certificate certificate
//   var hostedZone hostedZone
//
//   applicationLoadBalancerProps := &applicationLoadBalancerProps{
//   	listeners: []applicationListenerProps{
//   		&applicationListenerProps{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			certificate: certificate,
//   			port: jsii.Number(123),
//   			protocol: awscdk.Aws_elasticloadbalancingv2.applicationProtocol_HTTP,
//   			sslPolicy: awscdk.*Aws_elasticloadbalancingv2.sslPolicy_RECOMMENDED,
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	domainName: jsii.String("domainName"),
//   	domainZone: hostedZone,
//   	publicLoadBalancer: jsii.Boolean(false),
//   }
//
// Experimental.
type ApplicationLoadBalancerProps struct {
	// Listeners (at least one listener) attached to this load balancer.
	// Experimental.
	Listeners *[]*ApplicationListenerProps `field:"required" json:"listeners" yaml:"listeners"`
	// Name of the load balancer.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The domain name for the service, e.g. "api.example.com.".
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Experimental.
	DomainZone awsroute53.IHostedZone `field:"optional" json:"domainZone" yaml:"domainZone"`
	// Determines whether the Load Balancer will be internet-facing.
	// Experimental.
	PublicLoadBalancer *bool `field:"optional" json:"publicLoadBalancer" yaml:"publicLoadBalancer"`
}

