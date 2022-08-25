package awsservicediscovery


// A complex type that contains information about the Amazon Route 53 DNS records that you want AWS Cloud Map to create when you register an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dnsConfigProperty := &dnsConfigProperty{
//   	dnsRecords: []interface{}{
//   		&dnsRecordProperty{
//   			ttl: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	namespaceId: jsii.String("namespaceId"),
//   	routingPolicy: jsii.String("routingPolicy"),
//   }
//
type CfnService_DnsConfigProperty struct {
	// An array that contains one `DnsRecord` object for each Route 53 DNS record that you want AWS Cloud Map to create when you register an instance.
	DnsRecords interface{} `field:"required" json:"dnsRecords" yaml:"dnsRecords"`
	// The ID of the namespace to use for DNS configuration.
	//
	// > You must specify a value for `NamespaceId` either for `DnsConfig` or for the [service properties](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-service.html) . Don't specify a value in both places.
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// The routing policy that you want to apply to all Route 53 DNS records that AWS Cloud Map creates when you register an instance and specify this service.
	//
	// > If you want to use this service to register instances that create alias records, specify `WEIGHTED` for the routing policy.
	//
	// You can specify the following values:
	//
	// - **MULTIVALUE** - If you define a health check for the service and the health check is healthy, Route 53 returns the applicable value for up to eight instances.
	//
	// For example, suppose that the service includes configurations for one `A` record and a health check. You use the service to register 10 instances. Route 53 responds to DNS queries with IP addresses for up to eight healthy instances. If fewer than eight instances are healthy, Route 53 responds to every DNS query with the IP addresses for all of the healthy instances.
	//
	// If you don't define a health check for the service, Route 53 assumes that all instances are healthy and returns the values for up to eight instances.
	//
	// For more information about the multivalue routing policy, see [Multivalue Answer Routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-multivalue) in the *Route 53 Developer Guide* .
	// - **WEIGHTED** - Route 53 returns the applicable value from one randomly selected instance from among the instances that you registered using the same service. Currently, all records have the same weight, so you can't route more or less traffic to any instances.
	//
	// For example, suppose that the service includes configurations for one `A` record and a health check. You use the service to register 10 instances. Route 53 responds to DNS queries with the IP address for one randomly selected instance from among the healthy instances. If no instances are healthy, Route 53 responds to DNS queries as if all of the instances were healthy.
	//
	// If you don't define a health check for the service, Route 53 assumes that all instances are healthy and returns the applicable value for one randomly selected instance.
	//
	// For more information about the weighted routing policy, see [Weighted Routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-weighted) in the *Route 53 Developer Guide* .
	RoutingPolicy *string `field:"optional" json:"routingPolicy" yaml:"routingPolicy"`
}

