package awsbatch


// Properties for how to prepare compute resources that are provisioned for a compute environment.
// Experimental.
type AllocationStrategy string

const (
	// Batch will use the best fitting instance type will be used when assigning a batch job in this compute environment.
	// Experimental.
	AllocationStrategy_BEST_FIT AllocationStrategy = "BEST_FIT"
	// Batch will select additional instance types that are large enough to meet the requirements of the jobs in the queue, with a preference for instance types with a lower cost per unit vCPU.
	// Experimental.
	AllocationStrategy_BEST_FIT_PROGRESSIVE AllocationStrategy = "BEST_FIT_PROGRESSIVE"
	// This is only available for Spot Instance compute resources and will select additional instance types that are large enough to meet the requirements of the jobs in the queue, with a preference for instance types that are less likely to be interrupted.
	// Experimental.
	AllocationStrategy_SPOT_CAPACITY_OPTIMIZED AllocationStrategy = "SPOT_CAPACITY_OPTIMIZED"
)

