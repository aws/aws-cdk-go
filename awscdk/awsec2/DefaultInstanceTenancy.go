package awsec2


// The default tenancy of instances launched into the VPC.
type DefaultInstanceTenancy string

const (
	// Instances can be launched with any tenancy.
	DefaultInstanceTenancy_DEFAULT DefaultInstanceTenancy = "DEFAULT"
	// Any instance launched into the VPC automatically has dedicated tenancy, unless you launch it with the default tenancy.
	DefaultInstanceTenancy_DEDICATED DefaultInstanceTenancy = "DEDICATED"
	// Instances must be launched on dedicated hosts.
	//
	// This provides additional visibility
	// and control over instance placement at the physical host level.
	DefaultInstanceTenancy_HOST DefaultInstanceTenancy = "HOST"
)

