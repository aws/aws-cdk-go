package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Properties to define an network load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var hostedZone hostedZone
//
//   networkLoadBalancerProps := &NetworkLoadBalancerProps{
//   	Listeners: []networkListenerProps{
//   		&networkListenerProps{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Port: jsii.Number(123),
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
type NetworkLoadBalancerProps struct {
	// Listeners (at least one listener) attached to this load balancer.
	Listeners *[]*NetworkListenerProps `field:"required" json:"listeners" yaml:"listeners"`
	// Name of the load balancer.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The domain name for the service, e.g. "api.example.com.".
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone for the domain, e.g. "example.com.".
	DomainZone awsroute53.IHostedZone `field:"optional" json:"domainZone" yaml:"domainZone"`
	// Determines whether the Load Balancer will be internet-facing.
	PublicLoadBalancer *bool `field:"optional" json:"publicLoadBalancer" yaml:"publicLoadBalancer"`
}

