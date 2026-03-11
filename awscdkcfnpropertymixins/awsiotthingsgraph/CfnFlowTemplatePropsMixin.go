package awsiotthingsgraph

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiotthingsgraph/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnFlowTemplatePropsMixin := awscdkcfnpropertymixins.Aws_iotthingsgraph.NewCfnFlowTemplatePropsMixin(&CfnFlowTemplateMixinProps{
//   	CompatibleNamespaceVersion: jsii.Number(123),
//   	Definition: &DefinitionDocumentProperty{
//   		Language: jsii.String("language"),
//   		Text: jsii.String("text"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotthingsgraph-flowtemplate.html
//
type CfnFlowTemplatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFlowTemplateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFlowTemplatePropsMixin
type jsiiProxy_CfnFlowTemplatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFlowTemplatePropsMixin) Props() *CfnFlowTemplateMixinProps {
	var returns *CfnFlowTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowTemplatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTThingsGraph::FlowTemplate`.
func NewCfnFlowTemplatePropsMixin(props *CfnFlowTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFlowTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFlowTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlowTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotthingsgraph.CfnFlowTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTThingsGraph::FlowTemplate`.
func NewCfnFlowTemplatePropsMixin_Override(c CfnFlowTemplatePropsMixin, props *CfnFlowTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotthingsgraph.CfnFlowTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFlowTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iotthingsgraph.CfnFlowTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iotthingsgraph.CfnFlowTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlowTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFlowTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

