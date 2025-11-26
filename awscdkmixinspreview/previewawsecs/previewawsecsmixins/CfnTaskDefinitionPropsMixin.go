package previewawsecsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsecs/previewawsecsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Registers a new task definition from the supplied `family` and `containerDefinitions` .
//
// Optionally, you can add data volumes to your containers with the `volumes` parameter. For more information about task definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// You can specify a role for your task with the `taskRoleArn` parameter. When you specify a role for a task, its containers can then use the latest versions of the AWS CLI or SDKs to make API requests to the AWS services that are specified in the policy that's associated with the role. For more information, see [IAM Roles for Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// You can specify a Docker networking mode for the containers in your task definition with the `networkMode` parameter. If you specify the `awsvpc` network mode, the task is allocated an elastic network interface, and you must specify a [NetworkConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_NetworkConfiguration.html) when you create a service or run a task with the task definition. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTaskDefinitionPropsMixin := awscdkmixinspreview.Mixins.NewCfnTaskDefinitionPropsMixin(&CfnTaskDefinitionMixinProps{
//   	ContainerDefinitions: []interface{}{
//   		&ContainerDefinitionProperty{
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: jsii.Number(123),
//   			CredentialSpecs: []*string{
//   				jsii.String("credentialSpecs"),
//   			},
//   			DependsOn: []interface{}{
//   				&ContainerDependencyProperty{
//   					Condition: jsii.String("condition"),
//   					ContainerName: jsii.String("containerName"),
//   				},
//   			},
//   			DisableNetworking: jsii.Boolean(false),
//   			DnsSearchDomains: []*string{
//   				jsii.String("dnsSearchDomains"),
//   			},
//   			DnsServers: []*string{
//   				jsii.String("dnsServers"),
//   			},
//   			DockerLabels: map[string]*string{
//   				"dockerLabelsKey": jsii.String("dockerLabels"),
//   			},
//   			DockerSecurityOptions: []*string{
//   				jsii.String("dockerSecurityOptions"),
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
//   			ExtraHosts: []interface{}{
//   				&HostEntryProperty{
//   					Hostname: jsii.String("hostname"),
//   					IpAddress: jsii.String("ipAddress"),
//   				},
//   			},
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
//   			Hostname: jsii.String("hostname"),
//   			Image: jsii.String("image"),
//   			Interactive: jsii.Boolean(false),
//   			Links: []*string{
//   				jsii.String("links"),
//   			},
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
//   				MaxSwap: jsii.Number(123),
//   				SharedMemorySize: jsii.Number(123),
//   				Swappiness: jsii.Number(123),
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
//   			PortMappings: []interface{}{
//   				&PortMappingProperty{
//   					AppProtocol: jsii.String("appProtocol"),
//   					ContainerPort: jsii.Number(123),
//   					ContainerPortRange: jsii.String("containerPortRange"),
//   					HostPort: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					Protocol: jsii.String("protocol"),
//   				},
//   			},
//   			Privileged: jsii.Boolean(false),
//   			PseudoTerminal: jsii.Boolean(false),
//   			ReadonlyRootFilesystem: jsii.Boolean(false),
//   			RepositoryCredentials: &RepositoryCredentialsProperty{
//   				CredentialsParameter: jsii.String("credentialsParameter"),
//   			},
//   			ResourceRequirements: []interface{}{
//   				&ResourceRequirementProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
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
//   			VersionConsistency: jsii.String("versionConsistency"),
//   			VolumesFrom: []interface{}{
//   				&VolumeFromProperty{
//   					ReadOnly: jsii.Boolean(false),
//   					SourceContainer: jsii.String("sourceContainer"),
//   				},
//   			},
//   			WorkingDirectory: jsii.String("workingDirectory"),
//   		},
//   	},
//   	Cpu: jsii.String("cpu"),
//   	EnableFaultInjection: jsii.Boolean(false),
//   	EphemeralStorage: &EphemeralStorageProperty{
//   		SizeInGiB: jsii.Number(123),
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Family: jsii.String("family"),
//   	InferenceAccelerators: []interface{}{
//   		&InferenceAcceleratorProperty{
//   			DeviceName: jsii.String("deviceName"),
//   			DeviceType: jsii.String("deviceType"),
//   		},
//   	},
//   	IpcMode: jsii.String("ipcMode"),
//   	Memory: jsii.String("memory"),
//   	NetworkMode: jsii.String("networkMode"),
//   	PidMode: jsii.String("pidMode"),
//   	PlacementConstraints: []interface{}{
//   		&TaskDefinitionPlacementConstraintProperty{
//   			Expression: jsii.String("expression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ProxyConfiguration: &ProxyConfigurationProperty{
//   		ContainerName: jsii.String("containerName"),
//   		ProxyConfigurationProperties: []interface{}{
//   			&KeyValuePairProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	RequiresCompatibilities: []*string{
//   		jsii.String("requiresCompatibilities"),
//   	},
//   	RuntimePlatform: &RuntimePlatformProperty{
//   		CpuArchitecture: jsii.String("cpuArchitecture"),
//   		OperatingSystemFamily: jsii.String("operatingSystemFamily"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   	Volumes: []interface{}{
//   		&VolumeProperty{
//   			ConfiguredAtLaunch: jsii.Boolean(false),
//   			DockerVolumeConfiguration: &DockerVolumeConfigurationProperty{
//   				Autoprovision: jsii.Boolean(false),
//   				Driver: jsii.String("driver"),
//   				DriverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				Labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   				Scope: jsii.String("scope"),
//   			},
//   			EfsVolumeConfiguration: &EFSVolumeConfigurationProperty{
//   				AuthorizationConfig: &AuthorizationConfigProperty{
//   					AccessPointId: jsii.String("accessPointId"),
//   					Iam: jsii.String("iam"),
//   				},
//   				FilesystemId: jsii.String("filesystemId"),
//   				RootDirectory: jsii.String("rootDirectory"),
//   				TransitEncryption: jsii.String("transitEncryption"),
//   				TransitEncryptionPort: jsii.Number(123),
//   			},
//   			FSxWindowsFileServerVolumeConfiguration: &FSxWindowsFileServerVolumeConfigurationProperty{
//   				AuthorizationConfig: &FSxAuthorizationConfigProperty{
//   					CredentialsParameter: jsii.String("credentialsParameter"),
//   					Domain: jsii.String("domain"),
//   				},
//   				FileSystemId: jsii.String("fileSystemId"),
//   				RootDirectory: jsii.String("rootDirectory"),
//   			},
//   			Host: &HostVolumePropertiesProperty{
//   				SourcePath: jsii.String("sourcePath"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html
//
type CfnTaskDefinitionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTaskDefinitionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTaskDefinitionPropsMixin
type jsiiProxy_CfnTaskDefinitionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTaskDefinitionPropsMixin) Props() *CfnTaskDefinitionMixinProps {
	var returns *CfnTaskDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinitionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECS::TaskDefinition`.
func NewCfnTaskDefinitionPropsMixin(props *CfnTaskDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTaskDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTaskDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTaskDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECS::TaskDefinition`.
func NewCfnTaskDefinitionPropsMixin_Override(c CfnTaskDefinitionPropsMixin, props *CfnTaskDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTaskDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTaskDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTaskDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTaskDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

