package awskinesisanalyticsv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awskinesisanalyticsv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds an Amazon CloudWatch log stream to monitor application configuration errors.
//
// > Only one *ApplicationCloudWatchLoggingOption* resource can be attached per application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnApplicationCloudWatchLoggingOptionPropsMixin := awscdkcfnpropertymixins.Aws_kinesisanalyticsv2.NewCfnApplicationCloudWatchLoggingOptionPropsMixin(&CfnApplicationCloudWatchLoggingOptionMixinProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	CloudWatchLoggingOption: &CloudWatchLoggingOptionProperty{
//   		LogStreamArn: jsii.String("logStreamArn"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html
//
type CfnApplicationCloudWatchLoggingOptionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnApplicationCloudWatchLoggingOptionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationCloudWatchLoggingOptionPropsMixin
type jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption`.
func NewCfnApplicationCloudWatchLoggingOptionPropsMixin(props *CfnApplicationCloudWatchLoggingOptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnApplicationCloudWatchLoggingOptionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationCloudWatchLoggingOptionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOptionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption`.
func NewCfnApplicationCloudWatchLoggingOptionPropsMixin_Override(c CfnApplicationCloudWatchLoggingOptionPropsMixin, props *CfnApplicationCloudWatchLoggingOptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOptionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnApplicationCloudWatchLoggingOptionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationCloudWatchLoggingOptionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOptionPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalyticsv2.CfnApplicationCloudWatchLoggingOptionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationCloudWatchLoggingOptionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

