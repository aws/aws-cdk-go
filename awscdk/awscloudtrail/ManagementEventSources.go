package awscloudtrail


// Types of management event sources that can be excluded.
type ManagementEventSources string

const (
	// AWS Key Management Service (AWS KMS) events.
	ManagementEventSources_KMS ManagementEventSources = "KMS"
	// Data API events.
	ManagementEventSources_RDS_DATA_API ManagementEventSources = "RDS_DATA_API"
)

