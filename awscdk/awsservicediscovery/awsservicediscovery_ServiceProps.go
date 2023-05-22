package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var namespace iNamespace
//
//   serviceProps := &ServiceProps{
//   	Namespace: namespace,
//
//   	// the properties below are optional
//   	CustomHealthCheck: &HealthCheckCustomConfig{
//   		FailureThreshold: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	DnsRecordType: awscdk.Aws_servicediscovery.DnsRecordType_A,
//   	DnsTtl: duration,
//   	HealthCheck: &HealthCheckConfig{
//   		FailureThreshold: jsii.Number(123),
//   		ResourcePath: jsii.String("resourcePath"),
//   		Type: awscdk.*Aws_servicediscovery.HealthCheckType_HTTP,
//   	},
//   	LoadBalancer: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	RoutingPolicy: awscdk.*Aws_servicediscovery.RoutingPolicy_WEIGHTED,
//   }
//
// Experimental.
type ServiceProps struct {
	// Structure containing failure threshold for a custom health checker.
	//
	// Only one of healthCheckConfig or healthCheckCustomConfig can be specified.
	// See: https://docs.aws.amazon.com/cloud-map/latest/api/API_HealthCheckCustomConfig.html
	// Experimental.
	CustomHealthCheck *HealthCheckCustomConfig `field:"optional" json:"customHealthCheck" yaml:"customHealthCheck"`
	// A description of the service.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Settings for an optional health check.
	//
	// If you specify health check settings, AWS Cloud Map associates the health
	// check with the records that you specify in DnsConfig. Only one of healthCheckConfig or healthCheckCustomConfig can
	// be specified. Not valid for PrivateDnsNamespaces. If you use healthCheck, you can only register IP instances to
	// this service.
	// Experimental.
	HealthCheck *HealthCheckConfig `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// A name for the Service.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The DNS type of the record that you want AWS Cloud Map to create.
	//
	// Supported record types
	// include A, AAAA, A and AAAA (A_AAAA), CNAME, and SRV.
	// Experimental.
	DnsRecordType DnsRecordType `field:"optional" json:"dnsRecordType" yaml:"dnsRecordType"`
	// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this record.
	// Experimental.
	DnsTtl awscdk.Duration `field:"optional" json:"dnsTtl" yaml:"dnsTtl"`
	// Whether or not this service will have an Elastic LoadBalancer registered to it as an AliasTargetInstance.
	//
	// Setting this to `true` correctly configures the `routingPolicy`
	// and performs some additional validation.
	// Experimental.
	LoadBalancer *bool `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// The routing policy that you want to apply to all DNS records that AWS Cloud Map creates when you register an instance and specify this service.
	// Experimental.
	RoutingPolicy RoutingPolicy `field:"optional" json:"routingPolicy" yaml:"routingPolicy"`
	// The namespace that you want to use for DNS configuration.
	// Experimental.
	Namespace INamespace `field:"required" json:"namespace" yaml:"namespace"`
}

