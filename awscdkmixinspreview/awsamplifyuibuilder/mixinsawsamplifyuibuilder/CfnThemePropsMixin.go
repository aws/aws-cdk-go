package mixinsawsamplifyuibuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsamplifyuibuilder/mixinsawsamplifyuibuilder/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::AmplifyUIBuilder::Theme resource specifies a theme within an Amplify app.
//
// A theme is a collection of style settings that apply globally to the components associated with the app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var themeValuesProperty_ ThemeValuesProperty
//
//   cfnThemePropsMixin := awscdkmixinspreview.Mixins.NewCfnThemePropsMixin(&CfnThemeMixinProps{
//   	AppId: jsii.String("appId"),
//   	EnvironmentName: jsii.String("environmentName"),
//   	Name: jsii.String("name"),
//   	Overrides: []interface{}{
//   		&ThemeValuesProperty{
//   			Key: jsii.String("key"),
//   			Value: &ThemeValueProperty{
//   				Children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Values: []interface{}{
//   		&ThemeValuesProperty{
//   			Key: jsii.String("key"),
//   			Value: &ThemeValueProperty{
//   				Children: []interface{}{
//   					themeValuesProperty_,
//   				},
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amplifyuibuilder-theme.html
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


// Create a mixin to apply properties to `AWS::AmplifyUIBuilder::Theme`.
func NewCfnThemePropsMixin(props *CfnThemeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnThemePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnThemePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnThemePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amplifyuibuilder.mixins.CfnThemePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AmplifyUIBuilder::Theme`.
func NewCfnThemePropsMixin_Override(c CfnThemePropsMixin, props *CfnThemeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amplifyuibuilder.mixins.CfnThemePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_amplifyuibuilder.mixins.CfnThemePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_amplifyuibuilder.mixins.CfnThemePropsMixin",
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

