package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of the AWS::QuickSight::Brand Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnBrandPropsMixin := awscdkcfnpropertymixins.Aws_quicksight.NewCfnBrandPropsMixin(&CfnBrandMixinProps{
//   	BrandDefinition: &BrandDefinitionProperty{
//   		ApplicationTheme: &ApplicationThemeProperty{
//   			BrandColorPalette: &BrandColorPaletteProperty{
//   				Accent: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Danger: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Dimension: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Info: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Measure: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Primary: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Secondary: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Success: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   				Warning: &PaletteProperty{
//   					Background: jsii.String("background"),
//   					Foreground: jsii.String("foreground"),
//   				},
//   			},
//   			BrandElementStyle: &BrandElementStyleProperty{
//   				NavbarStyle: &NavbarStyleProperty{
//   					ContextualNavbar: &PaletteProperty{
//   						Background: jsii.String("background"),
//   						Foreground: jsii.String("foreground"),
//   					},
//   					GlobalNavbar: &PaletteProperty{
//   						Background: jsii.String("background"),
//   						Foreground: jsii.String("foreground"),
//   					},
//   				},
//   			},
//   		},
//   		BrandName: jsii.String("brandName"),
//   		Description: jsii.String("description"),
//   	},
//   	BrandId: jsii.String("brandId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-brand.html
//
type CfnBrandPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnBrandMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBrandPropsMixin
type jsiiProxy_CfnBrandPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnBrandPropsMixin) Props() *CfnBrandMixinProps {
	var returns *CfnBrandMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBrandPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::Brand`.
func NewCfnBrandPropsMixin(props *CfnBrandMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnBrandPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBrandPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBrandPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnBrandPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::Brand`.
func NewCfnBrandPropsMixin_Override(c CfnBrandPropsMixin, props *CfnBrandMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnBrandPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnBrandPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBrandPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnBrandPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBrandPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnBrandPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBrandPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnBrandPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

