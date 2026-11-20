package awscdk


// The type of asset hash.
//
// NOTE: the hash is used in order to identify a specific revision of the asset, and
// used for optimizing and caching deployment activities related to this asset such as
// packaging, uploading to Amazon S3, etc.
type AssetHashType string

const (
	// Base the asset hash on the contents of the source path.
	//
	// ## Without bundling
	//
	// When used without the `bundling` property, and the asset is just a regular file
	// copy operation: the hash will be based on the files in the source path
	// respecting the `exclude` patterns given (i.e., excluded files will not be
	// included in the hash). This matches the files that are actually copied by the
	// copy operation.
	//
	// ## With bundling
	//
	// When used with the `bundling` property, and the asset is produced by a
	// bundling operation: the hash will still be based on the files in the source
	// path while respecting the `exclude` patterns given. If those given files
	// did not change since the last bundling operation, the hash will remain the same
	// and the bundling operation itself will be skipped. However, do note the following
	// important caveat:
	//
	// - The bundling operation has access to all files in the source path, regardless
	//   of the `exclude` pattern you configure. It is your responsibility to ensure
	//   the `exclude` pattern accurately matches the files the bundling operation
	//   will *actually* access and base its final output on.
	//
	// It is always safe to fingerprint more files than strictly necessary. This will
	// potentially reduce the efficiency of cache hits, but not affect correctness.
	AssetHashType_SOURCE AssetHashType = "SOURCE"
	// Base the asset hash on the contents of the bundling output.
	//
	// Use `OUTPUT` when it is hard to isolate the full set of source files that
	// contribute to the final asset. Instead, the files produced by the `bundling` step
	// will be fingerprinted to produce the asset hash.
	//
	// Use of this value precludes caching of the transform operation: since the
	// asset hash can only be calculated *after* the asset is bundled, the asset must
	// be bundled on every synthesis.
	//
	// This option can only be used in combination with the `bundling` property.
	AssetHashType_OUTPUT AssetHashType = "OUTPUT"
	// Use a custom hash.
	//
	// The `AssetStaging` construct will not calculate the asset hash; instead, it
	// will use the asset hash given in the `assetHash` property.
	AssetHashType_CUSTOM AssetHashType = "CUSTOM"
)

