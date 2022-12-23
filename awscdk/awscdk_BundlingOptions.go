// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Bundling options.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   asset := assets.NewAsset(this, jsii.String("BundledAsset"), &assetProps{
//   	path: jsii.String("/path/to/asset"),
//   	bundling: &bundlingOptions{
//   		image: ambda.runtime_PYTHON_3_9_BundlingImage,
//   		command: []*string{
//   			jsii.String("bash"),
//   			jsii.String("-c"),
//   			jsii.String("pip install -r requirements.txt -t /asset-output && cp -au . /asset-output"),
//   		},
//   		securityOpt: jsii.String("no-new-privileges:true"),
//   		 // https://docs.docker.com/engine/reference/commandline/run/#optional-security-options---security-opt
//   		network: jsii.String("host"),
//   	},
//   })
//
type BundlingOptions struct {
	// The Docker image where the command will run.
	Image DockerImage `field:"required" json:"image" yaml:"image"`
	// The command to run in the Docker container.
	//
	// Example value: `['npm', 'install']`.
	// See: https://docs.docker.com/engine/reference/run/
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The entrypoint to run in the Docker container.
	//
	// Example value: `['/bin/sh', '-c']`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the Docker container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Local bundling provider.
	//
	// The provider implements a method `tryBundle()` which should return `true`
	// if local bundling was performed. If `false` is returned, docker bundling
	// will be done.
	Local ILocalBundling `field:"optional" json:"local" yaml:"local"`
	// Docker [Networking options](https://docs.docker.com/engine/reference/commandline/run/#connect-a-container-to-a-network---network).
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The type of output that this bundling operation is producing.
	OutputType BundlingOutput `field:"optional" json:"outputType" yaml:"outputType"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the Docker container.
	//
	// user | user:group | uid | uid:gid | user:gid | uid:group.
	// See: https://docs.docker.com/engine/reference/run/#user
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// Additional Docker volumes to mount.
	Volumes *[]*DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Where to mount the specified volumes from.
	// See: https://docs.docker.com/engine/reference/commandline/run/#mount-volumes-from-container---volumes-from
	//
	VolumesFrom *[]*string `field:"optional" json:"volumesFrom" yaml:"volumesFrom"`
	// Working directory inside the Docker container.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

