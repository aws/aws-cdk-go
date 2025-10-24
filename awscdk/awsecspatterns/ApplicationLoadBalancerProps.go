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
//   var certificate Certificate
//   var hostedZone HostedZone
//
//   applicationLoadBalancerProps := &ApplicationLoadBalancerProps{
//   	Listeners: []ApplicationListenerProps{
//   		&ApplicationListenerProps{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Certificate: certificate,
//   			Port: jsii.Number(123),
//   			Protocol: awscdk.Aws_elasticloadbalancingv2.ApplicationProtocol_HTTP,
//   			SslPolicy: awscdk.*Aws_elasticloadbalancingv2.SslPolicy_RECOMMENDED_TLS,
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DomainName: jsii.String("domainName"),
//   	DomainZone: hostedZone,
//   	IdleTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	PublicLoadBalancer: jsii.Boolean(false),
//   }
//
type ApplicationLoadBalancerProps struct {
	// Listeners (at least one listener) attached to this load balancer.
	Listeners *[]*ApplicationListenerProps `field:"required" json:"listeners" yaml:"listeners"`
	// Name of the load balancer.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The domain name for the service, e.g. "api.example.com.".
	// Default: - No domain name.
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	// Default: - No Route53 hosted domain zone.
	//
	DomainZone awsroute53.IHostedZone `field:"optional" json:"domainZone" yaml:"domainZone"`
	// The load balancer idle timeout, in seconds.
	//
	// Can be between 1 and 4000 seconds.
	// Default: - CloudFormation sets idle timeout to 60 seconds.
	//
	IdleTimeout awscdk.Duration `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Determines whether the Load Balancer will be internet-facing.
	// Default: true.
	//
	PublicLoadBalancer *bool `field:"optional" json:"publicLoadBalancer" yaml:"publicLoadBalancer"`
}

