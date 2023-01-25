package awscloudtrail


// Types of management event sources that can be excluded.
// Experimental.
type ManagementEventSources string

const (
	// AWS Key Management Service (AWS KMS) events.
	// Experimental.
	ManagementEventSources_KMS ManagementEventSources = "KMS"
	// Data API events.
	// Experimental.
	ManagementEventSources_RDS_DATA_API ManagementEventSources = "RDS_DATA_API"
)

