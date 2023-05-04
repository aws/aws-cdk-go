// Version 2 of the AWS Cloud Development Kit library
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
	BundlingFileAccess BundlingFileAccess `field:"optional" json:"bundlingFileAccess" yaml:"bundlingFileAccess"`
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

