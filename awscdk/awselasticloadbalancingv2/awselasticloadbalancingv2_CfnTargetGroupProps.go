package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTargetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTargetGroupProps := &cfnTargetGroupProps{
//   	healthCheckEnabled: jsii.Boolean(false),
//   	healthCheckIntervalSeconds: jsii.Number(123),
//   	healthCheckPath: jsii.String("healthCheckPath"),
//   	healthCheckPort: jsii.String("healthCheckPort"),
//   	healthCheckProtocol: jsii.String("healthCheckProtocol"),
//   	healthCheckTimeoutSeconds: jsii.Number(123),
//   	healthyThresholdCount: jsii.Number(123),
//   	ipAddressType: jsii.String("ipAddressType"),
//   	matcher: &matcherProperty{
//   		grpcCode: jsii.String("grpcCode"),
//   		httpCode: jsii.String("httpCode"),
//   	},
//   	name: jsii.String("name"),
//   	port: jsii.Number(123),
//   	protocol: jsii.String("protocol"),
//   	protocolVersion: jsii.String("protocolVersion"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targetGroupAttributes: []interface{}{
//   		&targetGroupAttributeProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	targets: []interface{}{
//   		&targetDescriptionProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			availabilityZone: jsii.String("availabilityZone"),
//   			port: jsii.Number(123),
//   		},
//   	},
//   	targetType: jsii.String("targetType"),
//   	unhealthyThresholdCount: jsii.Number(123),
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnTargetGroupProps struct {
	// Indicates whether health checks are enabled.
	//
	// If the target type is `lambda` , health checks are disabled by default but can be enabled. If the target type is `instance` , `ip` , or `alb` , health checks are always enabled and cannot be disabled.
	HealthCheckEnabled interface{} `field:"optional" json:"healthCheckEnabled" yaml:"healthCheckEnabled"`
	// The approximate amount of time, in seconds, between health checks of an individual target.
	//
	// If the target group protocol is HTTP or HTTPS, the default is 30 seconds. If the target group protocol is TCP, TLS, UDP, or TCP_UDP, the supported values are 10 and 30 seconds and the default is 30 seconds. If the target group protocol is GENEVE, the default is 10 seconds. If the target type is `lambda` , the default is 35 seconds.
	HealthCheckIntervalSeconds *float64 `field:"optional" json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	//
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /.
	//
	// [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is / AWS .ALB/healthcheck.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The port the load balancer uses when performing health checks on targets.
	//
	// If the protocol is HTTP, HTTPS, TCP, TLS, UDP, or TCP_UDP, the default is `traffic-port` , which is the port on which each target receives traffic from the load balancer. If the protocol is GENEVE, the default is port 80.
	HealthCheckPort *string `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	// The protocol the load balancer uses when performing health checks on targets.
	//
	// For Application Load Balancers, the default is HTTP. For Network Load Balancers and Gateway Load Balancers, the default is TCP. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. The GENEVE, TLS, UDP, and TCP_UDP protocols are not supported for health checks.
	HealthCheckProtocol *string `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	//
	// For target groups with a protocol of HTTP, HTTPS, or GENEVE, the default is 5 seconds. For target groups with a protocol of TCP or TLS, this value must be 6 seconds for HTTP health checks and 10 seconds for TCP and HTTPS health checks. If the target type is `lambda` , the default is 30 seconds.
	HealthCheckTimeoutSeconds *float64 `field:"optional" json:"healthCheckTimeoutSeconds" yaml:"healthCheckTimeoutSeconds"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy.
	//
	// For target groups with a protocol of HTTP or HTTPS, the default is 5. For target groups with a protocol of TCP, TLS, or GENEVE, the default is 3. If the target type is `lambda` , the default is 5.
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// The type of IP address used for this target group.
	//
	// The possible values are `ipv4` and `ipv6` . This is an optional parameter. If not specified, the IP address type defaults to `ipv4` .
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port on which the targets receive traffic.
	//
	// This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use for routing traffic to the targets.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, or TCP_UDP. For Gateway Load Balancers, the supported protocol is GENEVE. A TCP_UDP listener must be associated with a TCP_UDP target group. If the target is a Lambda function, this parameter does not apply.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// [HTTP/HTTPS protocol] The protocol version.
	//
	// The possible values are `GRPC` , `HTTP1` , and `HTTP2` .
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// The tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The attributes.
	TargetGroupAttributes interface{} `field:"optional" json:"targetGroupAttributes" yaml:"targetGroupAttributes"`
	// The targets.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The type of target that you must specify when registering targets with this target group.
	//
	// You can't specify targets for a target group using more than one target type.
	//
	// - `instance` - Register targets by instance ID. This is the default value.
	// - `ip` - Register targets by IP address. You can specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	// - `lambda` - Register a single Lambda function as a target.
	// - `alb` - Register a single Application Load Balancer as a target.
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// The number of consecutive health check failures required before considering a target unhealthy.
	//
	// If the target group protocol is HTTP or HTTPS, the default is 2. If the target group protocol is TCP or TLS, this value must be the same as the healthy threshold count. If the target group protocol is GENEVE, the default is 3. If the target type is `lambda` , the default is 2.
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
	// The identifier of the virtual private cloud (VPC).
	//
	// If the target is a Lambda function, this parameter does not apply. Otherwise, this parameter is required.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

