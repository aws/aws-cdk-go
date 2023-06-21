package awscdkgameliftalpha


// Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances in the game server group.
// Experimental.
type BalancingStrategy string

const (
	// Only Spot Instances are used in the game server group.
	//
	// If Spot Instances are unavailable or not viable for game hosting, the game server group provides no hosting capacity until Spot Instances can again be used.
	// Until then, no new instances are started, and the existing nonviable Spot Instances are terminated (after current gameplay ends) and are not replaced.
	// Experimental.
	BalancingStrategy_SPOT_ONLY BalancingStrategy = "SPOT_ONLY"
	// Spot Instances are used whenever available in the game server group.
	//
	// If Spot Instances are unavailable, the game server group continues to provide hosting capacity by falling back to On-Demand Instances.
	// Existing nonviable Spot Instances are terminated (after current gameplay ends) and are replaced with new On-Demand Instances.
	// Experimental.
	BalancingStrategy_SPOT_PREFERRED BalancingStrategy = "SPOT_PREFERRED"
	// Only On-Demand Instances are used in the game server group.
	//
	// No Spot Instances are used, even when available, while this balancing strategy is in force.
	// Experimental.
	BalancingStrategy_ON_DEMAND_ONLY BalancingStrategy = "ON_DEMAND_ONLY"
)

