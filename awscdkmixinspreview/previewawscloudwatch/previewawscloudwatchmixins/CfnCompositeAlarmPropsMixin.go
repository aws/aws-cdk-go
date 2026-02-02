package previewawscloudwatchmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudwatch/previewawscloudwatchmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudWatch::CompositeAlarm` type creates or updates a composite alarm.
//
// When you create a composite alarm, you specify a rule expression for the alarm that takes into account the alarm states of other alarms that you have created. The composite alarm goes into ALARM state only if all conditions of the rule are met.
//
// The alarms specified in a composite alarm's rule expression can include metric alarms and other composite alarms.
//
// Using composite alarms can reduce alarm noise. You can create multiple metric alarms, and also create a composite alarm and set up alerts only for the composite alarm. For example, you could create a composite alarm that goes into ALARM state only when more than one of the underlying metric alarms are in ALARM state.
//
// When this operation creates an alarm, the alarm state is immediately set to INSUFFICIENT_DATA. The alarm is then evaluated and its state is set appropriately. Any actions associated with the new state are then executed. For a composite alarm, this initial time after creation is the only time that the alarm can be in INSUFFICIENT_DATA state.
//
// When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCompositeAlarmPropsMixin := awscdkmixinspreview.Mixins.NewCfnCompositeAlarmPropsMixin(&CfnCompositeAlarmMixinProps{
//   	ActionsEnabled: jsii.Boolean(false),
//   	ActionsSuppressor: jsii.String("actionsSuppressor"),
//   	ActionsSuppressorExtensionPeriod: jsii.Number(123),
//   	ActionsSuppressorWaitPeriod: jsii.Number(123),
//   	AlarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	AlarmDescription: jsii.String("alarmDescription"),
//   	AlarmName: jsii.String("alarmName"),
//   	AlarmRule: jsii.String("alarmRule"),
//   	InsufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	OkActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-compositealarm.html
//
type CfnCompositeAlarmPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCompositeAlarmMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCompositeAlarmPropsMixin
type jsiiProxy_CfnCompositeAlarmPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCompositeAlarmPropsMixin) Props() *CfnCompositeAlarmMixinProps {
	var returns *CfnCompositeAlarmMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCompositeAlarmPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::CompositeAlarm`.
func NewCfnCompositeAlarmPropsMixin(props *CfnCompositeAlarmMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCompositeAlarmPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCompositeAlarmPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCompositeAlarmPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnCompositeAlarmPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::CompositeAlarm`.
func NewCfnCompositeAlarmPropsMixin_Override(c CfnCompositeAlarmPropsMixin, props *CfnCompositeAlarmMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnCompositeAlarmPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCompositeAlarmPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCompositeAlarmPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnCompositeAlarmPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCompositeAlarmPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnCompositeAlarmPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCompositeAlarmPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCompositeAlarmPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

