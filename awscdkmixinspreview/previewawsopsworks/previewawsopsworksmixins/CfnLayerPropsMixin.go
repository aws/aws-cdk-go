package previewawsopsworksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsopsworks/previewawsopsworksmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var customJson interface{}
//
//   cfnLayerPropsMixin := awscdkmixinspreview.Mixins.NewCfnLayerPropsMixin(&CfnLayerMixinProps{
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	AutoAssignElasticIps: jsii.Boolean(false),
//   	AutoAssignPublicIps: jsii.Boolean(false),
//   	CustomInstanceProfileArn: jsii.String("customInstanceProfileArn"),
//   	CustomJson: customJson,
//   	CustomRecipes: &RecipesProperty{
//   		Configure: []*string{
//   			jsii.String("configure"),
//   		},
//   		Deploy: []*string{
//   			jsii.String("deploy"),
//   		},
//   		Setup: []*string{
//   			jsii.String("setup"),
//   		},
//   		Shutdown: []*string{
//   			jsii.String("shutdown"),
//   		},
//   		Undeploy: []*string{
//   			jsii.String("undeploy"),
//   		},
//   	},
//   	CustomSecurityGroupIds: []*string{
//   		jsii.String("customSecurityGroupIds"),
//   	},
//   	EnableAutoHealing: jsii.Boolean(false),
//   	InstallUpdatesOnBoot: jsii.Boolean(false),
//   	LifecycleEventConfiguration: &LifecycleEventConfigurationProperty{
//   		ShutdownEventConfiguration: &ShutdownEventConfigurationProperty{
//   			DelayUntilElbConnectionsDrained: jsii.Boolean(false),
//   			ExecutionTimeout: jsii.Number(123),
//   		},
//   	},
//   	LoadBasedAutoScaling: &LoadBasedAutoScalingProperty{
//   		DownScaling: &AutoScalingThresholdsProperty{
//   			CpuThreshold: jsii.Number(123),
//   			IgnoreMetricsTime: jsii.Number(123),
//   			InstanceCount: jsii.Number(123),
//   			LoadThreshold: jsii.Number(123),
//   			MemoryThreshold: jsii.Number(123),
//   			ThresholdsWaitTime: jsii.Number(123),
//   		},
//   		Enable: jsii.Boolean(false),
//   		UpScaling: &AutoScalingThresholdsProperty{
//   			CpuThreshold: jsii.Number(123),
//   			IgnoreMetricsTime: jsii.Number(123),
//   			InstanceCount: jsii.Number(123),
//   			LoadThreshold: jsii.Number(123),
//   			MemoryThreshold: jsii.Number(123),
//   			ThresholdsWaitTime: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Packages: []*string{
//   		jsii.String("packages"),
//   	},
//   	Shortname: jsii.String("shortname"),
//   	StackId: jsii.String("stackId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	UseEbsOptimizedInstances: jsii.Boolean(false),
//   	VolumeConfigurations: []interface{}{
//   		&VolumeConfigurationProperty{
//   			Encrypted: jsii.Boolean(false),
//   			Iops: jsii.Number(123),
//   			MountPoint: jsii.String("mountPoint"),
//   			NumberOfDisks: jsii.Number(123),
//   			RaidLevel: jsii.Number(123),
//   			Size: jsii.Number(123),
//   			VolumeType: jsii.String("volumeType"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-layer.html
//
type CfnLayerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLayerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLayerPropsMixin
type jsiiProxy_CfnLayerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLayerPropsMixin) Props() *CfnLayerMixinProps {
	var returns *CfnLayerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OpsWorks::Layer`.
func NewCfnLayerPropsMixin(props *CfnLayerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLayerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLayerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLayerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnLayerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OpsWorks::Layer`.
func NewCfnLayerPropsMixin_Override(c CfnLayerPropsMixin, props *CfnLayerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnLayerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLayerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLayerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnLayerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLayerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_opsworks.mixins.CfnLayerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLayerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLayerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

