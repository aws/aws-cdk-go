// An experiment to bundle the entire CDK into a single module
package awscdk


// Docker run options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
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
//   	SecurityOpt: jsii.String("securityOpt"),
//   	User: jsii.String("user"),
//   	Volumes: []dockerVolume{
//   		&dockerVolume{
//   			ContainerPath: jsii.String("containerPath"),
//   			HostPath: jsii.String("hostPath"),
//
//   			// the properties below are optional
//   			Consistency: monocdk.DockerVolumeConsistency_CONSISTENT,
//   		},
//   	},
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
// Experimental.
type DockerRunOptions struct {
	// The command to run in the container.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The entrypoint to run in the container.
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	// Experimental.
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the container.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
	// Docker volumes to mount.
	// Experimental.
	Volumes *[]*DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Working directory inside the container.
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

