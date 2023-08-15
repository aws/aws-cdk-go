package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Service props needed to create a service in a given namespace.
//
// Used by createService() for PrivateDnsNamespace and
// PublicDnsNamespace.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &PublicDnsNamespaceProps{
//   	Name: jsii.String("foobar.com"),
//   })
//
//   service := namespace.CreateService(jsii.String("Service"), &DnsServiceProps{
//   	Name: jsii.String("foo"),
//   	DnsRecordType: servicediscovery.DnsRecordType_A,
//   	DnsTtl: cdk.Duration_Seconds(jsii.Number(30)),
//   	HealthCheck: &HealthCheckConfig{
//   		Type: servicediscovery.HealthCheckType_HTTPS,
//   		ResourcePath: jsii.String("/healthcheck"),
//   		FailureThreshold: jsii.Number(2),
//   	},
//   })
//
//   service.RegisterIpInstance(jsii.String("IpInstance"), &IpInstanceBaseProps{
//   	Ipv4: jsii.String("54.239.25.192"),
//   	Port: jsii.Number(443),
//   })
//
//   app.Synth()
//
type DnsServiceProps struct {
	// Structure containing failure threshold for a custom health checker.
	//
	// Only one of healthCheckConfig or healthCheckCustomConfig can be specified.
	// See: https://docs.aws.amazon.com/cloud-map/latest/api/API_HealthCheckCustomConfig.html
	// Default: none.
	//
	CustomHealthCheck *HealthCheckCustomConfig `field:"optional" json:"customHealthCheck" yaml:"customHealthCheck"`
	// A description of the service.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Settings for an optional health check.
	//
	// If you specify health check settings, AWS Cloud Map associates the health
	// check with the records that you specify in DnsConfig. Only one of healthCheckConfig or healthCheckCustomConfig can
	// be specified. Not valid for PrivateDnsNamespaces. If you use healthCheck, you can only register IP instances to
	// this service.
	// Default: none.
	//
	HealthCheck *HealthCheckConfig `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// A name for the Service.
	// Default: CloudFormation-generated name.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Controls how instances within this service can be discovered.
	// Default: DNS_AND_API.
	//
	DiscoveryType DiscoveryType `field:"optional" json:"discoveryType" yaml:"discoveryType"`
	// The DNS type of the record that you want AWS Cloud Map to create.
	//
	// Supported record types
	// include A, AAAA, A and AAAA (A_AAAA), CNAME, and SRV.
	// Default: A.
	//
	DnsRecordType DnsRecordType `field:"optional" json:"dnsRecordType" yaml:"dnsRecordType"`
	// The amount of time, in seconds, that you want DNS resolvers to cache the settings for this record.
	// Default: Duration.minutes(1)
	//
	DnsTtl awscdk.Duration `field:"optional" json:"dnsTtl" yaml:"dnsTtl"`
	// Whether or not this service will have an Elastic LoadBalancer registered to it as an AliasTargetInstance.
	//
	// Setting this to `true` correctly configures the `routingPolicy`
	// and performs some additional validation.
	// Default: false.
	//
	LoadBalancer *bool `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// The routing policy that you want to apply to all DNS records that AWS Cloud Map creates when you register an instance and specify this service.
	// Default: WEIGHTED for CNAME records and when loadBalancer is true, MULTIVALUE otherwise.
	//
	RoutingPolicy RoutingPolicy `field:"optional" json:"routingPolicy" yaml:"routingPolicy"`
}

