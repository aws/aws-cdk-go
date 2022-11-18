package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerImage containerImage
//   var logDriver logDriver
//   var secret secret
//
//   scheduledTaskImageProps := &scheduledTaskImageProps{
//   	image: containerImage,
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	logDriver: logDriver,
//   	secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   }
//
// Experimental.
type ScheduledTaskImageProps struct {
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	// Experimental.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The environment variables to pass to the container.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
}

