package awselasticloadbalancingv2


// Specifies an attribute for an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerAttributeProperty := &LoadBalancerAttributeProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattribute.html
//
type CfnLoadBalancer_LoadBalancerAttributeProperty struct {
	// The name of the attribute.
	//
	// The following attributes are supported by all load balancers:
	//
	// - `deletion_protection.enabled` - Indicates whether deletion protection is enabled. The value is `true` or `false` . The default is `false` .
	// - `load_balancing.cross_zone.enabled` - Indicates whether cross-zone load balancing is enabled. The possible values are `true` and `false` . The default for Network Load Balancers and Gateway Load Balancers is `false` . The default for Application Load Balancers is `true` , and cannot be changed.
	//
	// The following attributes are supported by both Application Load Balancers and Network Load Balancers:
	//
	// - `access_logs.s3.enabled` - Indicates whether access logs are enabled. The value is `true` or `false` . The default is `false` .
	// - `access_logs.s3.bucket` - The name of the S3 bucket for the access logs. This attribute is required if access logs are enabled. The bucket must exist in the same region as the load balancer and have a bucket policy that grants Elastic Load Balancing permissions to write to the bucket.
	// - `access_logs.s3.prefix` - The prefix for the location in the S3 bucket for the access logs.
	// - `ipv6.deny_all_igw_traffic` - Blocks internet gateway (IGW) access to the load balancer. It is set to `false` for internet-facing load balancers and `true` for internal load balancers, preventing unintended access to your internal load balancer through an internet gateway.
	//
	// The following attributes are supported by only Application Load Balancers:
	//
	// - `idle_timeout.timeout_seconds` - The idle timeout value, in seconds. The valid range is 1-4000 seconds. The default is 60 seconds.
	// - `client_keep_alive.seconds` - The client keep alive value, in seconds. The valid range is 60-604800 seconds. The default is 3600 seconds.
	// - `connection_logs.s3.enabled` - Indicates whether connection logs are enabled. The value is `true` or `false` . The default is `false` .
	// - `connection_logs.s3.bucket` - The name of the S3 bucket for the connection logs. This attribute is required if connection logs are enabled. The bucket must exist in the same region as the load balancer and have a bucket policy that grants Elastic Load Balancing permissions to write to the bucket.
	// - `connection_logs.s3.prefix` - The prefix for the location in the S3 bucket for the connection logs.
	// - `routing.http.desync_mitigation_mode` - Determines how the load balancer handles requests that might pose a security risk to your application. The possible values are `monitor` , `defensive` , and `strictest` . The default is `defensive` .
	// - `routing.http.drop_invalid_header_fields.enabled` - Indicates whether HTTP headers with invalid header fields are removed by the load balancer ( `true` ) or routed to targets ( `false` ). The default is `false` .
	// - `routing.http.preserve_host_header.enabled` - Indicates whether the Application Load Balancer should preserve the `Host` header in the HTTP request and send it to the target without any change. The possible values are `true` and `false` . The default is `false` .
	// - `routing.http.x_amzn_tls_version_and_cipher_suite.enabled` - Indicates whether the two headers ( `x-amzn-tls-version` and `x-amzn-tls-cipher-suite` ), which contain information about the negotiated TLS version and cipher suite, are added to the client request before sending it to the target. The `x-amzn-tls-version` header has information about the TLS protocol version negotiated with the client, and the `x-amzn-tls-cipher-suite` header has information about the cipher suite negotiated with the client. Both headers are in OpenSSL format. The possible values for the attribute are `true` and `false` . The default is `false` .
	// - `routing.http.xff_client_port.enabled` - Indicates whether the `X-Forwarded-For` header should preserve the source port that the client used to connect to the load balancer. The possible values are `true` and `false` . The default is `false` .
	// - `routing.http.xff_header_processing.mode` - Enables you to modify, preserve, or remove the `X-Forwarded-For` header in the HTTP request before the Application Load Balancer sends the request to the target. The possible values are `append` , `preserve` , and `remove` . The default is `append` .
	//
	// - If the value is `append` , the Application Load Balancer adds the client IP address (of the last hop) to the `X-Forwarded-For` header in the HTTP request before it sends it to targets.
	// - If the value is `preserve` the Application Load Balancer preserves the `X-Forwarded-For` header in the HTTP request, and sends it to targets without any change.
	// - If the value is `remove` , the Application Load Balancer removes the `X-Forwarded-For` header in the HTTP request before it sends it to targets.
	// - `routing.http2.enabled` - Indicates whether HTTP/2 is enabled. The possible values are `true` and `false` . The default is `true` . Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens.
	// - `waf.fail_open.enabled` - Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. The possible values are `true` and `false` . The default is `false` .
	//
	// The following attributes are supported by only Network Load Balancers:
	//
	// - `dns_record.client_routing_policy` - Indicates how traffic is distributed among the load balancer Availability Zones. The possible values are `availability_zone_affinity` with 100 percent zonal affinity, `partial_availability_zone_affinity` with 85 percent zonal affinity, and `any_availability_zone` with 0 percent zonal affinity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattribute.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattribute-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattribute.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattribute-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

