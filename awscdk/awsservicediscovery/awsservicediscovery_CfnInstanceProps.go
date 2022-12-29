package awsservicediscovery


// Properties for defining a `CfnInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instanceAttributes interface{}
//
//   cfnInstanceProps := &cfnInstanceProps{
//   	instanceAttributes: instanceAttributes,
//   	serviceId: jsii.String("serviceId"),
//
//   	// the properties below are optional
//   	instanceId: jsii.String("instanceId"),
//   }
//
type CfnInstanceProps struct {
	// A string map that contains the following information for the service that you specify in `ServiceId` :.
	//
	// - The attributes that apply to the records that are defined in the service.
	// - For each attribute, the applicable value.
	//
	// Supported attribute keys include the following:
	//
	// - **AWS_ALIAS_DNS_NAME** - If you want AWS Cloud Map to create a Route 53 alias record that routes traffic to an Elastic Load Balancing load balancer, specify the DNS name that is associated with the load balancer. For information about how to get the DNS name, see [AliasTarget->DNSName](https://docs.aws.amazon.com/Route53/latest/APIReference/API_AliasTarget.html#Route53-Type-AliasTarget-DNSName) in the *Route 53 API Reference* .
	//
	// Note the following:
	//
	// - The configuration for the service that is specified by `ServiceId` must include settings for an `A` record, an `AAAA` record, or both.
	// - In the service that is specified by `ServiceId` , the value of `RoutingPolicy` must be `WEIGHTED` .
	// - If the service that is specified by `ServiceId` includes `HealthCheckConfig` settings, AWS Cloud Map will create the health check, but it won't associate the health check with the alias record.
	// - Auto naming currently doesn't support creating alias records that route traffic to AWS resources other than ELB load balancers.
	// - If you specify a value for `AWS_ALIAS_DNS_NAME` , don't specify values for any of the `AWS_INSTANCE` attributes.
	// - **AWS_EC2_INSTANCE_ID** - *HTTP namespaces only.* The Amazon EC2 instance ID for the instance. The `AWS_INSTANCE_IPV4` attribute contains the primary private IPv4 address. When creating resources with a type of [AWS::ServiceDiscovery::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html) , if the `AWS_EC2_INSTANCE_ID` attribute is specified, the only other attribute that can be specified is `AWS_INIT_HEALTH_STATUS` . After the resource has been created, the `AWS_INSTANCE_IPV4` attribute contains the primary private IPv4 address.
	// - **AWS_INIT_HEALTH_STATUS** - If the service configuration includes `HealthCheckCustomConfig` , when creating resources with a type of [AWS::ServiceDiscovery::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html) you can optionally use `AWS_INIT_HEALTH_STATUS` to specify the initial status of the custom health check, `HEALTHY` or `UNHEALTHY` . If you don't specify a value for `AWS_INIT_HEALTH_STATUS` , the initial status is `HEALTHY` . This attribute can only be used when creating resources and will not be seen on existing resources.
	// - **AWS_INSTANCE_CNAME** - If the service configuration includes a `CNAME` record, the domain name that you want Route 53 to return in response to DNS queries, for example, `example.com` .
	//
	// This value is required if the service specified by `ServiceId` includes settings for an `CNAME` record.
	// - **AWS_INSTANCE_IPV4** - If the service configuration includes an `A` record, the IPv4 address that you want Route 53 to return in response to DNS queries, for example, `192.0.2.44` .
	//
	// This value is required if the service specified by `ServiceId` includes settings for an `A` record. If the service includes settings for an `SRV` record, you must specify a value for `AWS_INSTANCE_IPV4` , `AWS_INSTANCE_IPV6` , or both.
	// - **AWS_INSTANCE_IPV6** - If the service configuration includes an `AAAA` record, the IPv6 address that you want Route 53 to return in response to DNS queries, for example, `2001:0db8:85a3:0000:0000:abcd:0001:2345` .
	//
	// This value is required if the service specified by `ServiceId` includes settings for an `AAAA` record. If the service includes settings for an `SRV` record, you must specify a value for `AWS_INSTANCE_IPV4` , `AWS_INSTANCE_IPV6` , or both.
	// - **AWS_INSTANCE_PORT** - If the service includes an `SRV` record, the value that you want Route 53 to return for the port.
	//
	// If the service includes `HealthCheckConfig` , the port on the endpoint that you want Route 53 to send requests to.
	//
	// This value is required if you specified settings for an `SRV` record or a Route 53 health check when you created the service.
	InstanceAttributes interface{} `field:"required" json:"instanceAttributes" yaml:"instanceAttributes"`
	// The ID of the service that you want to use for settings for the instance.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// An identifier that you want to associate with the instance. Note the following:.
	//
	// - If the service that's specified by `ServiceId` includes settings for an `SRV` record, the value of `InstanceId` is automatically included as part of the value for the `SRV` record. For more information, see [DnsRecord > Type](https://docs.aws.amazon.com/cloud-map/latest/api/API_DnsRecord.html#cloudmap-Type-DnsRecord-Type) .
	// - You can use this value to update an existing instance.
	// - To register a new instance, you must specify a value that's unique among instances that you register by using the same service.
	// - If you specify an existing `InstanceId` and `ServiceId` , AWS Cloud Map updates the existing DNS records, if any. If there's also an existing health check, AWS Cloud Map deletes the old health check and creates a new one.
	//
	// > The health check isn't deleted immediately, so it will still appear for a while if you submit a `ListHealthChecks` request, for example.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
}

