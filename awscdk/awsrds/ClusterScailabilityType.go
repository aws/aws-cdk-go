package awsrds


// The scalability mode of the Aurora DB cluster.
// Deprecated: Use ClusterScalabilityType instead. This will be removed in the next major version.
type ClusterScailabilityType string

const (
	// The cluster uses normal DB instance creation.
	// Deprecated: Use ClusterScalabilityType instead. This will be removed in the next major version.
	ClusterScailabilityType_STANDARD ClusterScailabilityType = "STANDARD"
	// The cluster operates as an Aurora Limitless Database, allowing you to create a DB shard group for horizontal scaling (sharding) capabilities.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/limitless.html
	//
	// Deprecated: Use ClusterScalabilityType instead. This will be removed in the next major version.
	ClusterScailabilityType_LIMITLESS ClusterScailabilityType = "LIMITLESS"
)

