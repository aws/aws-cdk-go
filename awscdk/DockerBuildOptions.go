package awscdk


// Docker build options.
//
// Example:
//   lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: lambda.Code_FromAsset(jsii.String("/path/to/handler"), &AssetOptions{
//   		Bundling: &BundlingOptions{
//   			Image: awscdk.DockerImage_FromBuild(jsii.String("/path/to/dir/with/DockerFile"), &DockerBuildOptions{
//   				BuildArgs: map[string]*string{
//   					"ARG1": jsii.String("value1"),
//   				},
//   			}),
//   			Command: []*string{
//   				jsii.String("my"),
//   				jsii.String("cool"),
//   				jsii.String("command"),
//   			},
//   		},
//   	}),
//   	Runtime: lambda.Runtime_PYTHON_3_9(),
//   	Handler: jsii.String("index.handler"),
//   })
//
type DockerBuildOptions struct {
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
	CacheFrom *[]*DockerCacheOption `field:"optional" json:"cacheFrom" yaml:"cacheFrom"`
	// Cache to options to pass to the `docker build` command.
	// Default: - no cache to args are passed.
	//
	CacheTo *DockerCacheOption `field:"optional" json:"cacheTo" yaml:"cacheTo"`
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
}

