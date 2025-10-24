package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The properties in a container definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var appProtocol AppProtocol
//   var containerImage ContainerImage
//   var credentialSpec CredentialSpec
//   var environmentFile EnvironmentFile
//   var linuxParameters LinuxParameters
//   var logDriver LogDriver
//   var secret Secret
//   var taskDefinition TaskDefinition
//
//   containerDefinitionProps := &ContainerDefinitionProps{
//   	Image: containerImage,
//   	TaskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	ContainerName: jsii.String("containerName"),
//   	Cpu: jsii.Number(123),
//   	CredentialSpecs: []CredentialSpec{
//   		credentialSpec,
//   	},
//   	DisableNetworking: jsii.Boolean(false),
//   	DnsSearchDomains: []*string{
//   		jsii.String("dnsSearchDomains"),
//   	},
//   	DnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	DockerLabels: map[string]*string{
//   		"dockerLabelsKey": jsii.String("dockerLabels"),
//   	},
//   	DockerSecurityOptions: []*string{
//   		jsii.String("dockerSecurityOptions"),
//   	},
//   	EnableRestartPolicy: jsii.Boolean(false),
//   	EntryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	EnvironmentFiles: []EnvironmentFile{
//   		environmentFile,
//   	},
//   	Essential: jsii.Boolean(false),
//   	ExtraHosts: map[string]*string{
//   		"extraHostsKey": jsii.String("extraHosts"),
//   	},
//   	GpuCount: jsii.Number(123),
//   	HealthCheck: &HealthCheck{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//
//   		// the properties below are optional
//   		Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   		Retries: jsii.Number(123),
//   		StartPeriod: cdk.Duration_*Minutes(jsii.Number(30)),
//   		Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	},
//   	Hostname: jsii.String("hostname"),
//   	InferenceAcceleratorResources: []*string{
//   		jsii.String("inferenceAcceleratorResources"),
//   	},
//   	Interactive: jsii.Boolean(false),
//   	LinuxParameters: linuxParameters,
//   	Logging: logDriver,
//   	MemoryLimitMiB: jsii.Number(123),
//   	MemoryReservationMiB: jsii.Number(123),
//   	PortMappings: []PortMapping{
//   		&PortMapping{
//   			ContainerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			AppProtocol: appProtocol,
//   			ContainerPortRange: jsii.String("containerPortRange"),
//   			HostPort: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			Protocol: awscdk.Aws_ecs.Protocol_TCP,
//   		},
//   	},
//   	Privileged: jsii.Boolean(false),
//   	PseudoTerminal: jsii.Boolean(false),
//   	ReadonlyRootFilesystem: jsii.Boolean(false),
//   	RestartAttemptPeriod: cdk.Duration_*Minutes(jsii.Number(30)),
//   	RestartIgnoredExitCodes: []*f64{
//   		jsii.Number(123),
//   	},
//   	Secrets: map[string]Secret{
//   		"secretsKey": secret,
//   	},
//   	StartTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	StopTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	SystemControls: []SystemControl{
//   		&SystemControl{
//   			Namespace: jsii.String("namespace"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Ulimits: []Ulimit{
//   		&Ulimit{
//   			HardLimit: jsii.Number(123),
//   			Name: awscdk.*Aws_ecs.UlimitName_CORE,
//   			SoftLimit: jsii.Number(123),
//   		},
//   	},
//   	User: jsii.String("user"),
//   	VersionConsistency: awscdk.*Aws_ecs.VersionConsistency_ENABLED,
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
type ContainerDefinitionProps struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage.
	Image ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Default: - CMD value built into container image.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	// Default: - id of node associated with ContainerDefinition.
	//
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	// Default: - No minimum CPU units reserved.
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// A list of ARNs in SSM or Amazon S3 to a credential spec (`CredSpec`) file that configures the container for Active Directory authentication.
	//
	// We recommend that you use this parameter instead of the `dockerSecurityOptions`.
	//
	// Currently, only one credential spec is allowed per container definition.
	// Default: - No credential specs.
	//
	CredentialSpecs *[]CredentialSpec `field:"optional" json:"credentialSpecs" yaml:"credentialSpecs"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	// Default: false.
	//
	DisableNetworking *bool `field:"optional" json:"disableNetworking" yaml:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	// Default: - No search domains.
	//
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	// Default: - Default DNS servers.
	//
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// A key/value map of labels to add to the container.
	// Default: - No labels.
	//
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	// Default: - No security labels.
	//
	DockerSecurityOptions *[]*string `field:"optional" json:"dockerSecurityOptions" yaml:"dockerSecurityOptions"`
	// Enable a restart policy for a container.
	//
	// When you set up a restart policy, Amazon ECS can restart the container without needing to replace the task.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-restart-policy.html
	//
	// Default: - false unless `restartIgnoredExitCodes` or `restartAttemptPeriod` is set.
	//
	EnableRestartPolicy *bool `field:"optional" json:"enableRestartPolicy" yaml:"enableRestartPolicy"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	// Default: - Entry point configured in container.
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to the container.
	// Default: - No environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	// Default: - No environment files.
	//
	EnvironmentFiles *[]EnvironmentFile `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	// Default: true.
	//
	Essential *bool `field:"optional" json:"essential" yaml:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	// Default: - No extra hosts.
	//
	ExtraHosts *map[string]*string `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// The number of GPUs assigned to the container.
	// Default: - No GPUs assigned.
	//
	GpuCount *float64 `field:"optional" json:"gpuCount" yaml:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	// Default: - Health check configuration from container.
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The hostname to use for your container.
	// Default: - Automatic hostname.
	//
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The inference accelerators referenced by the container.
	// Default: - No inference accelerators assigned.
	//
	InferenceAcceleratorResources *[]*string `field:"optional" json:"inferenceAcceleratorResources" yaml:"inferenceAcceleratorResources"`
	// When this parameter is true, you can deploy containerized applications that require stdin or a tty to be allocated.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinition.html#cfn-ecs-taskdefinition-containerdefinition-interactive
	//
	// Default: - false.
	//
	Interactive *bool `field:"optional" json:"interactive" yaml:"interactive"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	// Default: - No Linux parameters.
	//
	LinuxParameters LinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration specification for the container.
	// Default: - Containers use the same logging driver that the Docker daemon uses.
	//
	Logging LogDriver `field:"optional" json:"logging" yaml:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	// Default: - No memory limit.
	//
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	// Default: - No memory reserved.
	//
	MemoryReservationMiB *float64 `field:"optional" json:"memoryReservationMiB" yaml:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	// Default: - No ports are mapped.
	//
	PortMappings *[]*PortMapping `field:"optional" json:"portMappings" yaml:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	// Default: false.
	//
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is true, a TTY is allocated.
	//
	// This parameter maps to Tty in the "Create a container section" of the
	// Docker Remote API and the --tty option to `docker run`.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_pseudoterminal
	//
	// Default: - false.
	//
	PseudoTerminal *bool `field:"optional" json:"pseudoTerminal" yaml:"pseudoTerminal"`
	// When this parameter is true, the container is given read-only access to its root file system.
	// Default: false.
	//
	ReadonlyRootFilesystem *bool `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// A period of time that the container must run for before a restart can be attempted.
	//
	// A container can be restarted only once every `restartAttemptPeriod` seconds.
	// If a container isn't able to run for this time period and exits early, it will not be restarted.
	//
	// This property can't be used if `enableRestartPolicy` is set to false.
	//
	// You can set a minimum `restartAttemptPeriod` of 60 seconds and a maximum `restartAttemptPeriod`
	// of 1800 seconds.
	// Default: - Duration.seconds(300) if `enableRestartPolicy` is true, otherwise no period.
	//
	RestartAttemptPeriod awscdk.Duration `field:"optional" json:"restartAttemptPeriod" yaml:"restartAttemptPeriod"`
	// A list of exit codes that Amazon ECS will ignore and not attempt a restart on.
	//
	// This property can't be used if `enableRestartPolicy` is set to false.
	//
	// You can specify a maximum of 50 container exit codes.
	// Default: - No exit codes are ignored.
	//
	RestartIgnoredExitCodes *[]*float64 `field:"optional" json:"restartIgnoredExitCodes" yaml:"restartIgnoredExitCodes"`
	// The secret environment variables to pass to the container.
	// Default: - No secret environment variables.
	//
	Secrets *map[string]Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	// Default: - none.
	//
	StartTimeout awscdk.Duration `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	// Default: - none.
	//
	StopTimeout awscdk.Duration `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	// Default: - No system controls are set.
	//
	SystemControls *[]*SystemControl `field:"optional" json:"systemControls" yaml:"systemControls"`
	// An array of ulimits to set in the container.
	Ulimits *[]*Ulimit `field:"optional" json:"ulimits" yaml:"ulimits"`
	// The user to use inside the container.
	//
	// This parameter maps to User in the Create a container section of the Docker Remote API and the --user option to docker run.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#ContainerDefinition-user
	//
	// Default: root.
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// Specifies whether Amazon ECS will resolve the container image tag provided in the container definition to an image digest.
	//
	// If you set the value for a container as disabled, Amazon ECS will
	// not resolve the provided container image tag to a digest and will use the
	// original image URI specified in the container definition for deployment.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinition.html#cfn-ecs-taskdefinition-containerdefinition-versionconsistency
	//
	// Default: VersionConsistency.DISABLED if `image` is a CDK asset, VersionConsistency.ENABLED otherwise
	//
	VersionConsistency VersionConsistency `field:"optional" json:"versionConsistency" yaml:"versionConsistency"`
	// The working directory in which to run commands inside the container.
	// Default: /.
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
	// The name of the task definition that includes this container definition.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
}

