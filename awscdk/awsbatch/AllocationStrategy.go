package awsbatch


// Determines how this compute environment chooses instances to spawn.
//
// Example:
//   var vpc IVpc
//
//
//   computeEnv := batch.NewManagedEc2EcsComputeEnvironment(this, jsii.String("myEc2ComputeEnv"), &ManagedEc2EcsComputeEnvironmentProps{
//   	Vpc: Vpc,
//   	AllocationStrategy: batch.AllocationStrategy_BEST_FIT,
//   })
//
// See: https://aws.amazon.com/blogs/compute/optimizing-for-cost-availability-and-throughput-by-selecting-your-aws-batch-allocation-strategy/
//
type AllocationStrategy string

const (
	// Batch chooses the lowest-cost instance type that fits all the jobs in the queue.
	//
	// If instances of that type are not available, the queue will not choose a new type;
	// instead, it will wait for the instance to become available.
	// This can stall your `Queue`, with your compute environment only using part of its max capacity
	// (or none at all) until the `BEST_FIT` instance becomes available.
	// This allocation strategy keeps costs lower but can limit scaling.
	// `BEST_FIT` isn't supported when updating compute environments.
	AllocationStrategy_BEST_FIT AllocationStrategy = "BEST_FIT"
	// This is the default Allocation Strategy if `spot` is `false` or unspecified.
	//
	// This strategy will examine the Jobs in the queue and choose whichever instance type meets the requirements
	// of the jobs in the queue and with the lowest cost per vCPU, just as `BEST_FIT`.
	// However, if not all of the capacity can be filled with this instance type,
	// it will choose a new next-best instance type to run any jobs that couldn’t fit into the `BEST_FIT` capacity.
	// To make the most use of this allocation strategy,
	// it is recommended to use as many instance classes as is feasible for your workload.
	AllocationStrategy_BEST_FIT_PROGRESSIVE AllocationStrategy = "BEST_FIT_PROGRESSIVE"
	// If your workflow tolerates interruptions, you should enable `spot` on your `ComputeEnvironment` and use `SPOT_CAPACITY_OPTIMIZED` (this is the default if `spot` is enabled).
	//
	// This will tell Batch to choose the instance types from the ones you’ve specified that have
	// the most spot capacity available to minimize the chance of interruption.
	// To get the most benefit from your spot instances,
	// you should allow Batch to choose from as many different instance types as possible.
	AllocationStrategy_SPOT_CAPACITY_OPTIMIZED AllocationStrategy = "SPOT_CAPACITY_OPTIMIZED"
	// The price and capacity optimized allocation strategy looks at both price and capacity to select the Spot Instance pools that are the least likely to be interrupted and have the lowest possible price.
	//
	// The Batch team recommends this over `SPOT_CAPACITY_OPTIMIZED` in most instances.
	AllocationStrategy_SPOT_PRICE_CAPACITY_OPTIMIZED AllocationStrategy = "SPOT_PRICE_CAPACITY_OPTIMIZED"
)

