package mixinsawsmedialive

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsmedialive/mixinsawsmedialive/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::MediaLive::CloudWatchAlarmTemplate Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCloudWatchAlarmTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnCloudWatchAlarmTemplatePropsMixin(&CfnCloudWatchAlarmTemplateMixinProps{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	DatapointsToAlarm: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	EvaluationPeriods: jsii.Number(123),
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//   	MetricName: jsii.String("metricName"),
//   	Name: jsii.String("name"),
//   	Period: jsii.Number(123),
//   	Statistic: jsii.String("statistic"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TargetResourceType: jsii.String("targetResourceType"),
//   	Threshold: jsii.Number(123),
//   	TreatMissingData: jsii.String("treatMissingData"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplate.html
//
type CfnCloudWatchAlarmTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCloudWatchAlarmTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCloudWatchAlarmTemplatePropsMixin
type jsiiProxy_CfnCloudWatchAlarmTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCloudWatchAlarmTemplatePropsMixin) Props() *CfnCloudWatchAlarmTemplateMixinProps {
	var returns *CfnCloudWatchAlarmTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudWatchAlarmTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaLive::CloudWatchAlarmTemplate`.
func NewCfnCloudWatchAlarmTemplatePropsMixin(props *CfnCloudWatchAlarmTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCloudWatchAlarmTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCloudWatchAlarmTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudWatchAlarmTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnCloudWatchAlarmTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaLive::CloudWatchAlarmTemplate`.
func NewCfnCloudWatchAlarmTemplatePropsMixin_Override(c CfnCloudWatchAlarmTemplatePropsMixin, props *CfnCloudWatchAlarmTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnCloudWatchAlarmTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCloudWatchAlarmTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudWatchAlarmTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnCloudWatchAlarmTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudWatchAlarmTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_medialive.mixins.CfnCloudWatchAlarmTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudWatchAlarmTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCloudWatchAlarmTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

