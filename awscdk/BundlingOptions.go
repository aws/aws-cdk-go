package awscdk


// Bundling options.
//
// Example:
//   asset := assets.NewAsset(this, jsii.String("BundledAsset"), &AssetProps{
//   	Path: path.join(__dirname, jsii.String("markdown-asset")),
//   	 // /asset-input and working directory in the container
//   	Bundling: &BundlingOptions{
//   		Image: awscdk.DockerImage_FromBuild(path.join(__dirname, jsii.String("alpine-markdown"))),
//   		 // Build an image
//   		Command: []*string{
//   			jsii.String("sh"),
//   			jsii.String("-c"),
//   			jsii.String(`
//   			            markdown index.md > /asset-output/index.html
//   			          `),
//   		},
//   	},
//   })
//
// Experimental.
type BundlingOptions struct {
	// The Docker image where the command will run.
	// Experimental.
	Image DockerImage `field:"required" json:"image" yaml:"image"`
	// The command to run in the Docker container.
	//
	// Example value: `['npm', 'install']`.
	// See: https://docs.docker.com/engine/reference/run/
	//
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The entrypoint to run in the Docker container.
	//
	// Example value: `['/bin/sh', '-c']`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the Docker container.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Local bundling provider.
	//
	// The provider implements a method `tryBundle()` which should return `true`
	// if local bundling was performed. If `false` is returned, docker bundling
	// will be done.
	// Experimental.
	Local ILocalBundling `field:"optional" json:"local" yaml:"local"`
	// The type of output that this bundling operation is producing.
	// Experimental.
	OutputType BundlingOutput `field:"optional" json:"outputType" yaml:"outputType"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	// Experimental.
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the Docker container.
	//
	// user | user:group | uid | uid:gid | user:gid | uid:group.
	// See: https://docs.docker.com/engine/reference/run/#user
	//
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
	// Additional Docker volumes to mount.
	// Experimental.
	Volumes *[]*DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Working directory inside the Docker container.
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

