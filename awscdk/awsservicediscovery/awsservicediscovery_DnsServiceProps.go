package awsservicediscovery

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Service props needed to create a service in a given namespace.
//
// Used by createService() for PrivateDnsNamespace and
// PublicDnsNamespace.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &publicDnsNamespaceProps{
//   	name: jsii.String("foobar.com"),
//   })
//
//   service := namespace.createService(jsii.String("Service"), &dnsServiceProps{
//   	name: jsii.String("foo"),
//   	dnsRecordType: servicediscovery.dnsRecordType_A,
//   	dnsTtl: cdk.duration.seconds(jsii.Number(30)),
//   	healthCheck: &healthCheckConfig{
//   		type: servicediscovery.healthCheckType_HTTPS,
//   		resourcePath: jsii.String("/healthcheck"),
//   		failureThreshold: jsii.Number(2),
//   	},
//   })
//
//   service.registerIpInstance(jsii.String("IpInstance"), &ipInstanceBaseProps{
//   	ipv4: jsii.String("54.239.25.192"),
//   	port: jsii.Number(443),
//   })
//
//   app.synth()
//
// Experimental.
type DnsServiceProps struct {
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
}

