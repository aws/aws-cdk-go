package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The health check command and associated configuration parameters for the container.
//
// Example:
//   var vpc vpc
//   var securityGroup securityGroup
//
//   queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
//   	vpc: vpc,
//   	memoryLimitMiB: jsii.Number(512),
//   	image: ecs.containerImage.fromRegistry(jsii.String("test")),
//   	healthCheck: &healthCheck{
//   		command: []*string{
//   			jsii.String("CMD-SHELL"),
//   			jsii.String("curl -f http://localhost/ || exit 1"),
//   		},
//   		// the properties below are optional
//   		interval: awscdk.Duration.minutes(jsii.Number(30)),
//   		retries: jsii.Number(123),
//   		startPeriod: awscdk.Duration.minutes(jsii.Number(30)),
//   		timeout: awscdk.Duration.minutes(jsii.Number(30)),
//   	},
//   })
//
type HealthCheck struct {
	// A string array representing the command that the container runs to determine if it is healthy.
	//
	// The string array must start with CMD to execute the command arguments directly, or
	// CMD-SHELL to run the command with the container's default shell.
	//
	// For example: [ "CMD-SHELL", "curl -f http://localhost/ || exit 1" ].
	Command *[]*string `field:"required" json:"command" yaml:"command"`
	// The time period in seconds between each health check execution.
	//
	// You may specify between 5 and 300 seconds.
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The number of times to retry a failed health check before the container is considered unhealthy.
	//
	// You may specify between 1 and 10 retries.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The optional grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries.
	//
	// You may specify between 0 and 300 seconds.
	StartPeriod awscdk.Duration `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// The time period in seconds to wait for a health check to succeed before it is considered a failure.
	//
	// You may specify between 2 and 60 seconds.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

