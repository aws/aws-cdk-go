package awscdk


// The type of asset hash.
//
// NOTE: the hash is used in order to identify a specific revision of the asset, and
// used for optimizing and caching deployment activities related to this asset such as
// packaging, uploading to Amazon S3, etc.
type AssetHashType string

const (
	// Based on the content of the source path.
	//
	// When bundling, use `SOURCE` when the content of the bundling output is not
	// stable across repeated bundling operations.
	AssetHashType_SOURCE AssetHashType = "SOURCE"
	// Based on the content of the bundling output.
	//
	// Use `OUTPUT` when the source of the asset is a top level folder containing
	// code and/or dependencies that are not directly linked to the asset.
	AssetHashType_OUTPUT AssetHashType = "OUTPUT"
	// Use a custom hash.
	AssetHashType_CUSTOM AssetHashType = "CUSTOM"
)

