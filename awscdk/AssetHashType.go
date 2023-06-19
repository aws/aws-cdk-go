package awscdk


// The type of asset hash.
//
// NOTE: the hash is used in order to identify a specific revision of the asset, and
// used for optimizing and caching deployment activities related to this asset such as
// packaging, uploading to Amazon S3, etc.
// Experimental.
type AssetHashType string

const (
	// Based on the content of the source path.
	//
	// When bundling, use `SOURCE` when the content of the bundling output is not
	// stable across repeated bundling operations.
	// Experimental.
	AssetHashType_SOURCE AssetHashType = "SOURCE"
	// Based on the content of the bundled path.
	// Deprecated: use `OUTPUT` instead.
	AssetHashType_BUNDLE AssetHashType = "BUNDLE"
	// Based on the content of the bundling output.
	//
	// Use `OUTPUT` when the source of the asset is a top level folder containing
	// code and/or dependencies that are not directly linked to the asset.
	// Experimental.
	AssetHashType_OUTPUT AssetHashType = "OUTPUT"
	// Use a custom hash.
	// Experimental.
	AssetHashType_CUSTOM AssetHashType = "CUSTOM"
)

