package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Properties to define an application load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
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
//   			sslPolicy: awscdk.*Aws_elasticloadbalancingv2.sslPolicy_RECOMMENDED_TLS,
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	domainName: jsii.String("domainName"),
//   	domainZone: hostedZone,
//   	idleTimeout: cdk.duration.minutes(jsii.Number(30)),
//   	publicLoadBalancer: jsii.Boolean(false),
//   }
//
type ApplicationLoadBalancerProps struct {
	// Listeners (at least one listener) attached to this load balancer.
	Listeners *[]*ApplicationListenerProps `field:"required" json:"listeners" yaml:"listeners"`
	// Name of the load balancer.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The domain name for the service, e.g. "api.example.com.".
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	DomainZone awsroute53.IHostedZone `field:"optional" json:"domainZone" yaml:"domainZone"`
	// The load balancer idle timeout, in seconds.
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Determines whether the Load Balancer will be internet-facing.
	PublicLoadBalancer *bool `field:"optional" json:"publicLoadBalancer" yaml:"publicLoadBalancer"`
}

