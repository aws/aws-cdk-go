package awseksv2


// Capacity type of the managed node group.
type CapacityType string

const (
	// spot instances.
	CapacityType_SPOT CapacityType = "SPOT"
	// on-demand instances.
	CapacityType_ON_DEMAND CapacityType = "ON_DEMAND"
	// capacity block instances.
	CapacityType_CAPACITY_BLOCK CapacityType = "CAPACITY_BLOCK"
)

