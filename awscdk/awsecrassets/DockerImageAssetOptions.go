package awsecrassets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for DockerImageAsset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkMode networkMode
//   var platform platform
//
//   dockerImageAssetOptions := &DockerImageAssetOptions{
//   	AssetName: jsii.String("assetName"),
//   	BuildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	BuildSecrets: map[string]*string{
//   		"buildSecretsKey": jsii.String("buildSecrets"),
//   	},
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
//   }
//
type DockerImageAssetOptions struct {
	// File paths matching the patterns will be excluded.
	//
	// See `ignoreMode` to set the matching behavior.
	// Has no effect on Assets bundled using the `bundling` property.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for `exclude` patterns.
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Unique identifier of the docker image asset and its potential revisions.
	//
	// Required if using AppScopedStagingSynthesizer.
	AssetName *string `field:"optional" json:"assetName" yaml:"assetName"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Build secrets.
	//
	// Docker BuildKit must be enabled to use build secrets.
	//
	// Example:
	//   import "github.com/aws/aws-cdk-go/awscdk"
	//
	//
	//   buildSecrets := map[string]*string{
	//   	"MY_SECRET": awscdk.DockerBuildSecret_fromSrc(jsii.String("file.txt")),
	//   }
	//
	// See: https://docs.docker.com/build/buildkit/
	//
	BuildSecrets *map[string]*string `field:"optional" json:"buildSecrets" yaml:"buildSecrets"`
	// Cache from options to pass to the `docker build` command.
	// See: https://docs.docker.com/build/cache/backends/
	//
	CacheFrom *[]*DockerCacheOption `field:"optional" json:"cacheFrom" yaml:"cacheFrom"`
	// Cache to options to pass to the `docker build` command.
	// See: https://docs.docker.com/build/cache/backends/
	//
	CacheTo *DockerCacheOption `field:"optional" json:"cacheTo" yaml:"cacheTo"`
	// Path to the Dockerfile (relative to the directory).
	File *string `field:"optional" json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	Invalidation *DockerImageAssetInvalidationOptions `field:"optional" json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Outputs to pass to the `docker build` command.
	// See: https://docs.docker.com/engine/reference/commandline/build/#custom-build-outputs
	//
	Outputs *[]*string `field:"optional" json:"outputs" yaml:"outputs"`
	// Platform to build for.
	//
	// _Requires Docker Buildx_.
	Platform Platform `field:"optional" json:"platform" yaml:"platform"`
	// Docker target to build to.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

