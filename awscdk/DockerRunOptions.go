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
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The entrypoint to run in the container.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docker [Networking options](https://docs.docker.com/engine/reference/commandline/run/#connect-a-container-to-a-network---network).
	Network *string `field:"optional" json:"network" yaml:"network"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the container.
	User *string `field:"optional" json:"user" yaml:"user"`
	// Docker volumes to mount.
	Volumes *[]*DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Where to mount the specified volumes from.
	// See: https://docs.docker.com/engine/reference/commandline/run/#mount-volumes-from-container---volumes-from
	//
	VolumesFrom *[]*string `field:"optional" json:"volumesFrom" yaml:"volumesFrom"`
	// Working directory inside the container.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

