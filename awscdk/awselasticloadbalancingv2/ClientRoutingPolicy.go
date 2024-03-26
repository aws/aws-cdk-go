package awselasticloadbalancingv2


// Indicates how traffic is distributed among the load balancer Availability Zones.
//
// Example:
//   var vpc vpc
//
//
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   	// Whether deletion protection is enabled.
//   	DeletionProtection: jsii.Boolean(true),
//
//   	// Whether cross-zone load balancing is enabled.
//   	CrossZoneEnabled: jsii.Boolean(true),
//
//   	// Whether the load balancer blocks traffic through the Internet Gateway (IGW).
//   	DenyAllIgwTraffic: jsii.Boolean(false),
//
//   	// Indicates how traffic is distributed among the load balancer Availability Zones.
//   	ClientRoutingPolicy: elbv2.ClientRoutingPolicy_AVAILABILITY_ZONE_AFFINITY,
//   })
//
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#zonal-dns-affinity
//
type ClientRoutingPolicy string

const (
	// 100 percent zonal affinity.
	ClientRoutingPolicy_AVAILABILITY_ZONE_AFFINITY ClientRoutingPolicy = "AVAILABILITY_ZONE_AFFINITY"
	// 85 percent zonal affinity.
	ClientRoutingPolicy_PARTIAL_AVAILABILITY_ZONE_AFFINITY ClientRoutingPolicy = "PARTIAL_AVAILABILITY_ZONE_AFFINITY"
	// No zonal affinity.
	ClientRoutingPolicy_ANY_AVAILABILITY_ZONE ClientRoutingPolicy = "ANY_AVAILABILITY_ZONE"
)

