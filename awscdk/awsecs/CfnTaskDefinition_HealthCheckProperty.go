package awsecs


// The `HealthCheck` property specifies an object representing a container health check.
//
// Health check parameters that are specified in a container definition override any Docker health checks that exist in the container image (such as those specified in a parent image or from the image's Dockerfile). This configuration maps to the `HEALTHCHECK` parameter of [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
//
// > The Amazon ECS container agent only monitors and reports on the health checks specified in the task definition. Amazon ECS does not monitor Docker health checks that are embedded in a container image and not specified in the container definition. Health check parameters that are specified in a container definition override any Docker health checks that exist in the container image.
//
// If a task is run manually, and not as part of a service, the task will continue its lifecycle regardless of its health status. For tasks that are part of a service, if the task reports as unhealthy then the task will be stopped and the service scheduler will replace it.
//
// The following are notes about container health check support:
//
// - Container health checks require version 1.17.0 or greater of the Amazon ECS container agent. For more information, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) .
// - Container health checks are supported for Fargate tasks if you are using platform version 1.1.0 or greater. For more information, see [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) .
// - Container health checks are not supported for tasks that are part of a service that is configured to use a Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckProperty := &HealthCheckProperty{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Interval: jsii.Number(123),
//   	Retries: jsii.Number(123),
//   	StartPeriod: jsii.Number(123),
//   	Timeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-healthcheck.html
//
type CfnTaskDefinition_HealthCheckProperty struct {
	// A string array representing the command that the container runs to determine if it is healthy.
	//
	// The string array must start with `CMD` to run the command arguments directly, or `CMD-SHELL` to run the command with the container's default shell.
	//
	// When you use the AWS Management Console JSON panel, the AWS Command Line Interface , or the APIs, enclose the list of commands in double quotes and brackets.
	//
	// `[ "CMD-SHELL", "curl -f http://localhost/ || exit 1" ]`
	//
	// You don't include the double quotes and brackets when you use the AWS Management Console.
	//
	// `CMD-SHELL, curl -f http://localhost/ || exit 1`
	//
	// An exit code of 0 indicates success, and non-zero exit code indicates failure. For more information, see `HealthCheck` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-healthcheck.html#cfn-ecs-taskdefinition-healthcheck-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The time period in seconds between each health check execution.
	//
	// You may specify between 5 and 300 seconds. The default value is 30 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-healthcheck.html#cfn-ecs-taskdefinition-healthcheck-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The number of times to retry a failed health check before the container is considered unhealthy.
	//
	// You may specify between 1 and 10 retries. The default value is 3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-healthcheck.html#cfn-ecs-taskdefinition-healthcheck-retries
	//
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The optional grace period to provide containers time to bootstrap before failed health checks count towards the maximum number of retries.
	//
	// You can specify between 0 and 300 seconds. By default, the `startPeriod` is off.
	//
	// > If a health check succeeds within the `startPeriod` , then the container is considered healthy and any subsequent failures count toward the maximum number of retries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-healthcheck.html#cfn-ecs-taskdefinition-healthcheck-startperiod
	//
	StartPeriod *float64 `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// The time period in seconds to wait for a health check to succeed before it is considered a failure.
	//
	// You may specify between 2 and 60 seconds. The default value is 5.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-healthcheck.html#cfn-ecs-taskdefinition-healthcheck-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

