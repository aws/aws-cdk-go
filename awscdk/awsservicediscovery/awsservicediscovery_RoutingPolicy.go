package awsservicediscovery


type RoutingPolicy string

const (
	// Route 53 returns the applicable value from one randomly selected instance from among the instances that you registered using the same service.
	RoutingPolicy_WEIGHTED RoutingPolicy = "WEIGHTED"
	// If you define a health check for the service and the health check is healthy, Route 53 returns the applicable value for up to eight instances.
	RoutingPolicy_MULTIVALUE RoutingPolicy = "MULTIVALUE"
)

