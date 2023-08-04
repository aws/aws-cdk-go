package awscdk


// Docker run options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerRunOptions := &DockerRunOptions{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	Network: jsii.String("network"),
//   	Platform: jsii.String("platform"),
//   	SecurityOpt: jsii.String("securityOpt"),
//   	User: jsii.String("user"),
//   	Volumes: []dockerVolume{
//   		&dockerVolume{
//   			ContainerPath: jsii.String("containerPath"),
//   			HostPath: jsii.String("hostPath"),
//
//   			// the properties below are optional
//   			Consistency: cdk.DockerVolumeConsistency_CONSISTENT,
//   		},
//   	},
//   	VolumesFrom: []*string{
//   		jsii.String("volumesFrom"),
//   	},
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
type DockerRunOptions struct {
	// The command to run in the container.
	// Default: - run the command defined in the image.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The entrypoint to run in the container.
	// Default: - run the entrypoint defined in the image.
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the container.
	// Default: - no environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docker [Networking options](https://docs.docker.com/engine/reference/commandline/run/#connect-a-container-to-a-network---network).
	// Default: - no networking options.
	//
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	// Default: - no platform specified.
	//
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	// Default: - no security options.
	//
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the container.
	// Default: - root or image default.
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// Docker volumes to mount.
	// Default: - no volumes are mounted.
	//
	Volumes *[]*DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Where to mount the specified volumes from.
	// See: https://docs.docker.com/engine/reference/commandline/run/#mount-volumes-from-container---volumes-from
	//
	// Default: - no containers are specified to mount volumes from.
	//
	VolumesFrom *[]*string `field:"optional" json:"volumesFrom" yaml:"volumesFrom"`
	// Working directory inside the container.
	// Default: - image default.
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

