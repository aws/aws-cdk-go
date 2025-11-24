package mixinsawsgamelift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsgamelift/mixinsawsgamelift/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The properties that describe a container group resource.
//
// You can update all properties of a container group definition properties. Updates to a container group definition are saved as new versions.
//
// *Used with:* [CreateContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateContainerGroupDefinition.html)
//
// *Returned by:* [DescribeContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeContainerGroupDefinition.html) , [ListContainerGroupDefinitions](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListContainerGroupDefinitions.html) , [UpdateContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateContainerGroupDefinition.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContainerGroupDefinitionPropsMixin := awscdkmixinspreview.Mixins.NewCfnContainerGroupDefinitionPropsMixin(&CfnContainerGroupDefinitionMixinProps{
//   	ContainerGroupType: jsii.String("containerGroupType"),
//   	GameServerContainerDefinition: &GameServerContainerDefinitionProperty{
//   		ContainerName: jsii.String("containerName"),
//   		DependsOn: []interface{}{
//   			&ContainerDependencyProperty{
//   				Condition: jsii.String("condition"),
//   				ContainerName: jsii.String("containerName"),
//   			},
//   		},
//   		EnvironmentOverride: []interface{}{
//   			&ContainerEnvironmentProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ImageUri: jsii.String("imageUri"),
//   		MountPoints: []interface{}{
//   			&ContainerMountPointProperty{
//   				AccessLevel: jsii.String("accessLevel"),
//   				ContainerPath: jsii.String("containerPath"),
//   				InstancePath: jsii.String("instancePath"),
//   			},
//   		},
//   		PortConfiguration: &PortConfigurationProperty{
//   			ContainerPortRanges: []interface{}{
//   				&ContainerPortRangeProperty{
//   					FromPort: jsii.Number(123),
//   					Protocol: jsii.String("protocol"),
//   					ToPort: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ResolvedImageDigest: jsii.String("resolvedImageDigest"),
//   		ServerSdkVersion: jsii.String("serverSdkVersion"),
//   	},
//   	Name: jsii.String("name"),
//   	OperatingSystem: jsii.String("operatingSystem"),
//   	SourceVersionNumber: jsii.Number(123),
//   	SupportContainerDefinitions: []interface{}{
//   		&SupportContainerDefinitionProperty{
//   			ContainerName: jsii.String("containerName"),
//   			DependsOn: []interface{}{
//   				&ContainerDependencyProperty{
//   					Condition: jsii.String("condition"),
//   					ContainerName: jsii.String("containerName"),
//   				},
//   			},
//   			EnvironmentOverride: []interface{}{
//   				&ContainerEnvironmentProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Essential: jsii.Boolean(false),
//   			HealthCheck: &ContainerHealthCheckProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Interval: jsii.Number(123),
//   				Retries: jsii.Number(123),
//   				StartPeriod: jsii.Number(123),
//   				Timeout: jsii.Number(123),
//   			},
//   			ImageUri: jsii.String("imageUri"),
//   			MemoryHardLimitMebibytes: jsii.Number(123),
//   			MountPoints: []interface{}{
//   				&ContainerMountPointProperty{
//   					AccessLevel: jsii.String("accessLevel"),
//   					ContainerPath: jsii.String("containerPath"),
//   					InstancePath: jsii.String("instancePath"),
//   				},
//   			},
//   			PortConfiguration: &PortConfigurationProperty{
//   				ContainerPortRanges: []interface{}{
//   					&ContainerPortRangeProperty{
//   						FromPort: jsii.Number(123),
//   						Protocol: jsii.String("protocol"),
//   						ToPort: jsii.Number(123),
//   					},
//   				},
//   			},
//   			ResolvedImageDigest: jsii.String("resolvedImageDigest"),
//   			Vcpu: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TotalMemoryLimitMebibytes: jsii.Number(123),
//   	TotalVcpuLimit: jsii.Number(123),
//   	VersionDescription: jsii.String("versionDescription"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containergroupdefinition.html
//
type CfnContainerGroupDefinitionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnContainerGroupDefinitionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContainerGroupDefinitionPropsMixin
type jsiiProxy_CfnContainerGroupDefinitionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnContainerGroupDefinitionPropsMixin) Props() *CfnContainerGroupDefinitionMixinProps {
	var returns *CfnContainerGroupDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerGroupDefinitionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GameLift::ContainerGroupDefinition`.
func NewCfnContainerGroupDefinitionPropsMixin(props *CfnContainerGroupDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnContainerGroupDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContainerGroupDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContainerGroupDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GameLift::ContainerGroupDefinition`.
func NewCfnContainerGroupDefinitionPropsMixin_Override(c CfnContainerGroupDefinitionPropsMixin, props *CfnContainerGroupDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnContainerGroupDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerGroupDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainerGroupDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerGroupDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainerGroupDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnContainerGroupDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

