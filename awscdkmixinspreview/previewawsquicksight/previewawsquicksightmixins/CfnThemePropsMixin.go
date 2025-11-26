package previewawsquicksightmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsquicksight/previewawsquicksightmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a theme.
//
// A *theme* is set of configuration options for color and layout. Themes apply to analyses and dashboards. For more information, see [Using Themes in Amazon Quick Suite](https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html) in the *Amazon Quick Suite User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnThemePropsMixin := awscdkmixinspreview.Mixins.NewCfnThemePropsMixin(&CfnThemeMixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	BaseThemeId: jsii.String("baseThemeId"),
//   	Configuration: &ThemeConfigurationProperty{
//   		DataColorPalette: &DataColorPaletteProperty{
//   			Colors: []*string{
//   				jsii.String("colors"),
//   			},
//   			EmptyFillColor: jsii.String("emptyFillColor"),
//   			MinMaxGradient: []*string{
//   				jsii.String("minMaxGradient"),
//   			},
//   		},
//   		Sheet: &SheetStyleProperty{
//   			Tile: &TileStyleProperty{
//   				Border: &BorderStyleProperty{
//   					Show: jsii.Boolean(false),
//   				},
//   			},
//   			TileLayout: &TileLayoutStyleProperty{
//   				Gutter: &GutterStyleProperty{
//   					Show: jsii.Boolean(false),
//   				},
//   				Margin: &MarginStyleProperty{
//   					Show: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		Typography: &TypographyProperty{
//   			FontFamilies: []interface{}{
//   				&FontProperty{
//   					FontFamily: jsii.String("fontFamily"),
//   				},
//   			},
//   		},
//   		UiColorPalette: &UIColorPaletteProperty{
//   			Accent: jsii.String("accent"),
//   			AccentForeground: jsii.String("accentForeground"),
//   			Danger: jsii.String("danger"),
//   			DangerForeground: jsii.String("dangerForeground"),
//   			Dimension: jsii.String("dimension"),
//   			DimensionForeground: jsii.String("dimensionForeground"),
//   			Measure: jsii.String("measure"),
//   			MeasureForeground: jsii.String("measureForeground"),
//   			PrimaryBackground: jsii.String("primaryBackground"),
//   			PrimaryForeground: jsii.String("primaryForeground"),
//   			SecondaryBackground: jsii.String("secondaryBackground"),
//   			SecondaryForeground: jsii.String("secondaryForeground"),
//   			Success: jsii.String("success"),
//   			SuccessForeground: jsii.String("successForeground"),
//   			Warning: jsii.String("warning"),
//   			WarningForeground: jsii.String("warningForeground"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   			Resource: jsii.String("resource"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThemeId: jsii.String("themeId"),
//   	VersionDescription: jsii.String("versionDescription"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-theme.html
//
type CfnThemePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnThemeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnThemePropsMixin
type jsiiProxy_CfnThemePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnThemePropsMixin) Props() *CfnThemeMixinProps {
	var returns *CfnThemeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnThemePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::Theme`.
func NewCfnThemePropsMixin(props *CfnThemeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnThemePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnThemePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnThemePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnThemePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::Theme`.
func NewCfnThemePropsMixin_Override(c CfnThemePropsMixin, props *CfnThemeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnThemePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnThemePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnThemePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnThemePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnThemePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnThemePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnThemePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnThemePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

