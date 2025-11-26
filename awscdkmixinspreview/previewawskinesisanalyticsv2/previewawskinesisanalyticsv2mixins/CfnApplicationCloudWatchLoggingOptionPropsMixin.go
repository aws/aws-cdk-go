package previewawskinesisanalyticsv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawskinesisanalyticsv2/previewawskinesisanalyticsv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds an Amazon CloudWatch log stream to monitor application configuration errors.
//
// > Only one *ApplicationCloudWatchLoggingOption* resource can be attached per application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationCloudWatchLoggingOptionPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationCloudWatchLoggingOptionPropsMixin(&CfnApplicationCloudWatchLoggingOptionMixinProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	CloudWatchLoggingOption: &CloudWatchLoggingOptionProperty{
//   		LogStreamArn: jsii.String("logStreamArn"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html
//
type CfnApplicationCloudWatchLoggingOptionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationCloudWatchLoggingOptionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationCloudWatchLoggingOptionPropsMixin
type jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin) Props() *CfnApplicationCloudWatchLoggingOptionMixinProps {
	var returns *CfnApplicationCloudWatchLoggingOptionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption`.
func NewCfnApplicationCloudWatchLoggingOptionPropsMixin(props *CfnApplicationCloudWatchLoggingOptionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationCloudWatchLoggingOptionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationCloudWatchLoggingOptionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationCloudWatchLoggingOptionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption`.
func NewCfnApplicationCloudWatchLoggingOptionPropsMixin_Override(c CfnApplicationCloudWatchLoggingOptionPropsMixin, props *CfnApplicationCloudWatchLoggingOptionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationCloudWatchLoggingOptionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationCloudWatchLoggingOptionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationCloudWatchLoggingOptionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationCloudWatchLoggingOptionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationCloudWatchLoggingOptionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_kinesisanalyticsv2.mixins.CfnApplicationCloudWatchLoggingOptionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

