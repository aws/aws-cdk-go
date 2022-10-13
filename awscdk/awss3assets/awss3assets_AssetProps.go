package awss3assets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Example:
//   asset := assets.NewAsset(this, jsii.String("BundledAsset"), &assetProps{
//   	path: path.join(__dirname, jsii.String("markdown-asset")),
//   	 // /asset-input and working directory in the container
//   	bundling: &bundlingOptions{
//   		image: awscdk.DockerImage.fromBuild(path.join(__dirname, jsii.String("alpine-markdown"))),
//   		 // Build an image
//   		command: []*string{
//   			jsii.String("sh"),
//   			jsii.String("-c"),
//   			jsii.String("\n            markdown index.md > /asset-output/index.html\n          "),
//   		},
//   	},
//   })
//
type AssetProps struct {
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
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Specifies the type of hash to calculate for this asset.
	//
	// If `assetHash` is configured, this option must be `undefined` or
	// `AssetHashType.CUSTOM`.
	AssetHashType awscdk.AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Bundle the asset by executing a command in a Docker container or a custom bundling provider.
	//
	// The asset path will be mounted at `/asset-input`. The Docker
	// container is responsible for putting content at `/asset-output`.
	// The content at `/asset-output` will be zipped and used as the
	// final asset.
	Bundling *awscdk.BundlingOptions `field:"optional" json:"bundling" yaml:"bundling"`
	// File paths matching the patterns will be excluded.
	//
	// See `ignoreMode` to set the matching behavior.
	// Has no effect on Assets bundled using the `bundling` property.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for `exclude` patterns.
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// A list of principals that should be able to read this asset from S3.
	//
	// You can use `asset.grantRead(principal)` to grant read permissions later.
	Readers *[]awsiam.IGrantable `field:"optional" json:"readers" yaml:"readers"`
	// The disk location of the asset.
	//
	// The path should refer to one of the following:
	// - A regular file or a .zip file, in which case the file will be uploaded as-is to S3.
	// - A directory, in which case it will be archived into a .zip file and uploaded to S3.
	Path *string `field:"required" json:"path" yaml:"path"`
}

