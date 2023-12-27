package awsecrassets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for DockerImageAssets.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &DockerImageAssetProps{
//   	Directory: path.join(__dirname, jsii.String("my-image")),
//   	BuildArgs: map[string]*string{
//   		"HTTP_PROXY": jsii.String("http://10.20.30.2:1234"),
//   	},
//   	Invalidation: &DockerImageAssetInvalidationOptions{
//   		BuildArgs: jsii.Boolean(false),
//   	},
//   })
//
type DockerImageAssetProps struct {
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
	CacheFrom *[]*DockerCacheOption `field:"optional" json:"cacheFrom" yaml:"cacheFrom"`
	// Cache to options to pass to the `docker build` command.
	// See: https://docs.docker.com/build/cache/backends/
	//
	// Default: - no cache to options are passed to the build command.
	//
	CacheTo *DockerCacheOption `field:"optional" json:"cacheTo" yaml:"cacheTo"`
	// Path to the Dockerfile (relative to the directory).
	// Default: 'Dockerfile'.
	//
	File *string `field:"optional" json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	// Default: - hash all parameters.
	//
	Invalidation *DockerImageAssetInvalidationOptions `field:"optional" json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	// Default: - no networking mode specified (the default networking mode `NetworkMode.DEFAULT` will be used)
	//
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
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
	Platform Platform `field:"optional" json:"platform" yaml:"platform"`
	// Docker target to build to.
	// Default: - no target.
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The directory where the Dockerfile is stored.
	//
	// Any directory inside with a name that matches the CDK output folder (cdk.out by default) will be excluded from the asset
	Directory *string `field:"required" json:"directory" yaml:"directory"`
}

