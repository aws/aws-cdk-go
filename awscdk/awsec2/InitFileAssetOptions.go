package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Additional options for creating an InitFile from an asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var grantable iGrantable
//   var initServiceRestartHandle initServiceRestartHandle
//   var key key
//   var localBundling iLocalBundling
//
//   initFileAssetOptions := &InitFileAssetOptions{
//   	AssetHash: jsii.String("assetHash"),
//   	AssetHashType: cdk.AssetHashType_SOURCE,
//   	Base64Encoded: jsii.Boolean(false),
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
//   		Volumes: []dockerVolume{
//   			&dockerVolume{
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
//   	DeployTime: jsii.Boolean(false),
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	FollowSymlinks: cdk.SymlinkFollowMode_NEVER,
//   	Group: jsii.String("group"),
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   	Mode: jsii.String("mode"),
//   	Owner: jsii.String("owner"),
//   	Readers: []*iGrantable{
//   		grantable,
//   	},
//   	ServiceRestartHandles: []*initServiceRestartHandle{
//   		initServiceRestartHandle,
//   	},
//   	SourceKMSKey: key,
//   }
//
type InitFileAssetOptions struct {
	// True if the inlined content (from a string or file) should be treated as base64 encoded.
	//
	// Only applicable for inlined string and file content.
	// Default: false.
	//
	Base64Encoded *bool `field:"optional" json:"base64Encoded" yaml:"base64Encoded"`
	// The name of the owning group for this file.
	//
	// Not supported for Windows systems.
	// Default: 'root'.
	//
	Group *string `field:"optional" json:"group" yaml:"group"`
	// A six-digit octal value representing the mode for this file.
	//
	// Use the first three digits for symlinks and the last three digits for
	// setting permissions. To create a symlink, specify 120xxx, where xxx
	// defines the permissions of the target file. To specify permissions for a
	// file, use the last three digits, such as 000644.
	//
	// Not supported for Windows systems.
	// Default: '000644'.
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The name of the owning user for this file.
	//
	// Not supported for Windows systems.
	// Default: 'root'.
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Restart the given service after this file has been written.
	// Default: - Do not restart any service.
	//
	ServiceRestartHandles *[]InitServiceRestartHandle `field:"optional" json:"serviceRestartHandles" yaml:"serviceRestartHandles"`
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
	// The ARN of the KMS key used to encrypt the handler code.
	// Default: - the default server-side encryption with Amazon S3 managed keys(SSE-S3) key will be used.
	//
	SourceKMSKey awskms.IKey `field:"optional" json:"sourceKMSKey" yaml:"sourceKMSKey"`
}

