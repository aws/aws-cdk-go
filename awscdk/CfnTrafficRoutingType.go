package awscdk


// The possible types of traffic shifting for the blue-green deployment configuration.
//
// The type of the {@link CfnTrafficRoutingConfig.type} property.
// Experimental.
type CfnTrafficRoutingType string

const (
	// Switch from blue to green at once.
	// Experimental.
	CfnTrafficRoutingType_ALL_AT_ONCE CfnTrafficRoutingType = "ALL_AT_ONCE"
	// Specifies a configuration that shifts traffic from blue to green in two increments.
	// Experimental.
	CfnTrafficRoutingType_TIME_BASED_CANARY CfnTrafficRoutingType = "TIME_BASED_CANARY"
	// Specifies a configuration that shifts traffic from blue to green in equal increments, with an equal number of minutes between each increment.
	// Experimental.
	CfnTrafficRoutingType_TIME_BASED_LINEAR CfnTrafficRoutingType = "TIME_BASED_LINEAR"
)

