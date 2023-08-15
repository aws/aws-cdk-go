package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for adding new targets to a listener.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var asg autoScalingGroup
//   var vpc vpc
//
//
//   // Create the load balancer in a VPC. 'internetFacing' is 'false'
//   // by default, which creates an internal load balancer.
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   // Add a listener and open up the load balancer's security group
//   // to the world.
//   listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//
//   	// 'open: true' is the default, you can leave it out if you want. Set it
//   	// to 'false' and use `listener.connections` if you want to be selective
//   	// about who can access the load balancer.
//   	Open: jsii.Boolean(true),
//   })
//
//   // Create an AutoScaling group and add it as a load balancing
//   // target to the listener.
//   listener.AddTargets(jsii.String("ApplicationFleet"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(8080),
//   	Targets: []iApplicationLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
type AddApplicationTargetsProps struct {
	// Rule applies if matches the conditions.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
	//
	// Default: - No conditions.
	//
	Conditions *[]ListenerCondition `field:"optional" json:"conditions" yaml:"conditions"`
	// Priority of this target group.
	//
	// The rule with the lowest priority will be used for every request.
	// If priority is not given, these target groups will be added as
	// defaults, and must not have conditions.
	//
	// Priorities must be unique.
	// Default: Target groups are used as defaults.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	// Default: Duration.minutes(5)
	//
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	// Default: - The default value for each property in this configuration varies depending on the target.
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The load balancing algorithm to select targets for routing requests.
	// Default: round_robin.
	//
	LoadBalancingAlgorithmType TargetGroupLoadBalancingAlgorithmType `field:"optional" json:"loadBalancingAlgorithmType" yaml:"loadBalancingAlgorithmType"`
	// The port on which the listener listens for requests.
	// Default: Determined from protocol if known.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The protocol to use.
	// Default: Determined from port if known.
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
	// Default: Stickiness disabled.
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
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	// Default: Automatically generated.
	//
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. All target must be of the same type.
	Targets *[]IApplicationLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

