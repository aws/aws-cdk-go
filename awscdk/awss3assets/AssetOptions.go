package awss3assets

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/assets"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Example:
//   lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("my-python-handler")), &AssetOptions{
//   		Bundling: &BundlingOptions{
//   			Image: lambda.Runtime_PYTHON_3_9().BundlingImage,
//   			Command: []*string{
//   				jsii.String("bash"),
//   				jsii.String("-c"),
//   				jsii.String("pip install -r requirements.txt -t /asset-output && cp -au . /asset-output"),
//   			},
//   		},
//   	}),
//   	Runtime: lambda.Runtime_PYTHON_3_9(),
//   	Handler: jsii.String("index.handler"),
//   })
//
// Experimental.
type AssetOptions struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead.
	Follow assets.FollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// A strategy for how to handle symlinks.
	// Experimental.
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
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
	AssetHashType awscdk.AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Bundle the asset by executing a command in a Docker container or a custom bundling provider.
	//
	// The asset path will be mounted at `/asset-input`. The Docker
	// container is responsible for putting content at `/asset-output`.
	// The content at `/asset-output` will be zipped and used as the
	// final asset.
	// Experimental.
	Bundling *awscdk.BundlingOptions `field:"optional" json:"bundling" yaml:"bundling"`
	// A list of principals that should be able to read this asset from S3.
	//
	// You can use `asset.grantRead(principal)` to grant read permissions later.
	// Experimental.
	Readers *[]awsiam.IGrantable `field:"optional" json:"readers" yaml:"readers"`
	// Custom hash to use when identifying the specific version of the asset.
	//
	// For consistency,
	// this custom hash will be SHA256 hashed and encoded as hex. The resulting hash will be
	// the asset hash.
	//
	// NOTE: the source hash is used in order to identify a specific revision of the asset,
	// and used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the source hash,
	// you will need to make sure it is updated every time the source changes, or otherwise
	// it is possible that some deployments will not be invalidated.
	// Deprecated: see `assetHash` and `assetHashType`.
	SourceHash *string `field:"optional" json:"sourceHash" yaml:"sourceHash"`
}

