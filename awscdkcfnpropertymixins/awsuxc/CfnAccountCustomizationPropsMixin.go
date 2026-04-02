package awsuxc

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsuxc/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for managing AWS account-level UX customization settings, including account color, visible services, and visible regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAccountCustomizationPropsMixin := awscdkcfnpropertymixins.Aws_uxc.NewCfnAccountCustomizationPropsMixin(&CfnAccountCustomizationMixinProps{
//   	AccountColor: jsii.String("accountColor"),
//   	VisibleRegions: []*string{
//   		jsii.String("visibleRegions"),
//   	},
//   	VisibleServices: []*string{
//   		jsii.String("visibleServices"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-uxc-accountcustomization.html
//
type CfnAccountCustomizationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAccountCustomizationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccountCustomizationPropsMixin
type jsiiProxy_CfnAccountCustomizationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAccountCustomizationPropsMixin) Props() *CfnAccountCustomizationMixinProps {
	var returns *CfnAccountCustomizationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccountCustomizationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::UXC::AccountCustomization`.
func NewCfnAccountCustomizationPropsMixin(props *CfnAccountCustomizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAccountCustomizationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccountCustomizationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccountCustomizationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_uxc.CfnAccountCustomizationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::UXC::AccountCustomization`.
func NewCfnAccountCustomizationPropsMixin_Override(c CfnAccountCustomizationPropsMixin, props *CfnAccountCustomizationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_uxc.CfnAccountCustomizationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAccountCustomizationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccountCustomizationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_uxc.CfnAccountCustomizationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccountCustomizationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_uxc.CfnAccountCustomizationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccountCustomizationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAccountCustomizationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

