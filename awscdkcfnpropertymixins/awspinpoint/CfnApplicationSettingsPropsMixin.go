package awspinpoint

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the settings for an Amazon Pinpoint application.
//
// In Amazon Pinpoint, an *application* (also referred to as an *app* or *project* ) is a collection of related settings, customer information, segments, and campaigns, and other types of Amazon Pinpoint resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnApplicationSettingsPropsMixin := awscdkcfnpropertymixins.Aws_pinpoint.NewCfnApplicationSettingsPropsMixin(&CfnApplicationSettingsMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	CampaignHook: &CampaignHookProperty{
//   		LambdaFunctionName: jsii.String("lambdaFunctionName"),
//   		Mode: jsii.String("mode"),
//   		WebUrl: jsii.String("webUrl"),
//   	},
//   	CloudWatchMetricsEnabled: jsii.Boolean(false),
//   	Limits: &LimitsProperty{
//   		Daily: jsii.Number(123),
//   		MaximumDuration: jsii.Number(123),
//   		MessagesPerSecond: jsii.Number(123),
//   		Total: jsii.Number(123),
//   	},
//   	QuietTime: &QuietTimeProperty{
//   		End: jsii.String("end"),
//   		Start: jsii.String("start"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-applicationsettings.html
//
type CfnApplicationSettingsPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnApplicationSettingsMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationSettingsPropsMixin
type jsiiProxy_CfnApplicationSettingsPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnApplicationSettingsPropsMixin) Props() *CfnApplicationSettingsMixinProps {
	var returns *CfnApplicationSettingsMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationSettingsPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Pinpoint::ApplicationSettings`.
func NewCfnApplicationSettingsPropsMixin(props *CfnApplicationSettingsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnApplicationSettingsPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationSettingsPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationSettingsPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_pinpoint.CfnApplicationSettingsPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Pinpoint::ApplicationSettings`.
func NewCfnApplicationSettingsPropsMixin_Override(c CfnApplicationSettingsPropsMixin, props *CfnApplicationSettingsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_pinpoint.CfnApplicationSettingsPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnApplicationSettingsPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationSettingsPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_pinpoint.CfnApplicationSettingsPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationSettingsPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_pinpoint.CfnApplicationSettingsPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationSettingsPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnApplicationSettingsPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

