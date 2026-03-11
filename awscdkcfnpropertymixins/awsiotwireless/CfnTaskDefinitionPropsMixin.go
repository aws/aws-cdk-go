package awsiotwireless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a gateway task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTaskDefinitionPropsMixin := awscdkcfnpropertymixins.Aws_iotwireless.NewCfnTaskDefinitionPropsMixin(&CfnTaskDefinitionMixinProps{
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-taskdefinition.html
//
type CfnTaskDefinitionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTaskDefinitionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTaskDefinitionPropsMixin
type jsiiProxy_CfnTaskDefinitionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnTaskDefinitionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTWireless::TaskDefinition`.
func NewCfnTaskDefinitionPropsMixin(props *CfnTaskDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTaskDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTaskDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTaskDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnTaskDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTWireless::TaskDefinition`.
func NewCfnTaskDefinitionPropsMixin_Override(c CfnTaskDefinitionPropsMixin, props *CfnTaskDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnTaskDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTaskDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTaskDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnTaskDefinitionPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_iotwireless.CfnTaskDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTaskDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

