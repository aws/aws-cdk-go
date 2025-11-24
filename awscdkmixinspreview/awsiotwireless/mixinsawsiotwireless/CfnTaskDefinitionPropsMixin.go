package mixinsawsiotwireless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiotwireless/mixinsawsiotwireless/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a gateway task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTaskDefinitionPropsMixin := awscdkmixinspreview.Mixins.NewCfnTaskDefinitionPropsMixin(&CfnTaskDefinitionMixinProps{
//   	AutoCreateTasks: jsii.Boolean(false),
//   	LoRaWanUpdateGatewayTaskEntry: &LoRaWANUpdateGatewayTaskEntryProperty{
//   		CurrentVersion: &LoRaWANGatewayVersionProperty{
//   			Model: jsii.String("model"),
//   			PackageVersion: jsii.String("packageVersion"),
//   			Station: jsii.String("station"),
//   		},
//   		UpdateVersion: &LoRaWANGatewayVersionProperty{
//   			Model: jsii.String("model"),
//   			PackageVersion: jsii.String("packageVersion"),
//   			Station: jsii.String("station"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskDefinitionType: jsii.String("taskDefinitionType"),
//   	Update: &UpdateWirelessGatewayTaskCreateProperty{
//   		LoRaWan: &LoRaWANUpdateGatewayTaskCreateProperty{
//   			CurrentVersion: &LoRaWANGatewayVersionProperty{
//   				Model: jsii.String("model"),
//   				PackageVersion: jsii.String("packageVersion"),
//   				Station: jsii.String("station"),
//   			},
//   			SigKeyCrc: jsii.Number(123),
//   			UpdateSignature: jsii.String("updateSignature"),
//   			UpdateVersion: &LoRaWANGatewayVersionProperty{
//   				Model: jsii.String("model"),
//   				PackageVersion: jsii.String("packageVersion"),
//   				Station: jsii.String("station"),
//   			},
//   		},
//   		UpdateDataRole: jsii.String("updateDataRole"),
//   		UpdateDataSource: jsii.String("updateDataSource"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-taskdefinition.html
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


// Create a mixin to apply properties to `AWS::IoTWireless::TaskDefinition`.
func NewCfnTaskDefinitionPropsMixin(props *CfnTaskDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTaskDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTaskDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTaskDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnTaskDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTWireless::TaskDefinition`.
func NewCfnTaskDefinitionPropsMixin_Override(c CfnTaskDefinitionPropsMixin, props *CfnTaskDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnTaskDefinitionPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnTaskDefinitionPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_iotwireless.mixins.CfnTaskDefinitionPropsMixin",
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

