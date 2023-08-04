package awss3assets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewAsset(this, jsii.String("BundledAsset"), &AssetProps{
//   	Path: jsii.String("/path/to/asset"),
//   	Bundling: &BundlingOptions{
//   		Image: cdk.DockerImage_FromRegistry(jsii.String("alpine")),
//   		Command: []*string{
//   			jsii.String("command-that-produces-an-archive.sh"),
//   		},
//   		OutputType: cdk.BundlingOutput_NOT_ARCHIVED,
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
	AssetHashType awscdk.AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Bundle the asset by executing a command in a Docker container or a custom bundling provider.
	//
	// The asset path will be mounted at `/asset-input`. The Docker
	// container is responsible for putting content at `/asset-output`.
	// The content at `/asset-output` will be zipped and used as the
	// final asset.
	// Default: - uploaded as-is to S3 if the asset is a regular file or a .zip file,
	// archived into a .zip file and uploaded to S3 otherwise
	//
	Bundling *awscdk.BundlingOptions `field:"optional" json:"bundling" yaml:"bundling"`
	// File paths matching the patterns will be excluded.
	//
	// See `ignoreMode` to set the matching behavior.
	// Has no effect on Assets bundled using the `bundling` property.
	// Default: - nothing is excluded.
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Default: SymlinkFollowMode.NEVER
	//
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for `exclude` patterns.
	// Default: IgnoreMode.GLOB
	//
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Whether or not the asset needs to exist beyond deployment time;
	//
	// i.e.
	// are copied over to a different location and not needed afterwards.
	// Setting this property to true has an impact on the lifecycle of the asset,
	// because we will assume that it is safe to delete after the CloudFormation
	// deployment succeeds.
	//
	// For example, Lambda Function assets are copied over to Lambda during
	// deployment. Therefore, it is not necessary to store the asset in S3, so
	// we consider those deployTime assets.
	// Default: false.
	//
	DeployTime *bool `field:"optional" json:"deployTime" yaml:"deployTime"`
	// A list of principals that should be able to read this asset from S3.
	//
	// You can use `asset.grantRead(principal)` to grant read permissions later.
	// Default: - No principals that can read file asset.
	//
	Readers *[]awsiam.IGrantable `field:"optional" json:"readers" yaml:"readers"`
	// The disk location of the asset.
	//
	// The path should refer to one of the following:
	// - A regular file or a .zip file, in which case the file will be uploaded as-is to S3.
	// - A directory, in which case it will be archived into a .zip file and uploaded to S3.
	Path *string `field:"required" json:"path" yaml:"path"`
}

