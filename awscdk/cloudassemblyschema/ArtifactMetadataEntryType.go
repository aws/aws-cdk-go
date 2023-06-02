package cloudassemblyschema


// Type of artifact metadata entry.
type ArtifactMetadataEntryType string

const (
	// Asset in metadata.
	ArtifactMetadataEntryType_ASSET ArtifactMetadataEntryType = "ASSET"
	// Metadata key used to print INFO-level messages by the toolkit when an app is syntheized.
	ArtifactMetadataEntryType_INFO ArtifactMetadataEntryType = "INFO"
	// Metadata key used to print WARNING-level messages by the toolkit when an app is syntheized.
	ArtifactMetadataEntryType_WARN ArtifactMetadataEntryType = "WARN"
	// Metadata key used to print ERROR-level messages by the toolkit when an app is syntheized.
	ArtifactMetadataEntryType_ERROR ArtifactMetadataEntryType = "ERROR"
	// Represents the CloudFormation logical ID of a resource at a certain path.
	ArtifactMetadataEntryType_LOGICAL_ID ArtifactMetadataEntryType = "LOGICAL_ID"
	// Represents tags of a stack.
	ArtifactMetadataEntryType_STACK_TAGS ArtifactMetadataEntryType = "STACK_TAGS"
)

