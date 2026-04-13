package awsecs


// Container definition for daemon task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   daemonContainerDefinitionProperty := &DaemonContainerDefinitionProperty{
//   	Image: jsii.String("image"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
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
//   				Size: jsii.Number(123),
//
//   				// the properties below are optional
//   				ContainerPath: jsii.String("containerPath"),
//   				MountOptions: []*string{
//   					jsii.String("mountOptions"),
//   				},
//   			},
//   		},
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		LogDriver: jsii.String("logDriver"),
//
//   		// the properties below are optional
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
type CfnDaemonTaskDefinition_DaemonContainerDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-image
	//
	Image *string `field:"required" json:"image" yaml:"image"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-cpu
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-dependson
	//
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-environmentfiles
	//
	EnvironmentFiles interface{} `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-essential
	//
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-firelensconfiguration
	//
	FirelensConfiguration interface{} `field:"optional" json:"firelensConfiguration" yaml:"firelensConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-interactive
	//
	Interactive interface{} `field:"optional" json:"interactive" yaml:"interactive"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-linuxparameters
	//
	LinuxParameters interface{} `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-memory
	//
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-memoryreservation
	//
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-mountpoints
	//
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-privileged
	//
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-pseudoterminal
	//
	PseudoTerminal interface{} `field:"optional" json:"pseudoTerminal" yaml:"pseudoTerminal"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-readonlyrootfilesystem
	//
	ReadonlyRootFilesystem interface{} `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-repositorycredentials
	//
	RepositoryCredentials interface{} `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-restartpolicy
	//
	RestartPolicy interface{} `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-secrets
	//
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-starttimeout
	//
	StartTimeout *float64 `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-stoptimeout
	//
	StopTimeout *float64 `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-systemcontrols
	//
	SystemControls interface{} `field:"optional" json:"systemControls" yaml:"systemControls"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-ulimits
	//
	Ulimits interface{} `field:"optional" json:"ulimits" yaml:"ulimits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-user
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.html#cfn-ecs-daemontaskdefinition-daemoncontainerdefinition-workingdirectory
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

