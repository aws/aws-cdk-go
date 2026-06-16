package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDaemonTaskDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDaemonTaskDefinitionMixinProps := &CfnDaemonTaskDefinitionMixinProps{
//   	ContainerDefinitions: []interface{}{
//   		&DaemonContainerDefinitionProperty{
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: jsii.Number(123),
//   			DependsOn: []interface{}{
//   				&ContainerDependencyProperty{
//   					Condition: jsii.String("condition"),
//   					ContainerName: jsii.String("containerName"),
//   				},
//   			},
//   			EntryPoint: []*string{
//   				jsii.String("entryPoint"),
//   			},
//   			Environment: []interface{}{
//   				&KeyValuePairProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			EnvironmentFiles: []interface{}{
//   				&EnvironmentFileProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Essential: jsii.Boolean(false),
//   			FirelensConfiguration: &FirelensConfigurationProperty{
//   				Options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			HealthCheck: &HealthCheckProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Interval: jsii.Number(123),
//   				Retries: jsii.Number(123),
//   				StartPeriod: jsii.Number(123),
//   				Timeout: jsii.Number(123),
//   			},
//   			Image: jsii.String("image"),
//   			Interactive: jsii.Boolean(false),
//   			LinuxParameters: &LinuxParametersProperty{
//   				Capabilities: &KernelCapabilitiesProperty{
//   					Add: []*string{
//   						jsii.String("add"),
//   					},
//   					Drop: []*string{
//   						jsii.String("drop"),
//   					},
//   				},
//   				Devices: []interface{}{
//   					&DeviceProperty{
//   						ContainerPath: jsii.String("containerPath"),
//   						HostPath: jsii.String("hostPath"),
//   						Permissions: []*string{
//   							jsii.String("permissions"),
//   						},
//   					},
//   				},
//   				InitProcessEnabled: jsii.Boolean(false),
//   				Tmpfs: []interface{}{
//   					&TmpfsProperty{
//   						ContainerPath: jsii.String("containerPath"),
//   						MountOptions: []*string{
//   							jsii.String("mountOptions"),
//   						},
//   						Size: jsii.Number(123),
//   					},
//   				},
//   			},
//   			LogConfiguration: &LogConfigurationProperty{
//   				LogDriver: jsii.String("logDriver"),
//   				Options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				SecretOptions: []interface{}{
//   					&SecretProperty{
//   						Name: jsii.String("name"),
//   						ValueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   			},
//   			Memory: jsii.Number(123),
//   			MemoryReservation: jsii.Number(123),
//   			MountPoints: []interface{}{
//   				&MountPointProperty{
//   					ContainerPath: jsii.String("containerPath"),
//   					ReadOnly: jsii.Boolean(false),
//   					SourceVolume: jsii.String("sourceVolume"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Privileged: jsii.Boolean(false),
//   			PseudoTerminal: jsii.Boolean(false),
//   			ReadonlyRootFilesystem: jsii.Boolean(false),
//   			RepositoryCredentials: &RepositoryCredentialsProperty{
//   				CredentialsParameter: jsii.String("credentialsParameter"),
//   			},
//   			RestartPolicy: &RestartPolicyProperty{
//   				Enabled: jsii.Boolean(false),
//   				IgnoredExitCodes: []interface{}{
//   					jsii.Number(123),
//   				},
//   				RestartAttemptPeriod: jsii.Number(123),
//   			},
//   			Secrets: []interface{}{
//   				&SecretProperty{
//   					Name: jsii.String("name"),
//   					ValueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   			StartTimeout: jsii.Number(123),
//   			StopTimeout: jsii.Number(123),
//   			SystemControls: []interface{}{
//   				&SystemControlProperty{
//   					Namespace: jsii.String("namespace"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Ulimits: []interface{}{
//   				&UlimitProperty{
//   					HardLimit: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					SoftLimit: jsii.Number(123),
//   				},
//   			},
//   			User: jsii.String("user"),
//   			WorkingDirectory: jsii.String("workingDirectory"),
//   		},
//   	},
//   	Cpu: jsii.String("cpu"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Family: jsii.String("family"),
//   	IpcMode: jsii.String("ipcMode"),
//   	Memory: jsii.String("memory"),
//   	PidMode: jsii.String("pidMode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   	Volumes: []interface{}{
//   		&VolumeProperty{
//   			Host: &HostVolumePropertiesProperty{
//   				SourcePath: jsii.String("sourcePath"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html
//
type CfnDaemonTaskDefinitionMixinProps struct {
	// A list of container definitions in JSON format that describe the containers that make up the daemon task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-containerdefinitions
	//
	ContainerDefinitions interface{} `field:"optional" json:"containerDefinitions" yaml:"containerDefinitions"`
	// The number of CPU units used by the daemon task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-cpu
	//
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make Amazon Web Services API calls on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of a family that this daemon task definition is registered to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-family
	//
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The IPC namespace mode for the daemon.
	//
	// The valid values are ``none`` and ``shared``. The default is ``none``.
	//  If ``none`` is specified or no value is provided, the daemon runs with its own IPC namespace, isolated from other tasks. If ``shared`` is specified, the daemon joins the host IPC namespace, making it accessible to non-daemon tasks that use ``ipcMode: "host"`` or other daemons that use ``ipcMode: "shared"``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-ipcmode
	//
	IpcMode *string `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// The amount of memory (in MiB) used by the daemon task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-memory
	//
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// The PID namespace mode for the daemon.
	//
	// The valid values are ``none`` and ``shared``. The default is ``none``.
	//  If ``none`` is specified or no value is provided, the daemon runs with its own PID namespace, isolated from other tasks. If ``shared`` is specified, the daemon joins the host PID namespace, making it accessible to non-daemon tasks that use ``pidMode: "host"`` or other daemons that use ``pidMode: "shared"``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-pidmode
	//
	PidMode *string `field:"optional" json:"pidMode" yaml:"pidMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The short name or full Amazon Resource Name (ARN) of the IAM role that grants containers in the daemon task permission to call Amazon Web Services APIs on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-taskrolearn
	//
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
	// The list of data volume definitions for the daemon task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html#cfn-ecs-daemontaskdefinition-volumes
	//
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

