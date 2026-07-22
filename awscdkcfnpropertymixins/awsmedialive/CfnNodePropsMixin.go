package awsmedialive

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::MediaLive::Node Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNodePropsMixin := awscdkcfnpropertymixins.Aws_medialive.NewCfnNodePropsMixin(&CfnNodeMixinProps{
//   	ClusterId: jsii.String("clusterId"),
//   	Name: jsii.String("name"),
//   	NodeInterfaceMappings: []interface{}{
//   		&NodeInterfaceMappingProperty{
//   			LogicalInterfaceName: jsii.String("logicalInterfaceName"),
//   			NetworkInterfaceMode: jsii.String("networkInterfaceMode"),
//   			PhysicalInterfaceName: jsii.String("physicalInterfaceName"),
//   		},
//   	},
//   	Role: jsii.String("role"),
//   	SdiSourceMappings: []interface{}{
//   		&SdiSourceMappingProperty{
//   			CardNumber: jsii.Number(123),
//   			ChannelNumber: jsii.Number(123),
//   			SdiSource: jsii.String("sdiSource"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-node.html
//
type CfnNodePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNodeMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNodePropsMixin
type jsiiProxy_CfnNodePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNodePropsMixin) Props() *CfnNodeMixinProps {
	var returns *CfnNodeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNodePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaLive::Node`.
func NewCfnNodePropsMixin(props *CfnNodeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNodePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNodePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNodePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_medialive.CfnNodePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaLive::Node`.
func NewCfnNodePropsMixin_Override(c CfnNodePropsMixin, props *CfnNodeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_medialive.CfnNodePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNodePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNodePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_medialive.CfnNodePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNodePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_medialive.CfnNodePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNodePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNodePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

