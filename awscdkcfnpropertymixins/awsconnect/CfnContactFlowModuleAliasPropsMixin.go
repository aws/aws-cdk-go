package awsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for ContactFlowModuleAlias.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnContactFlowModuleAliasPropsMixin := awscdkcfnpropertymixins.Aws_connect.NewCfnContactFlowModuleAliasPropsMixin(&CfnContactFlowModuleAliasMixinProps{
//   	ContactFlowModuleId: jsii.String("contactFlowModuleId"),
//   	ContactFlowModuleVersion: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-contactflowmodulealias.html
//
type CfnContactFlowModuleAliasPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnContactFlowModuleAliasMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContactFlowModuleAliasPropsMixin
type jsiiProxy_CfnContactFlowModuleAliasPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnContactFlowModuleAliasPropsMixin) Props() *CfnContactFlowModuleAliasMixinProps {
	var returns *CfnContactFlowModuleAliasMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContactFlowModuleAliasPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::ContactFlowModuleAlias`.
func NewCfnContactFlowModuleAliasPropsMixin(props *CfnContactFlowModuleAliasMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnContactFlowModuleAliasPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContactFlowModuleAliasPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContactFlowModuleAliasPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnContactFlowModuleAliasPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::ContactFlowModuleAlias`.
func NewCfnContactFlowModuleAliasPropsMixin_Override(c CfnContactFlowModuleAliasPropsMixin, props *CfnContactFlowModuleAliasMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnContactFlowModuleAliasPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnContactFlowModuleAliasPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContactFlowModuleAliasPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnContactFlowModuleAliasPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContactFlowModuleAliasPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnContactFlowModuleAliasPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContactFlowModuleAliasPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnContactFlowModuleAliasPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

