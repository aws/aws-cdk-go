package awsautoscaling


// Indicates how to allocate instance types to fulfill Spot capacity.
type SpotAllocationStrategy string

const (
	// The Auto Scaling group launches instances using the Spot pools with the lowest price, and evenly allocates your instances across the number of Spot pools that you specify.
	SpotAllocationStrategy_LOWEST_PRICE SpotAllocationStrategy = "LOWEST_PRICE"
	// The Auto Scaling group launches instances using Spot pools that are optimally chosen based on the available Spot capacity.
	//
	// Recommended.
	SpotAllocationStrategy_CAPACITY_OPTIMIZED SpotAllocationStrategy = "CAPACITY_OPTIMIZED"
	// When you use this strategy, you need to set the order of instance types in the list of launch template overrides from highest to lowest priority (from first to last in the list).
	//
	// Amazon EC2 Auto Scaling
	// honors the instance type priorities on a best-effort basis but optimizes for capacity first.
	SpotAllocationStrategy_CAPACITY_OPTIMIZED_PRIORITIZED SpotAllocationStrategy = "CAPACITY_OPTIMIZED_PRIORITIZED"
	// The price and capacity optimized allocation strategy looks at both price and capacity to select the Spot Instance pools that are the least likely to be interrupted and have the lowest possible price.
	SpotAllocationStrategy_PRICE_CAPACITY_OPTIMIZED SpotAllocationStrategy = "PRICE_CAPACITY_OPTIMIZED"
)

