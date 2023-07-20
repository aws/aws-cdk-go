package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTargetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTargetGroupProps := &CfnTargetGroupProps{
//   	HealthCheckEnabled: jsii.Boolean(false),
//   	HealthCheckIntervalSeconds: jsii.Number(123),
//   	HealthCheckPath: jsii.String("healthCheckPath"),
//   	HealthCheckPort: jsii.String("healthCheckPort"),
//   	HealthCheckProtocol: jsii.String("healthCheckProtocol"),
//   	HealthCheckTimeoutSeconds: jsii.Number(123),
//   	HealthyThresholdCount: jsii.Number(123),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Matcher: &MatcherProperty{
//   		GrpcCode: jsii.String("grpcCode"),
//   		HttpCode: jsii.String("httpCode"),
//   	},
//   	Name: jsii.String("name"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	ProtocolVersion: jsii.String("protocolVersion"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetGroupAttributes: []interface{}{
//   		&TargetGroupAttributeProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Targets: []interface{}{
//   		&TargetDescriptionProperty{
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			Port: jsii.Number(123),
//   		},
//   	},
//   	TargetType: jsii.String("targetType"),
//   	UnhealthyThresholdCount: jsii.Number(123),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html
//
type CfnTargetGroupProps struct {
	// Indicates whether health checks are enabled.
	//
	// If the target type is `lambda` , health checks are disabled by default but can be enabled. If the target type is `instance` , `ip` , or `alb` , health checks are always enabled and cannot be disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckenabled
	//
	HealthCheckEnabled interface{} `field:"optional" json:"healthCheckEnabled" yaml:"healthCheckEnabled"`
	// The approximate amount of time, in seconds, between health checks of an individual target.
	//
	// The range is 5-300. If the target group protocol is TCP, TLS, UDP, TCP_UDP, HTTP or HTTPS, the default is 30 seconds. If the target group protocol is GENEVE, the default is 10 seconds. If the target type is `lambda` , the default is 35 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckintervalseconds
	//
	HealthCheckIntervalSeconds *float64 `field:"optional" json:"healthCheckIntervalSeconds" yaml:"healthCheckIntervalSeconds"`
	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	//
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /.
	//
	// [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is / AWS .ALB/healthcheck.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckpath
	//
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// The port the load balancer uses when performing health checks on targets.
	//
	// If the protocol is HTTP, HTTPS, TCP, TLS, UDP, or TCP_UDP, the default is `traffic-port` , which is the port on which each target receives traffic from the load balancer. If the protocol is GENEVE, the default is port 80.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckport
	//
	HealthCheckPort *string `field:"optional" json:"healthCheckPort" yaml:"healthCheckPort"`
	// The protocol the load balancer uses when performing health checks on targets.
	//
	// For Application Load Balancers, the default is HTTP. For Network Load Balancers and Gateway Load Balancers, the default is TCP. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. The GENEVE, TLS, UDP, and TCP_UDP protocols are not supported for health checks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthcheckprotocol
	//
	HealthCheckProtocol *string `field:"optional" json:"healthCheckProtocol" yaml:"healthCheckProtocol"`
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	//
	// The range is 2–120 seconds. For target groups with a protocol of HTTP, the default is 6 seconds. For target groups with a protocol of TCP, TLS or HTTPS, the default is 10 seconds. For target groups with a protocol of GENEVE, the default is 5 seconds. If the target type is `lambda` , the default is 30 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthchecktimeoutseconds
	//
	HealthCheckTimeoutSeconds *float64 `field:"optional" json:"healthCheckTimeoutSeconds" yaml:"healthCheckTimeoutSeconds"`
	// The number of consecutive health check successes required before considering a target healthy.
	//
	// The range is 2-10. If the target group protocol is TCP, TCP_UDP, UDP, TLS, HTTP or HTTPS, the default is 5. For target groups with a protocol of GENEVE, the default is 5. If the target type is `lambda` , the default is 5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-healthythresholdcount
	//
	HealthyThresholdCount *float64 `field:"optional" json:"healthyThresholdCount" yaml:"healthyThresholdCount"`
	// The type of IP address used for this target group.
	//
	// The possible values are `ipv4` and `ipv6` . This is an optional parameter. If not specified, the IP address type defaults to `ipv4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
	//
	// For target groups with a protocol of TCP, TCP_UDP, UDP or TLS the range is 200-599. For target groups with a protocol of HTTP or HTTPS, the range is 200-499. For target groups with a protocol of GENEVE, the range is 200-399.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-matcher
	//
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port on which the targets receive traffic.
	//
	// This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use for routing traffic to the targets.
	//
	// For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, or TCP_UDP. For Gateway Load Balancers, the supported protocol is GENEVE. A TCP_UDP listener must be associated with a TCP_UDP target group. If the target is a Lambda function, this parameter does not apply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// [HTTP/HTTPS protocol] The protocol version.
	//
	// The possible values are `GRPC` , `HTTP1` , and `HTTP2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-protocolversion
	//
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// The tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-targetgroupattributes
	//
	TargetGroupAttributes interface{} `field:"optional" json:"targetGroupAttributes" yaml:"targetGroupAttributes"`
	// The targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// The type of target that you must specify when registering targets with this target group.
	//
	// You can't specify targets for a target group using more than one target type.
	//
	// - `instance` - Register targets by instance ID. This is the default value.
	// - `ip` - Register targets by IP address. You can specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group, the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10). You can't specify publicly routable IP addresses.
	// - `lambda` - Register a single Lambda function as a target.
	// - `alb` - Register a single Application Load Balancer as a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-targettype
	//
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// The number of consecutive health check failures required before considering a target unhealthy.
	//
	// The range is 2-10. If the target group protocol is TCP, TCP_UDP, UDP, TLS, HTTP or HTTPS, the default is 2. For target groups with a protocol of GENEVE, the default is 2. If the target type is `lambda` , the default is 5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-unhealthythresholdcount
	//
	UnhealthyThresholdCount *float64 `field:"optional" json:"unhealthyThresholdCount" yaml:"unhealthyThresholdCount"`
	// The identifier of the virtual private cloud (VPC).
	//
	// If the target is a Lambda function, this parameter does not apply. Otherwise, this parameter is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#cfn-elasticloadbalancingv2-targetgroup-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

