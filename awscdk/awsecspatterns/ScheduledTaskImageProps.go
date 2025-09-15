package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
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
//   scheduledTaskImageProps := &ScheduledTaskImageProps{
//   	Image: containerImage,
//
//   	// the properties below are optional
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	ContainerName: jsii.String("containerName"),
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	LogDriver: logDriver,
//   	Secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   }
//
type ScheduledTaskImageProps struct {
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	// Default: - none.
	//
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Default: - CMD value built into container image.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Optional name for the container added.
	// Default: - ScheduledContainer.
	//
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The environment variables to pass to the container.
	// Default: none.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The log driver to use.
	// Default: - AwsLogDriver if enableLogging is true.
	//
	LogDriver awsecs.LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The secret to expose to the container as an environment variable.
	// Default: - No secret environment variables.
	//
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
}

