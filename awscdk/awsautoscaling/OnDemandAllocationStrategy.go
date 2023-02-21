package awsautoscaling


// Indicates how to allocate instance types to fulfill On-Demand capacity.
type OnDemandAllocationStrategy string

const (
	// This strategy uses the order of instance types in the LaunchTemplateOverrides to define the launch priority of each instance type.
	//
	// The first instance type in the array is prioritized higher than the
	// last. If all your On-Demand capacity cannot be fulfilled using your highest priority instance, then
	// the Auto Scaling group launches the remaining capacity using the second priority instance type, and
	// so on.
	OnDemandAllocationStrategy_PRIORITIZED OnDemandAllocationStrategy = "PRIORITIZED"
)

