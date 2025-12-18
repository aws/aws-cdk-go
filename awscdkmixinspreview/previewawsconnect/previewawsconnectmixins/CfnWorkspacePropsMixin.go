package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconnect/previewawsconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Contains information about a workspace, which defines the user experience by mapping views to pages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkspacePropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkspacePropsMixin(&CfnWorkspaceMixinProps{
//   	Associations: []*string{
//   		jsii.String("associations"),
//   	},
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Media: []interface{}{
//   		&MediaItemProperty{
//   			Source: jsii.String("source"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Pages: []interface{}{
//   		&WorkspacePageProperty{
//   			InputData: jsii.String("inputData"),
//   			Page: jsii.String("page"),
//   			ResourceArn: jsii.String("resourceArn"),
//   			Slug: jsii.String("slug"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Theme: &WorkspaceThemeProperty{
//   		Dark: &WorkspaceThemeConfigProperty{
//   			Palette: &WorkspaceThemePaletteProperty{
//   				Canvas: &PaletteCanvasProperty{
//   					ActiveBackground: jsii.String("activeBackground"),
//   					ContainerBackground: jsii.String("containerBackground"),
//   					PageBackground: jsii.String("pageBackground"),
//   				},
//   				Header: &PaletteHeaderProperty{
//   					Background: jsii.String("background"),
//   					InvertActionsColors: jsii.Boolean(false),
//   					Text: jsii.String("text"),
//   					TextHover: jsii.String("textHover"),
//   				},
//   				Navigation: &PaletteNavigationProperty{
//   					Background: jsii.String("background"),
//   					InvertActionsColors: jsii.Boolean(false),
//   					Text: jsii.String("text"),
//   					TextActive: jsii.String("textActive"),
//   					TextBackgroundActive: jsii.String("textBackgroundActive"),
//   					TextBackgroundHover: jsii.String("textBackgroundHover"),
//   					TextHover: jsii.String("textHover"),
//   				},
//   				Primary: &PalettePrimaryProperty{
//   					Active: jsii.String("active"),
//   					ContrastText: jsii.String("contrastText"),
//   					Default: jsii.String("default"),
//   				},
//   			},
//   			Typography: &WorkspaceThemeTypographyProperty{
//   				FontFamily: &FontFamilyProperty{
//   					Default: jsii.String("default"),
//   				},
//   			},
//   		},
//   		Light: &WorkspaceThemeConfigProperty{
//   			Palette: &WorkspaceThemePaletteProperty{
//   				Canvas: &PaletteCanvasProperty{
//   					ActiveBackground: jsii.String("activeBackground"),
//   					ContainerBackground: jsii.String("containerBackground"),
//   					PageBackground: jsii.String("pageBackground"),
//   				},
//   				Header: &PaletteHeaderProperty{
//   					Background: jsii.String("background"),
//   					InvertActionsColors: jsii.Boolean(false),
//   					Text: jsii.String("text"),
//   					TextHover: jsii.String("textHover"),
//   				},
//   				Navigation: &PaletteNavigationProperty{
//   					Background: jsii.String("background"),
//   					InvertActionsColors: jsii.Boolean(false),
//   					Text: jsii.String("text"),
//   					TextActive: jsii.String("textActive"),
//   					TextBackgroundActive: jsii.String("textBackgroundActive"),
//   					TextBackgroundHover: jsii.String("textBackgroundHover"),
//   					TextHover: jsii.String("textHover"),
//   				},
//   				Primary: &PalettePrimaryProperty{
//   					Active: jsii.String("active"),
//   					ContrastText: jsii.String("contrastText"),
//   					Default: jsii.String("default"),
//   				},
//   			},
//   			Typography: &WorkspaceThemeTypographyProperty{
//   				FontFamily: &FontFamilyProperty{
//   					Default: jsii.String("default"),
//   				},
//   			},
//   		},
//   	},
//   	Title: jsii.String("title"),
//   	Visibility: jsii.String("visibility"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-workspace.html
//
type CfnWorkspacePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkspaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkspacePropsMixin
type jsiiProxy_CfnWorkspacePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkspacePropsMixin) Props() *CfnWorkspaceMixinProps {
	var returns *CfnWorkspaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::Workspace`.
func NewCfnWorkspacePropsMixin(props *CfnWorkspaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkspacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkspacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkspacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::Workspace`.
func NewCfnWorkspacePropsMixin_Override(c CfnWorkspacePropsMixin, props *CfnWorkspaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkspacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnWorkspacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnWorkspacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspacePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWorkspacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

