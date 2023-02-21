package awsstepfunctionstasks


// S3 Data Distribution Type.
type S3DataDistributionType string

const (
	// Fully replicated S3 Data Distribution Type.
	S3DataDistributionType_FULLY_REPLICATED S3DataDistributionType = "FULLY_REPLICATED"
	// Sharded By S3 Key Data Distribution Type.
	S3DataDistributionType_SHARDED_BY_S3_KEY S3DataDistributionType = "SHARDED_BY_S3_KEY"
)

