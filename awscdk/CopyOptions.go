package awscdk


// Options applied when copying directories.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyOptions := &CopyOptions{
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	Follow: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   }
//
type CopyOptions struct {
	// File paths matching the patterns will be excluded. See `ignoreMode` to set the matching behavior.
	//
	// ## Without bundling
	//
	// When used without the `bundling` property, the asset is just a regular file
	// copy operation. In that case, this controls the files that will be included
	// in the copy. Files matching the `exclude` patterns will not be included in
	// the copy.
	//
	// ## With bundling
	//
	// When used with the `bundling` property, the asset is produced by a
	// bundling operation. In that case, the value of `exclude` is only meaningful
	// if the `assetHashType` is set to `AssetHashType.SOURCE`, and ignored for
	// other asset hash types. The `exclude` patterns will be used to determine
	// which files contribute to the asset, and the bundling command is skipped
	// if the source files have not changed since the last bundling operation.
	//
	// **NOTE:** The bundling operation still has access to all files in the source path,
	// regardless of the `exclude` patterns. It is your responsibility to ensure
	// the `exclude` pattern accurately matches the files the bundling operation
	// will *actually* access and base its final output on.
	//
	// For `OUTPUT` or `CUSTOM` asset hash types, the `exclude` pattern is ignored.
	// Default: - nothing is excluded.
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	//
	// Symlinks may be copied, materialized or rejected. Note that if symlinks are
	// copied they should be relative; absolute symlinks will lose their
	// meaning if the cloud assembly is moved to another machine.
	//
	// ## Without bundling
	//
	// When used without the `bundling` property, the asset is just a regular file
	// copy operation. In that case, symlinks are materialized or blocked with respect
	// to the source path root as they are being copied.
	//
	// ## With bundling
	//
	// When used with the `bundling` property, the asset is produced by a
	// bundling operation. In that case, the bundling command is fully
	// in charge of producing a set of files in an output directory, and
	// the `AssetStaging` construct will never affect symlink materialization.
	//
	// If the `assetHashType` is `AssetHashType.SOURCE`, this flag affects what
	// the source hash is based on, and the recommended value is one of
	// `SymlinkFollowMode.EXTERNAL` or `SymlinkFollowMode.BLOCK_EXTERNAL`.
	//
	// If the `SymlinkFollowMode.BLOCK_EXTERNAL` mode is used:
	//
	// - Symlinks pointing outside the source tree will be blocked, but only
	//   if using `AssetHashType.SOURCE` (not for any of the other asset hash types).
	// - Symlinks pointing outside the bundle output will be blocked.
	// Default: SymlinkFollowMode.EXTERNAL
	//
	Follow SymlinkFollowMode `field:"optional" json:"follow" yaml:"follow"`
	// What syntax the `exclude` patterns are written in.
	// Default: - IgnoreMode.GLOB for file assets, `IgnoreMode.DOCKER` for image assets (requires that the `@aws-cdk/aws-ecr-assets:dockerIgnoreSupport` feature flag is set, which is recommended)
	//
	IgnoreMode IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
}

