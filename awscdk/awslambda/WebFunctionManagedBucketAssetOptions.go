package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Options for `WebFunctionCode.fromAssetToManagedBucket()`.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc IVpc
//
//
//   lambda.WebFunctionCode_FromAssetToManagedBucket(path.join(__dirname, jsii.String("my-app")), &WebFunctionManagedBucketAssetOptions{
//   	CopyOptions: &WebFunctionCodeCopyOptions{
//   		Vpc: *Vpc,
//   		VpcSubnets: &SubnetSelection{
//   			SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   		},
//   	},
//   })
//
type WebFunctionManagedBucketAssetOptions struct {
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
	// to the `sourcePath` directory as they are being copied.
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
	// - symlinks pointing outside the source tree will be blocked if using `AssetHashType.SOURCE`.
	// - symlinks pointing outside the bundle output will be blocked.
	// Default: SymlinkFollowMode.EXTERNAL
	//
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// What syntax the `exclude` patterns are written in.
	// Default: - IgnoreMode.GLOB for file assets, `IgnoreMode.DOCKER` for image assets (requires that the `@aws-cdk/aws-ecr-assets:dockerIgnoreSupport` feature flag is set, which is recommended)
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
	// A display name for this asset.
	//
	// If supplied, the display name will be used in locations where the asset
	// identifier is printed, like in the CLI progress information. If the same
	// asset is added multiple times, the display name of the first occurrence is
	// used.
	//
	// The default is the construct path of the Asset construct, with respect to
	// the enclosing stack. If the asset is produced by a construct helper
	// function (such as `lambda.Code.fromAsset()`), this will look like
	// `MyFunction/Code`.
	//
	// We use the stack-relative construct path so that in the common case where
	// you have multiple stacks with the same asset, we won't show something like
	// `/MyBetaStack/MyFunction/Code` when you are actually deploying to
	// production.
	// Default: - Stack-relative construct path.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// A list of principals that should be able to read this asset from S3.
	//
	// You can use `asset.grantRead(principal)` to grant read permissions later.
	// Default: - No principals that can read file asset.
	//
	Readers *[]awsiam.IGrantable `field:"optional" json:"readers" yaml:"readers"`
	// The ARN of the KMS key used to encrypt the handler code.
	// Default: - the default server-side encryption with Amazon S3 managed keys(SSE-S3) key will be used.
	//
	SourceKMSKey interfacesawskms.IKeyRef `field:"optional" json:"sourceKMSKey" yaml:"sourceKMSKey"`
	// Configuration for the custom resource that copies the staged asset into the destination bucket.
	//
	// Use this if your organization requires the copy to run in a VPC, or to supply
	// your own execution role.
	// Default: - the copy runs outside a VPC with a role created for it.
	//
	CopyOptions *WebFunctionCodeCopyOptions `field:"optional" json:"copyOptions" yaml:"copyOptions"`
}

