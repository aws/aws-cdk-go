package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for defining an Application Target Group.
//
// Example:
//   var vpc vpc
//
//
//   tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TG"), &ApplicationTargetGroupProps{
//   	TargetType: elbv2.TargetType_IP,
//   	Port: jsii.Number(50051),
//   	Protocol: elbv2.ApplicationProtocol_HTTP,
//   	ProtocolVersion: elbv2.ApplicationProtocolVersion_GRPC,
//   	HealthCheck: &HealthCheck{
//   		Enabled: jsii.Boolean(true),
//   		HealthyGrpcCodes: jsii.String("0-99"),
//   	},
//   	Vpc: Vpc,
//   })
//
type ApplicationTargetGroupProps struct {
	// Indicates whether cross zone load balancing is enabled.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattribute.html
	//
	// Default: - use load balancer configuration.
	//
	CrossZoneEnabled *bool `field:"optional" json:"crossZoneEnabled" yaml:"crossZoneEnabled"`
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	// Default: 300.
	//
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	// Default: - The default value for each property in this configuration varies depending on the target.
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The type of IP addresses of the targets registered with the target group.
	// Default: undefined - ELB defaults to IPv4.
	//
	IpAddressType TargetGroupIpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	// Default: - Automatically generated.
	//
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	// Default: - Determined automatically.
	//
	TargetType TargetType `field:"optional" json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`.
	// Default: - undefined.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Indicates whether anomaly mitigation is enabled.
	//
	// Only available when `loadBalancingAlgorithmType` is `TargetGroupLoadBalancingAlgorithmType.WEIGHTED_RANDOM`
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#automatic-target-weights
	//
	// Default: false.
	//
	EnableAnomalyMitigation *bool `field:"optional" json:"enableAnomalyMitigation" yaml:"enableAnomalyMitigation"`
	// The load balancing algorithm to select targets for routing requests.
	// Default: TargetGroupLoadBalancingAlgorithmType.ROUND_ROBIN
	//
	LoadBalancingAlgorithmType TargetGroupLoadBalancingAlgorithmType `field:"optional" json:"loadBalancingAlgorithmType" yaml:"loadBalancingAlgorithmType"`
	// Indicates whether the target group supports multi-value headers.
	//
	// If the value is true, the request and response headers exchanged between
	// the load balancer and the Lambda function include arrays of values or strings.
	//
	// Only applicable for Lambda targets.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
	//
	// Default: false.
	//
	MultiValueHeadersEnabled *bool `field:"optional" json:"multiValueHeadersEnabled" yaml:"multiValueHeadersEnabled"`
	// The port on which the target receives traffic.
	//
	// This is not applicable for Lambda targets.
	// Default: - Determined from protocol if known.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol used for communication with the target.
	//
	// This is not applicable for Lambda targets.
	// Default: - Determined from port if known.
	//
	Protocol ApplicationProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The protocol version to use.
	// Default: ApplicationProtocolVersion.HTTP1
	//
	ProtocolVersion ApplicationProtocolVersion `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// The time period during which the load balancer sends a newly registered target a linearly increasing share of the traffic to the target group.
	//
	// The range is 30-900 seconds (15 minutes).
	// Default: 0.
	//
	SlowStart awscdk.Duration `field:"optional" json:"slowStart" yaml:"slowStart"`
	// The stickiness cookie expiration period.
	//
	// Setting this value enables load balancer stickiness.
	//
	// After this period, the cookie is considered stale. The minimum value is
	// 1 second and the maximum value is 7 days (604800 seconds).
	// Default: - Stickiness is disabled.
	//
	StickinessCookieDuration awscdk.Duration `field:"optional" json:"stickinessCookieDuration" yaml:"stickinessCookieDuration"`
	// The name of an application-based stickiness cookie.
	//
	// Names that start with the following prefixes are not allowed: AWSALB, AWSALBAPP,
	// and AWSALBTG; they're reserved for use by the load balancer.
	//
	// Note: `stickinessCookieName` parameter depends on the presence of `stickinessCookieDuration` parameter.
	// If `stickinessCookieDuration` is not set, `stickinessCookieName` will be omitted.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/sticky-sessions.html
	//
	// Default: - If `stickinessCookieDuration` is set, a load-balancer generated cookie is used. Otherwise, no stickiness is defined.
	//
	StickinessCookieName *string `field:"optional" json:"stickinessCookieName" yaml:"stickinessCookieName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	// Default: - No targets.
	//
	Targets *[]IApplicationLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

