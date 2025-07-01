package cloudassemblyschema


// Properties for how to produce a Docker image from a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageSource := &DockerImageSource{
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
//   	Directory: jsii.String("directory"),
//   	DockerBuildArgs: map[string]*string{
//   		"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   	},
//   	DockerBuildSecrets: map[string]*string{
//   		"dockerBuildSecretsKey": jsii.String("dockerBuildSecrets"),
//   	},
//   	DockerBuildSsh: jsii.String("dockerBuildSsh"),
//   	DockerBuildTarget: jsii.String("dockerBuildTarget"),
//   	DockerFile: jsii.String("dockerFile"),
//   	DockerOutputs: []*string{
//   		jsii.String("dockerOutputs"),
//   	},
//   	Executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	NetworkMode: jsii.String("networkMode"),
//   	Platform: jsii.String("platform"),
//   }
//
type DockerImageSource struct {
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
	// The directory containing the Docker image build instructions.
	//
	// This path is relative to the asset manifest location.
	// Default: - Exactly one of `directory` and `executable` is required.
	//
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Additional build arguments.
	//
	// Only allowed when `directory` is set.
	// Default: - No additional build arguments.
	//
	DockerBuildArgs *map[string]*string `field:"optional" json:"dockerBuildArgs" yaml:"dockerBuildArgs"`
	// Additional build secrets.
	//
	// Only allowed when `directory` is set.
	// Default: - No additional build secrets.
	//
	DockerBuildSecrets *map[string]*string `field:"optional" json:"dockerBuildSecrets" yaml:"dockerBuildSecrets"`
	// SSH agent socket or keys.
	//
	// Requires building with docker buildkit.
	// Default: - No ssh flag is set.
	//
	DockerBuildSsh *string `field:"optional" json:"dockerBuildSsh" yaml:"dockerBuildSsh"`
	// Target build stage in a Dockerfile with multiple build stages.
	//
	// Only allowed when `directory` is set.
	// Default: - The last stage in the Dockerfile.
	//
	DockerBuildTarget *string `field:"optional" json:"dockerBuildTarget" yaml:"dockerBuildTarget"`
	// The name of the file with build instructions.
	//
	// Only allowed when `directory` is set.
	// Default: "Dockerfile".
	//
	DockerFile *string `field:"optional" json:"dockerFile" yaml:"dockerFile"`
	// Outputs.
	// See: https://docs.docker.com/engine/reference/commandline/build/#custom-build-outputs
	//
	// Default: - no outputs are passed to the build command (default outputs are used).
	//
	DockerOutputs *[]*string `field:"optional" json:"dockerOutputs" yaml:"dockerOutputs"`
	// A command-line executable that returns the name of a local Docker image on stdout after being run.
	// Default: - Exactly one of `directory` and `executable` is required.
	//
	Executable *[]*string `field:"optional" json:"executable" yaml:"executable"`
	// Networking mode for the RUN commands during build. _Requires Docker Engine API v1.25+_.
	//
	// Specify this property to build images on a specific networking mode.
	// Default: - no networking mode specified.
	//
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Platform to build for. _Requires Docker Buildx_.
	//
	// Specify this property to build images on a specific platform/architecture.
	// Default: - current machine platform.
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

