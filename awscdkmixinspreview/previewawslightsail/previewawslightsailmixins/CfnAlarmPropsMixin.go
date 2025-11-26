package previewawslightsailmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslightsail/previewawslightsailmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lightsail::Alarm` resource specifies an alarm that can be used to monitor a single metric for one of your Lightsail resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAlarmPropsMixin := awscdkmixinspreview.Mixins.NewCfnAlarmPropsMixin(&CfnAlarmMixinProps{
//   	AlarmName: jsii.String("alarmName"),
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	ContactProtocols: []*string{
//   		jsii.String("contactProtocols"),
//   	},
//   	DatapointsToAlarm: jsii.Number(123),
//   	EvaluationPeriods: jsii.Number(123),
//   	MetricName: jsii.String("metricName"),
//   	MonitoredResourceName: jsii.String("monitoredResourceName"),
//   	NotificationEnabled: jsii.Boolean(false),
//   	NotificationTriggers: []*string{
//   		jsii.String("notificationTriggers"),
//   	},
//   	Threshold: jsii.Number(123),
//   	TreatMissingData: jsii.String("treatMissingData"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-alarm.html
//
type CfnAlarmPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAlarmMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAlarmPropsMixin
type jsiiProxy_CfnAlarmPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAlarmPropsMixin) Props() *CfnAlarmMixinProps {
	var returns *CfnAlarmMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lightsail::Alarm`.
func NewCfnAlarmPropsMixin(props *CfnAlarmMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAlarmPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAlarmPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAlarmPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnAlarmPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lightsail::Alarm`.
func NewCfnAlarmPropsMixin_Override(c CfnAlarmPropsMixin, props *CfnAlarmMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnAlarmPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAlarmPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAlarmPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnAlarmPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlarmPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lightsail.mixins.CfnAlarmPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlarmPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAlarmPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

