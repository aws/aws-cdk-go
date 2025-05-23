package pipelines


// Type of the asset that is being published.
type AssetType string

const (
	// A file.
	AssetType_FILE AssetType = "FILE"
	// A Docker image.
	AssetType_DOCKER_IMAGE AssetType = "DOCKER_IMAGE"
)

