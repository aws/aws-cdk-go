package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
)

// Properties to initialize a new AssetImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkMode networkMode
//   var platform platform
//
//   assetImageCodeProps := &AssetImageCodeProps{
//   	AssetName: jsii.String("assetName"),
//   	BuildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	BuildSecrets: map[string]*string{
//   		"buildSecretsKey": jsii.String("buildSecrets"),
//   	},
//   	BuildSsh: jsii.String("buildSsh"),
//   	CacheDisabled: jsii.Boolean(false),
//   	CacheFrom: []dockerCacheOption{
//   		&dockerCacheOption{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Params: map[string]*string{
//   				"paramsKey": jsii.String("params"),
//   			},
//   		},
//   	},
//   	CacheTo: &dockerCacheOption{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Params: map[string]*string{
//   			"paramsKey": jsii.String("params"),
//   		},
//   	},
//   	Cmd: []*string{
//   		jsii.String("cmd"),
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	ExtraHash: jsii.String("extraHash"),
//   	File: jsii.String("file"),
//   	FollowSymlinks: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   	Invalidation: &DockerImageAssetInvalidationOptions{
//   		BuildArgs: jsii.Boolean(false),
//   		BuildSecrets: jsii.Boolean(false),
//   		BuildSsh: jsii.Boolean(false),
//   		ExtraHash: jsii.Boolean(false),
//   		File: jsii.Boolean(false),
//   		NetworkMode: jsii.Boolean(false),
//   		Outputs: jsii.Boolean(false),
//   		Platform: jsii.Boolean(false),
//   		RepositoryName: jsii.Boolean(false),
//   		Target: jsii.Boolean(false),
//   	},
//   	NetworkMode: networkMode,
//   	Outputs: []*string{
//   		jsii.String("outputs"),
//   	},
//   	Platform: platform,
//   	Target: jsii.String("target"),
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
type AssetImageCodeProps struct {
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
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Default: - hash is only based on source content.
	//
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Unique identifier of the docker image asset and its potential revisions.
	//
	// Required if using AppScopedStagingSynthesizer.
	// Default: - no asset name.
	//
	AssetName *string `field:"optional" json:"assetName" yaml:"assetName"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	// Default: - no build args are passed.
	//
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Build secrets.
	//
	// Docker BuildKit must be enabled to use build secrets.
	//
	// Example:
	//   import { DockerBuildSecret } from 'aws-cdk-lib';
	//
	//   const buildSecrets = {
	//     'MY_SECRET': DockerBuildSecret.fromSrc('file.txt')
	//   };
	//
	// See: https://docs.docker.com/build/buildkit/
	//
	// Default: - no build secrets.
	//
	BuildSecrets *map[string]*string `field:"optional" json:"buildSecrets" yaml:"buildSecrets"`
	// SSH agent socket or keys to pass to the `docker build` command.
	//
	// Docker BuildKit must be enabled to use the ssh flag.
	// See: https://docs.docker.com/build/buildkit/
	//
	// Default: - no --ssh flag.
	//
	BuildSsh *string `field:"optional" json:"buildSsh" yaml:"buildSsh"`
	// Disable the cache and pass `--no-cache` to the `docker build` command.
	// Default: - cache is used.
	//
	CacheDisabled *bool `field:"optional" json:"cacheDisabled" yaml:"cacheDisabled"`
	// Cache from options to pass to the `docker build` command.
	// See: https://docs.docker.com/build/cache/backends/
	//
	// Default: - no cache from options are passed to the build command.
	//
	CacheFrom *[]*awsecrassets.DockerCacheOption `field:"optional" json:"cacheFrom" yaml:"cacheFrom"`
	// Cache to options to pass to the `docker build` command.
	// See: https://docs.docker.com/build/cache/backends/
	//
	// Default: - no cache to options are passed to the build command.
	//
	CacheTo *awsecrassets.DockerCacheOption `field:"optional" json:"cacheTo" yaml:"cacheTo"`
	// A display name for this asset.
	//
	// If supplied, the display name will be used in locations where the asset
	// identifier is printed, like in the CLI progress information. If the same
	// asset is added multiple times, the display name of the first occurrence is
	// used.
	//
	// If `assetName` is given, it will also be used as the default `displayName`.
	// Otherwise, the default is the construct path of the ImageAsset construct,
	// with respect to the enclosing stack. If the asset is produced by a
	// construct helper function (such as `lambda.Code.fromAssetImage()`), this
	// will look like `MyFunction/AssetImage`.
	//
	// We use the stack-relative construct path so that in the common case where
	// you have multiple stacks with the same asset, we won't show something like
	// `/MyBetaStack/MyFunction/Code` when you are actually deploying to
	// production.
	// Default: - Stack-relative construct path.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Path to the Dockerfile (relative to the directory).
	// Default: 'Dockerfile'.
	//
	File *string `field:"optional" json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	// Default: - hash all parameters.
	//
	Invalidation *awsecrassets.DockerImageAssetInvalidationOptions `field:"optional" json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	// Default: - no networking mode specified (the default networking mode `NetworkMode.DEFAULT` will be used)
	//
	NetworkMode awsecrassets.NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Outputs to pass to the `docker build` command.
	// See: https://docs.docker.com/engine/reference/commandline/build/#custom-build-outputs
	//
	// Default: - no outputs are passed to the build command (default outputs are used).
	//
	Outputs *[]*string `field:"optional" json:"outputs" yaml:"outputs"`
	// Platform to build for.
	//
	// _Requires Docker Buildx_.
	// Default: - no platform specified (the current machine architecture will be used).
	//
	Platform awsecrassets.Platform `field:"optional" json:"platform" yaml:"platform"`
	// Docker target to build to.
	// Default: - no target.
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Specify or override the CMD on the specified Docker image or Dockerfile.
	//
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	// Default: - use the CMD specified in the docker image or Dockerfile.
	//
	Cmd *[]*string `field:"optional" json:"cmd" yaml:"cmd"`
	// Specify or override the ENTRYPOINT on the specified Docker image or Dockerfile.
	//
	// An ENTRYPOINT allows you to configure a container that will run as an executable.
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	// Default: - use the ENTRYPOINT in the docker image or Dockerfile.
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Specify or override the WORKDIR on the specified Docker image or Dockerfile.
	//
	// A WORKDIR allows you to configure the working directory the container will use.
	// See: https://docs.docker.com/engine/reference/builder/#workdir
	//
	// Default: - use the WORKDIR in the docker image or Dockerfile.
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

