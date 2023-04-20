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
//   applicationLoadBalancerProps := &ApplicationLoadBalancerProps{
//   	Listeners: []applicationListenerProps{
//   		&applicationListenerProps{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Certificate: certificate,
//   			Port: jsii.Number(123),
//   			Protocol: awscdk.Aws_elasticloadbalancingv2.ApplicationProtocol_HTTP,
//   			SslPolicy: awscdk.*Aws_elasticloadbalancingv2.SslPolicy_RECOMMENDED,
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DomainName: jsii.String("domainName"),
//   	DomainZone: hostedZone,
//   	PublicLoadBalancer: jsii.Boolean(false),
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

