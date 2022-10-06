// An experiment to bundle the entire CDK into a single module
package awscdk


// Docker run options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerRunOptions := &dockerRunOptions{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	securityOpt: jsii.String("securityOpt"),
//   	user: jsii.String("user"),
//   	volumes: []dockerVolume{
//   		&dockerVolume{
//   			containerPath: jsii.String("containerPath"),
//   			hostPath: jsii.String("hostPath"),
//
//   			// the properties below are optional
//   			consistency: monocdk.dockerVolumeConsistency_CONSISTENT,
//   		},
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
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

