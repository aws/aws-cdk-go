package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var namespace iNamespace
//
//   serviceProps := &serviceProps{
//   	namespace: namespace,
//
//   	// the properties below are optional
//   	customHealthCheck: &healthCheckCustomConfig{
//   		failureThreshold: jsii.Number(123),
//   	},
//   	description: jsii.String("description"),
//   	discoveryType: awscdk.Aws_servicediscovery.discoveryType_API,
//   	dnsRecordType: awscdk.*Aws_servicediscovery.dnsRecordType_A,
//   	dnsTtl: cdk.duration.minutes(jsii.Number(30)),
//   	healthCheck: &healthCheckConfig{
//   		failureThreshold: jsii.Number(123),
//   		resourcePath: jsii.String("resourcePath"),
//   		type: awscdk.*Aws_servicediscovery.healthCheckType_HTTP,
//   	},
//   	loadBalancer: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	routingPolicy: awscdk.*Aws_servicediscovery.routingPolicy_WEIGHTED,
//   }
//
type ServiceProps struct {
	// Structure containing failure threshold for a custom health checker.
	//
	// Only one of healthCheckConfig or healthCheckCustomConfig can be specified.
	// See: https://docs.aws.amazon.com/cloud-map/latest/api/API_HealthCheckCustomConfig.html
	CustomHealthCheck *HealthCheckCustomConfig `field:"optional" json:"customHealthCheck" yaml:"customHealthCheck"`
	// A description of the service.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Settings for an optional health check.
	//
	// If you specify health check settings, AWS Cloud Map associates the health
	// check with the records that you specify in DnsConfig. Only one of healthCheckConfig or healthCheckCustomConfig can
	// be specified. Not valid for PrivateDnsNamespaces. If you use healthCheck, you can only register IP instances to
	// this service.
	HealthCheck *HealthCheckConfig `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// A name for the Service.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Controls how instances within this service can be discovered.
	DiscoveryType DiscoveryType `field:"optional" json:"discoveryType" yaml:"discoveryType"`
	// The DNS type of the record that you want AWS Cloud Map to create.
	//
	// Supported record types
	// include A, AAAA, A and AAAA (A_AAAA), CNAME, and SRV.
	DnsRecordType DnsRecordType `field:"optional" json:"dnsRecordType" yaml:"dnsRecordType"`
	// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this record.
	DnsTtl awscdk.Duration `field:"optional" json:"dnsTtl" yaml:"dnsTtl"`
	// Whether or not this service will have an Elastic LoadBalancer registered to it as an AliasTargetInstance.
	//
	// Setting this to `true` correctly configures the `routingPolicy`
	// and performs some additional validation.
	LoadBalancer *bool `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// The routing policy that you want to apply to all DNS records that AWS Cloud Map creates when you register an instance and specify this service.
	RoutingPolicy RoutingPolicy `field:"optional" json:"routingPolicy" yaml:"routingPolicy"`
	// The namespace that you want to use for DNS configuration.
	Namespace INamespace `field:"required" json:"namespace" yaml:"namespace"`
}

