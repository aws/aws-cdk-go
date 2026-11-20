package awscdk


// Initialization properties for `AssetStaging`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage DockerImage
//   var localBundling ILocalBundling
//
//   assetStagingProps := &AssetStagingProps{
//   	SourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	AssetHash: jsii.String("assetHash"),
//   	AssetHashType: cdk.AssetHashType_SOURCE,
//   	Bundling: &BundlingOptions{
//   		Image: dockerImage,
//
//   		// the properties below are optional
//   		BundlingFileAccess: cdk.BundlingFileAccess_VOLUME_COPY,
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		Entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		Local: localBundling,
//   		Network: jsii.String("network"),
//   		OutputType: cdk.BundlingOutput_ARCHIVED,
//   		Platform: jsii.String("platform"),
//   		SecurityOpt: jsii.String("securityOpt"),
//   		User: jsii.String("user"),
//   		Volumes: []DockerVolume{
//   			&DockerVolume{
//   				ContainerPath: jsii.String("containerPath"),
//   				HostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				Consistency: cdk.DockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		VolumesFrom: []*string{
//   			jsii.String("volumesFrom"),
//   		},
//   		WorkingDirectory: jsii.String("workingDirectory"),
//   	},
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	ExtraHash: jsii.String("extraHash"),
//   	Follow: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   }
//
type AssetStagingProps struct {
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
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Default: - hash is only based on the file fingerprint.
	//
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Specify a custom hash for this asset.
	//
	// If `assetHashType` is set it must
	// be set to `AssetHashType.CUSTOM`. For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	// Default: - based on `assetHashType`.
	//
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Specifies the type of hash to calculate for this asset.
	//
	// If `assetHash` is configured, this option must be `undefined` or
	// `AssetHashType.CUSTOM`.
	// Default: - the default is `AssetHashType.SOURCE`, but if `assetHash` is
	// explicitly specified this value defaults to `AssetHashType.CUSTOM`.
	//
	AssetHashType AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Bundle the asset by executing a command in a Docker container or a custom bundling provider.
	//
	// The asset path will be mounted at `/asset-input`. The Docker
	// container is responsible for putting content at `/asset-output`.
	// The content at `/asset-output` will be zipped and used as the
	// final asset.
	// Default: - uploaded as-is to S3 if the asset is a regular file or a .zip file,
	// archived into a .zip file and uploaded to S3 otherwise
	//
	Bundling *BundlingOptions `field:"optional" json:"bundling" yaml:"bundling"`
	// The source file or directory to copy from.
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
}

