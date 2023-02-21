package pipelines


// Type of the asset that is being published.
// Experimental.
type AssetType string

const (
	// A file.
	// Experimental.
	AssetType_FILE AssetType = "FILE"
	// A Docker image.
	// Experimental.
	AssetType_DOCKER_IMAGE AssetType = "DOCKER_IMAGE"
)

