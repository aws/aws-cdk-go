package awsec2


// Determines how this placement group spreads instances.
type PlacementGroupSpreadLevel string

const (
	// Host spread level placement groups are only available with AWS Outposts.
	//
	// For host spread level placement groups, there are no restrictions for running instances per Outposts.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups-outpost.html
	//
	PlacementGroupSpreadLevel_HOST PlacementGroupSpreadLevel = "HOST"
	// Each instance is launched on a separate rack.
	//
	// Each has its own network and power source.
	// A rack spread placement group can span multiple Availability Zones in the same Region.
	// For rack spread level placement groups, you can have a maximum of seven running instances per Availability Zone per group.
	PlacementGroupSpreadLevel_RACK PlacementGroupSpreadLevel = "RACK"
)

