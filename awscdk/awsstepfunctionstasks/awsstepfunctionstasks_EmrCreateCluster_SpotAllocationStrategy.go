package awsstepfunctionstasks


// Spot Allocation Strategies.
//
// Specifies the strategy to use in launching Spot Instance fleets. For example, "capacity-optimized" launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_SpotProvisioningSpecification.html
//
type EmrCreateCluster_SpotAllocationStrategy string

const (
	// Capacity-optimized, which launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
	EmrCreateCluster_SpotAllocationStrategy_CAPACITY_OPTIMIZED EmrCreateCluster_SpotAllocationStrategy = "CAPACITY_OPTIMIZED"
)

