package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CloudWatch::LogAlarm.
//
// A LogAlarm evaluates scheduled query results from CloudWatch Logs and triggers actions when thresholds are breached.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLogAlarmPropsMixin := awscdkcfnpropertymixins.Aws_cloudwatch.NewCfnLogAlarmPropsMixin(&CfnLogAlarmMixinProps{
//   	ActionLogLineCount: jsii.Number(123),
//   	ActionLogLineRoleArn: jsii.String("actionLogLineRoleArn"),
//   	ActionsEnabled: jsii.Boolean(false),
//   	AlarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	AlarmDescription: jsii.String("alarmDescription"),
//   	AlarmName: jsii.String("alarmName"),
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	InsufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	OkActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   	QueryResultsToAlarm: jsii.Number(123),
//   	QueryResultsToEvaluate: jsii.Number(123),
//   	ScheduledQueryConfiguration: &ScheduledQueryConfigurationProperty{
//   		AggregationExpression: jsii.String("aggregationExpression"),
//   		LogGroupIdentifiers: []*string{
//   			jsii.String("logGroupIdentifiers"),
//   		},
//   		QueryLanguage: jsii.String("queryLanguage"),
//   		QueryString: jsii.String("queryString"),
//   		ScheduleConfiguration: &ScheduleConfigurationProperty{
//   			EndTimeOffset: jsii.Number(123),
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//   			StartTimeOffset: jsii.Number(123),
//   		},
//   		ScheduledQueryRoleArn: jsii.String("scheduledQueryRoleArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Threshold: jsii.Number(123),
//   	TreatMissingData: jsii.String("treatMissingData"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-logalarm.html
//
type CfnLogAlarmPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLogAlarmMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLogAlarmPropsMixin
type jsiiProxy_CfnLogAlarmPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLogAlarmPropsMixin) Props() *CfnLogAlarmMixinProps {
	var returns *CfnLogAlarmMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogAlarmPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::LogAlarm`.
func NewCfnLogAlarmPropsMixin(props *CfnLogAlarmMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLogAlarmPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLogAlarmPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLogAlarmPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::LogAlarm`.
func NewCfnLogAlarmPropsMixin_Override(c CfnLogAlarmPropsMixin, props *CfnLogAlarmMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLogAlarmPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLogAlarmPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLogAlarmPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnLogAlarmPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLogAlarmPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLogAlarmPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

