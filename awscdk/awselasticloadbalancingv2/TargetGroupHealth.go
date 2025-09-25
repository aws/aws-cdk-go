package awselasticloadbalancingv2


// Properties for configuring a target group health.
//
// Example:
//   var vpc vpc
//
//
//   targetGroup := elbv2.NewApplicationTargetGroup(this, jsii.String("TargetGroup"), &ApplicationTargetGroupProps{
//   	Vpc: Vpc,
//   	Port: jsii.Number(80),
//   	TargetGroupHealth: &TargetGroupHealth{
//   		DnsMinimumHealthyTargetCount: jsii.Number(3),
//   		DnsMinimumHealthyTargetPercentage: jsii.Number(70),
//   		RoutingMinimumHealthyTargetCount: jsii.Number(2),
//   		RoutingMinimumHealthyTargetPercentage: jsii.Number(50),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-attributes
//
type TargetGroupHealth struct {
	// The minimum number of targets that must be healthy for DNS failover.
	//
	// If below this value, mark the zone as unhealthy in DNS.
	// Use 0 for "off".
	// Default: 1.
	//
	DnsMinimumHealthyTargetCount *float64 `field:"optional" json:"dnsMinimumHealthyTargetCount" yaml:"dnsMinimumHealthyTargetCount"`
	// The minimum percentage of targets that must be healthy for DNS failover.
	//
	// If below this value, mark the zone as unhealthy in DNS.
	// Use 0 for "off".
	// Default: 0.
	//
	DnsMinimumHealthyTargetPercentage *float64 `field:"optional" json:"dnsMinimumHealthyTargetPercentage" yaml:"dnsMinimumHealthyTargetPercentage"`
	// The minimum number of targets that must be healthy for unhealthy state routing.
	//
	// If below this value, send traffic to all targets including unhealthy ones.
	// Default: 1.
	//
	RoutingMinimumHealthyTargetCount *float64 `field:"optional" json:"routingMinimumHealthyTargetCount" yaml:"routingMinimumHealthyTargetCount"`
	// The minimum percentage of targets that must be healthy for unhealthy state routing.
	//
	// If below this value, send traffic to all targets including unhealthy ones.
	// Use 0 for "off".
	// Default: 0.
	//
	RoutingMinimumHealthyTargetPercentage *float64 `field:"optional" json:"routingMinimumHealthyTargetPercentage" yaml:"routingMinimumHealthyTargetPercentage"`
}

