// An experiment to bundle the entire CDK into a single module
package awscdk


// Initialization properties for `AssetStaging`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var localBundling iLocalBundling
//
//   assetStagingProps := &assetStagingProps{
//   	sourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	assetHash: jsii.String("assetHash"),
//   	assetHashType: monocdk.assetHashType_SOURCE,
//   	bundling: &bundlingOptions{
//   		image: dockerImage,
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		local: localBundling,
//   		outputType: monocdk.bundlingOutput_ARCHIVED,
//   		securityOpt: jsii.String("securityOpt"),
//   		user: jsii.String("user"),
//   		volumes: []dockerVolume{
//   			&dockerVolume{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				consistency: monocdk.dockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	follow: monocdk.symlinkFollowMode_NEVER,
//   	ignoreMode: monocdk.ignoreMode_GLOB,
//   }
//
// Experimental.
type AssetStagingProps struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Experimental.
	Follow SymlinkFollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Experimental.
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
	// Experimental.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Specifies the type of hash to calculate for this asset.
	//
	// If `assetHash` is configured, this option must be `undefined` or
	// `AssetHashType.CUSTOM`.
	// Experimental.
	AssetHashType AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Bundle the asset by executing a command in a Docker container or a custom bundling provider.
	//
	// The asset path will be mounted at `/asset-input`. The Docker
	// container is responsible for putting content at `/asset-output`.
	// The content at `/asset-output` will be zipped and used as the
	// final asset.
	// Experimental.
	Bundling *BundlingOptions `field:"optional" json:"bundling" yaml:"bundling"`
	// The source file or directory to copy from.
	// Experimental.
	SourcePath *string `field:"required" json:"sourcePath" yaml:"sourcePath"`
}

