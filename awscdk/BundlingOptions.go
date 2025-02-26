package awscdk


// Bundling options.
//
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
type BundlingOptions struct {
	// The Docker image where the command will run.
	Image DockerImage `field:"required" json:"image" yaml:"image"`
	// The access mechanism used to make source files available to the bundling container and to return the bundling output back to the host.
	// Default: - BundlingFileAccess.BIND_MOUNT
	//
	BundlingFileAccess BundlingFileAccess `field:"optional" json:"bundlingFileAccess" yaml:"bundlingFileAccess"`
	// The command to run in the Docker container.
	//
	// Example value: `['npm', 'install']`.
	// See: https://docs.docker.com/engine/reference/run/
	//
	// Default: - run the command defined in the image.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The entrypoint to run in the Docker container.
	//
	// Example value: `['/bin/sh', '-c']`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	// Default: - run the entrypoint defined in the image.
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the Docker container.
	// Default: - no environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Local bundling provider.
	//
	// The provider implements a method `tryBundle()` which should return `true`
	// if local bundling was performed. If `false` is returned, docker bundling
	// will be done.
	// Default: - bundling will only be performed in a Docker container.
	//
	Local ILocalBundling `field:"optional" json:"local" yaml:"local"`
	// Docker [Networking options](https://docs.docker.com/engine/reference/commandline/run/#connect-a-container-to-a-network---network).
	// Default: - no networking options.
	//
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The type of output that this bundling operation is producing.
	// Default: BundlingOutput.AUTO_DISCOVER
	//
	OutputType BundlingOutput `field:"optional" json:"outputType" yaml:"outputType"`
	// Platform to build for. _Requires Docker Buildx_.
	//
	// Specify this property to build images on a specific platform.
	// Default: - no platform specified (the current machine architecture will be used).
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	// Default: - no security options.
	//
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the Docker container.
	//
	// user | user:group | uid | uid:gid | user:gid | uid:group.
	// See: https://docs.docker.com/engine/reference/run/#user
	//
	// Default: - uid:gid of the current user or 1000:1000 on Windows.
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// Additional Docker volumes to mount.
	// Default: - no additional volumes are mounted.
	//
	Volumes *[]*DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Where to mount the specified volumes from.
	// See: https://docs.docker.com/engine/reference/commandline/run/#mount-volumes-from-container---volumes-from
	//
	// Default: - no containers are specified to mount volumes from.
	//
	VolumesFrom *[]*string `field:"optional" json:"volumesFrom" yaml:"volumesFrom"`
	// Working directory inside the Docker container.
	// Default: /asset-input.
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

