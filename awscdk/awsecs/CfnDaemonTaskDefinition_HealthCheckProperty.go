package awsecs


// An object representing a container health check.
//
// Health check parameters that are specified in a container definition override any Docker health checks that exist in the container image (such as those specified in a parent image or from the image's Dockerfile). This configuration maps to the ``HEALTHCHECK`` parameter of docker run.
//   The Amazon ECS container agent only monitors and reports on the health checks specified in the task definition. Amazon ECS does not monitor Docker health checks that are embedded in a container image and not specified in the container definition. Health check parameters that are specified in a container definition override any Docker health checks that exist in the container image.
//   You can view the health status of both individual containers and a task with the DescribeTasks API operation or when viewing the task details in the console.
//  The health check is designed to make sure that your containers survive agent restarts, upgrades, or temporary unavailability.
//  Amazon ECS performs health checks on containers with the default that launched the container instance or the task.
//  The following describes the possible ``healthStatus`` values for a container:
//   +  ``HEALTHY``-The container health check has passed successfully.
//   +  ``UNHEALTHY``-The container health check has failed.
//   +  ``UNKNOWN``-The container health check is being evaluated, there's no container health check defined, or Amazon ECS doesn't have the health status of the container.
//
//  The following describes the possible ``healthStatus`` values based on the container health checker status of essential containers in the task with the following priority order (high to low):
//   +  ``UNHEALTHY``-One or more essential containers have failed their health check.
//   +  ``UNKNOWN``-Any essential container running within the task is in an ``UNKNOWN`` state and no other essential containers have an ``UNHEALTHY`` state.
//   +  ``HEALTHY``-All essential containers within the task have passed their health checks.
//
//  Consider the following task health example with 2 containers.
//   +  If Container1 is ``UNHEALTHY`` and Container2 is ``UNKNOWN``, the task health is ``UNHEALTHY``.
//   +  If Container1 is ``UNHEALTHY`` and Container2 is ``HEALTHY``, the task health is ``UNHEALTHY``.
//   +  If Container1 is ``HEALTHY`` and Container2 is ``UNKNOWN``, the task health is ``UNKNOWN``.
//   +  If Container1 is ``HEALTHY`` and Container2 is ``HEALTHY``, the task health is ``HEALTHY``.
//
//  Consider the following task health example with 3 containers.
//   +  If Container1 is ``UNHEALTHY`` and Container2 is ``UNKNOWN``, and Container3 is ``UNKNOWN``, the task health is ``UNHEALTHY``.
//   +  If Container1 is ``UNHEALTHY`` and Container2 is ``UNKNOWN``, and Container3 is ``HEALTHY``, the task health is ``UNHEALTHY``.
//   +  If Container1 is ``UNHEALTHY`` and Container2 is ``HEALTHY``, and Container3 is ``HEALTHY``, the task health is ``UNHEALTHY``.
//   +  If Container1 is ``HEALTHY`` and Container2 is ``UNKNOWN``, and Container3 is ``HEALTHY``, the task health is ``UNKNOWN``.
//   +  If Container1 is ``HEALTHY`` and Container2 is ``UNKNOWN``, and Container3 is ``UNKNOWN``, the task health is ``UNKNOWN``.
//   +  If Container1 is ``HEALTHY`` and Container2 is ``HEALTHY``, and Container3 is ``HEALTHY``, the task health is ``HEALTHY``.
//
//  If a task is run manually, and not as part of a service, the task will continue its lifecycle regardless of its health status. For tasks that are part of a service, if the task reports as unhealthy then the task will be stopped and the service scheduler will replace it.
//  When a container health check fails for a task that is part of a service, the following process occurs:
//   1.  The task is marked as ``UNHEALTHY``.
//   1.  The unhealthy task will be stopped, and during the stopping process, it will go through the following states:
//   +  ``DEACTIVATING`` - In this state, Amazon ECS performs additional steps before stopping the task. For example, for tasks that are part of services configured to use Elastic Load Balancing target groups, target groups will be deregistered in this state.
//   +  ``STOPPING`` - The task is in the process of being stopped.
//   +  ``DEPROVISIONING`` - Resources associated with the task are being cleaned up.
//   +  ``STOPPED`` - The task has been completely stopped.
//
//   1.  After the old task stops, a new task will be launched to ensure service operation, and the new task will go through the following lifecycle:
//   +  ``PROVISIONING`` - Resources required for the task are being provisioned.
//   +  ``PENDING`` - The task is waiting to be placed on a container instance.
//   +  ``ACTIVATING`` - In this state, Amazon ECS pulls container images, creates containers, configures task networking, registers load balancer target groups, and configures service discovery status.
//   +  ``RUNNING`` - The task is running and performing its work.
//
//
//  For more detailed information about task lifecycle states, see [Task lifecycle](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-lifecycle-explanation.html) in the *Amazon Elastic Container Service Developer Guide*.
//  The following are notes about container health check support:
//   +  If the Amazon ECS container agent becomes disconnected from the Amazon ECS service, this won't cause a container to transition to an ``UNHEALTHY`` status. This is by design, to ensure that containers remain running during agent restarts or temporary unavailability. The health check status is the "last heard from" response from the Amazon ECS agent, so if the container was considered ``HEALTHY`` prior to the disconnect, that status will remain until the agent reconnects and another health check occurs. There are no assumptions made about the status of the container health checks.
//   +  Container health checks require version ``1.17.0`` or greater of the Amazon ECS container agent. For more information, see [Updating the Amazon ECS container agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html).
//   +  Container health checks are supported for Fargate tasks if you're using platform version ``1.1.0`` or greater. For more information, see [platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
//   +  Container health checks aren't supported for tasks that are part of a service that's configured to use a Classic Load Balancer.
//
//  For an example of how to specify a task definition with multiple containers where container dependency is specified, see [Container dependency](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/example_task_definitions.html#example_task_definition-containerdependency) in the *Amazon Elastic Container Service Developer Guide*.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html
//
type CfnDaemonTaskDefinition_HealthCheckProperty struct {
	// A string array representing the command that the container runs to determine if it is healthy.
	//
	// The string array must start with ``CMD`` to run the command arguments directly, or ``CMD-SHELL`` to run the command with the container's default shell.
	//   When you use the AWS Management Console JSON panel, the CLIlong, or the APIs, enclose the list of commands in double quotes and brackets.
	//   ``[ "CMD-SHELL", "curl -f http://localhost/ || exit 1" ]``
	//  You don't include the double quotes and brackets when you use the AWS Management Console.
	//   ``CMD-SHELL, curl -f http://localhost/ || exit 1``
	//  An exit code of 0 indicates success, and non-zero exit code indicates failure. For more information, see ``HealthCheck`` in the docker container create command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The time period in seconds between each health check execution.
	//
	// You may specify between 5 and 300 seconds. The default value is 30 seconds. This value applies only when you specify a ``command``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The number of times to retry a failed health check before the container is considered unhealthy.
	//
	// You may specify between 1 and 10 retries. The default value is 3. This value applies only when you specify a ``command``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-retries
	//
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The optional grace period to provide containers time to bootstrap before failed health checks count towards the maximum number of retries.
	//
	// You can specify between 0 and 300 seconds. By default, the ``startPeriod`` is off. This value applies only when you specify a ``command``.
	//   If a health check succeeds within the ``startPeriod``, then the container is considered healthy and any subsequent failures count toward the maximum number of retries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-startperiod
	//
	StartPeriod *float64 `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// The time period in seconds to wait for a health check to succeed before it is considered a failure.
	//
	// You may specify between 2 and 60 seconds. The default value is 5. This value applies only when you specify a ``command``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-healthcheck.html#cfn-ecs-daemontaskdefinition-healthcheck-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

