package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options when creating an asset from a Docker build.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerBuildAssetOptions := &DockerBuildAssetOptions{
//   	BuildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
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
//   	File: jsii.String("file"),
//   	ImagePath: jsii.String("imagePath"),
//   	OutputPath: jsii.String("outputPath"),
//   	Platform: jsii.String("platform"),
//   	TargetStage: jsii.String("targetStage"),
//   }
//
type DockerBuildAssetOptions struct {
	// Build args.
	// Default: - no build args.
	//
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Disable the cache and pass `--no-cache` to the `docker build` command.
	// Default: - cache is used.
	//
	CacheDisabled *bool `field:"optional" json:"cacheDisabled" yaml:"cacheDisabled"`
	// Cache from options to pass to the `docker build` command.
	// Default: - no cache from args are passed.
	//
	CacheFrom *[]*awscdk.DockerCacheOption `field:"optional" json:"cacheFrom" yaml:"cacheFrom"`
	// Cache to options to pass to the `docker build` command.
	// Default: - no cache to args are passed.
	//
	CacheTo *awscdk.DockerCacheOption `field:"optional" json:"cacheTo" yaml:"cacheTo"`
	// Name of the Dockerfile, must relative to the docker build path.
	// Default: `Dockerfile`.
	//
	File *string `field:"optional" json:"file" yaml:"file"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	// Default: - no platform specified.
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Set build target for multi-stage container builds. Any stage defined afterwards will be ignored.
	//
	// Example value: `build-env`.
	// Default: - Build all stages defined in the Dockerfile.
	//
	TargetStage *string `field:"optional" json:"targetStage" yaml:"targetStage"`
	// The path in the Docker image where the asset is located after the build operation.
	// Default: /asset.
	//
	ImagePath *string `field:"optional" json:"imagePath" yaml:"imagePath"`
	// The path on the local filesystem where the asset will be copied using `docker cp`.
	// Default: - a unique temporary directory in the system temp directory.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
}

