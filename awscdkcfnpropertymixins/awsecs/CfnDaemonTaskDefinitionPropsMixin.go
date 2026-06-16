package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The details of a daemon task definition.
//
// A daemon task definition is a template that describes the containers that form a daemon. Daemons deploy cross-cutting software agents independently across your Amazon ECS infrastructure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDaemonTaskDefinitionPropsMixin := awscdkcfnpropertymixins.Aws_ecs.NewCfnDaemonTaskDefinitionPropsMixin(&CfnDaemonTaskDefinitionMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-daemontaskdefinition.html
//
type CfnDaemonTaskDefinitionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDaemonTaskDefinitionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDaemonTaskDefinitionPropsMixin
type jsiiProxy_CfnDaemonTaskDefinitionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDaemonTaskDefinitionPropsMixin) Props() *CfnDaemonTaskDefinitionMixinProps {
	var returns *CfnDaemonTaskDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDaemonTaskDefinitionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECS::DaemonTaskDefinition`.
func NewCfnDaemonTaskDefinitionPropsMixin(props *CfnDaemonTaskDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDaemonTaskDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDaemonTaskDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDaemonTaskDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECS::DaemonTaskDefinition`.
func NewCfnDaemonTaskDefinitionPropsMixin_Override(c CfnDaemonTaskDefinitionPropsMixin, props *CfnDaemonTaskDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDaemonTaskDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDaemonTaskDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDaemonTaskDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnDaemonTaskDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDaemonTaskDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDaemonTaskDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

