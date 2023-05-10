package cloudassemblyschema


// Type of artifact metadata entry.
// Experimental.
type ArtifactMetadataEntryType string

const (
	// Asset in metadata.
	// Experimental.
	ArtifactMetadataEntryType_ASSET ArtifactMetadataEntryType = "ASSET"
	// Metadata key used to print INFO-level messages by the toolkit when an app is syntheized.
	// Experimental.
	ArtifactMetadataEntryType_INFO ArtifactMetadataEntryType = "INFO"
	// Metadata key used to print WARNING-level messages by the toolkit when an app is syntheized.
	// Experimental.
	ArtifactMetadataEntryType_WARN ArtifactMetadataEntryType = "WARN"
	// Metadata key used to print ERROR-level messages by the toolkit when an app is syntheized.
	// Experimental.
	ArtifactMetadataEntryType_ERROR ArtifactMetadataEntryType = "ERROR"
	// Represents the CloudFormation logical ID of a resource at a certain path.
	// Experimental.
	ArtifactMetadataEntryType_LOGICAL_ID ArtifactMetadataEntryType = "LOGICAL_ID"
	// Represents tags of a stack.
	// Experimental.
	ArtifactMetadataEntryType_STACK_TAGS ArtifactMetadataEntryType = "STACK_TAGS"
)

