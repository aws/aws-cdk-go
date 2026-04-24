package awsecs


// A container definition for a daemon task.
//
// Daemon container definitions describe the containers that run as part of a daemon task on container instances managed by capacity providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   daemonContainerDefinitionProperty := &DaemonContainerDefinitionProperty{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Cpu: jsii.Number(123),
//   	DependsOn: []interface{}{
//   		&ContainerDependencyProperty{
//   			Condition: jsii.String("condition"),
//   			ContainerName: jsii.String("containerName"),
//   		},
//   	},
//   	EntryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	Environment: []interface{}{
//   		&KeyValuePairProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	EnvironmentFiles: []interface{}{
//   		&EnvironmentFileProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Essential: jsii.Boolean(false),
//   	FirelensConfiguration: &FirelensConfigurationProperty{
//   		Options: map[string]*string{
//   			"optionsKey": jsii.String("options"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	HealthCheck: &HealthCheckProperty{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		Interval: jsii.Number(123),
//   		Retries: jsii.Number(123),
//   		StartPeriod: jsii.Number(123),
//   		Timeout: jsii.Number(123),
//   	},
//   	Image: jsii.String("image"),
//   	Interactive: jsii.Boolean(false),
//   	LinuxParameters: &LinuxParametersProperty{
//   		Capabilities: &KernelCapabilitiesProperty{
//   			Add: []*string{
//   				jsii.String("add"),
//   			},
//   			Drop: []*string{
//   				jsii.String("drop"),
//   			},
//   		},
//   		Devices: []interface{}{
//   			&DeviceProperty{
//   				ContainerPath: jsii.String("containerPath"),
//   				HostPath: jsii.String("hostPath"),
//   				Permissions: []*string{
//   					jsii.String("permissions"),
//   				},
//   			},
//   		},
//   		InitProcessEnabled: jsii.Boolean(false),
//   		Tmpfs: []interface{}{
//   			&TmpfsProperty{
//   				ContainerPath: jsii.String("containerPath"),
//   				MountOptions: []*string{
//   					jsii.String("mountOptions"),
//   				},
//   				Size: jsii.Number(123),
//   			},
//   		},
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		LogDriver: jsii.String("logDriver"),
//   		Options: map[string]*string{
//   			"optionsKey": jsii.String("options"),
//   		},
//   		SecretOptions: []interface{}{
//   			&SecretProperty{
//   				Name: jsii.String("name"),
//   				ValueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   	Memory: jsii.Number(123),
//   	MemoryReservation: jsii.Number(123),
//   	MountPoints: []interface{}{
//   		&MountPointProperty{
//   			ContainerPath: jsii.String("containerPath"),
//   			ReadOnly: jsii.Boolean(false),
//   			SourceVolume: jsii.String("sourceVolume"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Privileged: jsii.Boolean(false),
//   	PseudoTerminal: jsii.Boolean(false),
//   	ReadonlyRootFilesystem: jsii.Boolean(false),
//   	RepositoryCredentials: &RepositoryCredentialsProperty{
//   		CredentialsParameter: jsii.String("credentialsParameter"),
//   	},
//   	RestartPolicy: &RestartPolicyProperty{
//   		Enabled: jsii.Boolean(false),
//   		IgnoredExitCodes: []interface{}{
//   			jsii.Number(123),
//   		},
//   		RestartAttemptPeriod: jsii.Number(123),
//   	},
//   	Secrets: []interface{}{
//   		&SecretProperty{
//   			Name: jsii.String("name"),
//   			ValueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   	StartTimeout: jsii.Number(123),
//   	StopTimeout: jsii.Number(123),
//   	SystemControls: []interface{}{
//   		&SystemControlProperty{
//   			Namespace: jsii.String("namespace"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Ulimits: []interface{}{
//   		&UlimitProperty{
//   			HardLimit: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			SoftLimit: jsii.Number(123),
//   		},
//   	},
//   	User: jsii.String("user"),
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html
//
type CfnDaemonTaskDefinitionPropsMixin_DaemonContainerDefinitionProperty struct {
	// The command that's passed to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of ``cpu`` units reserved for the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-cpu
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The dependencies defined for container startup and shutdown.
	//
	// A container can contain multiple dependencies on other containers in a task definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-dependson
	//
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The entry point that's passed to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// A list of files containing the environment variables to pass to a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-environmentfiles
	//
	EnvironmentFiles interface{} `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// If the ``essential`` parameter of a container is marked as ``true``, and that container fails or stops for any reason, all other containers that are part of the task are stopped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-essential
	//
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// The FireLens configuration for the container.
	//
	// This is used to specify and configure a log router for container logs. For more information, see [Custom log routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-firelensconfiguration
	//
	FirelensConfiguration interface{} `field:"optional" json:"firelensConfiguration" yaml:"firelensConfiguration"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The image used to start the container.
	//
	// This string is passed directly to the Docker daemon. Images in the Docker Hub registry are available by default. Other repositories are specified with either ``repository-url/image:tag`` or ``repository-url/image@digest``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-image
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
	// When this parameter is ``true``, you can deploy containerized applications that require ``stdin`` or a ``tty`` to be allocated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-interactive
	//
	Interactive interface{} `field:"optional" json:"interactive" yaml:"interactive"`
	// The Linux-specific options that are applied to the container, such as Linux [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-linuxparameters
	//
	LinuxParameters interface{} `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration for the container.
	//
	// This parameter maps to ``LogConfig`` in the docker container create command and the ``--log-driver`` option to docker run.
	//  By default, containers use the same logging driver that the Docker daemon uses. However, the container might use a different logging driver than the Docker daemon by specifying a log driver configuration in the container definition.
	//  Understand the following when specifying a log configuration for your containers.
	//   +  Amazon ECS currently supports a subset of the logging drivers available to the Docker daemon. Additional log drivers may be available in future releases of the Amazon ECS container agent.
	//  For tasks on FARGATElong, the supported log drivers are ``awslogs``, ``splunk``, and ``awsfirelens``.
	//  For tasks hosted on Amazon EC2 instances, the supported log drivers are ``awslogs``, ``fluentd``, ``gelf``, ``json-file``, ``journald``,``syslog``, ``splunk``, and ``awsfirelens``.
	//   +  This parameter requires version 1.18 of the Docker Remote API or greater on your container instance.
	//   +  For tasks that are hosted on Amazon EC2 instances, the Amazon ECS container agent must register the available logging drivers with the ``ECS_AVAILABLE_LOGGING_DRIVERS`` environment variable before containers placed on that instance can use these log configuration options. For more information, see [Amazon ECS container agent configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the *Amazon Elastic Container Service Developer Guide*.
	//   +  For tasks that are on FARGATElong, because you don't have access to the underlying infrastructure your tasks are hosted on, any additional software needed must be installed outside of the task. For example, the Fluentd output aggregators or a remote host running Logstash to send Gelf logs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If the container attempts to exceed the memory specified here, the container is killed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-memory
	//
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The soft limit (in MiB) of memory to reserve for the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-memoryreservation
	//
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// The mount points for data volumes in your container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-mountpoints
	//
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// The name of the container.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the ``root`` user).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-privileged
	//
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is ``true``, a TTY is allocated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-pseudoterminal
	//
	PseudoTerminal interface{} `field:"optional" json:"pseudoTerminal" yaml:"pseudoTerminal"`
	// When this parameter is true, the container is given read-only access to its root file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-readonlyrootfilesystem
	//
	ReadonlyRootFilesystem interface{} `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The repository credentials for private registry authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-repositorycredentials
	//
	RepositoryCredentials interface{} `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-restartpolicy
	//
	RestartPolicy interface{} `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// The secrets to pass to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-secrets
	//
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-starttimeout
	//
	StartTimeout *float64 `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-stoptimeout
	//
	StopTimeout *float64 `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-systemcontrols
	//
	SystemControls interface{} `field:"optional" json:"systemControls" yaml:"systemControls"`
	// A list of ``ulimits`` to set in the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-ulimits
	//
	Ulimits interface{} `field:"optional" json:"ulimits" yaml:"ulimits"`
	// The user to use inside the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-user
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// The working directory to run commands inside the container in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-workingdirectory
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

