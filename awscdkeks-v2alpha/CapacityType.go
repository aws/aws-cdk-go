package awscdkeks-v2alpha


// Capacity type of the managed node group.
// Experimental.
type CapacityType string

const (
	// spot instances.
	// Experimental.
	CapacityType_SPOT CapacityType = "SPOT"
	// on-demand instances.
	// Experimental.
	CapacityType_ON_DEMAND CapacityType = "ON_DEMAND"
	// capacity block instances.
	// Experimental.
	CapacityType_CAPACITY_BLOCK CapacityType = "CAPACITY_BLOCK"
)

