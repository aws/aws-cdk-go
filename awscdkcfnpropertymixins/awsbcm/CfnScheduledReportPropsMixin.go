package awsbcm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbcm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates and manages a scheduled report for an AWS Billing and Cost Management Dashboards dashboard.
//
// A scheduled report renders a snapshot of the dashboard on a cron schedule, using an IAM role that the service assumes to read the underlying cost and usage data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnScheduledReportPropsMixin := awscdkcfnpropertymixins.Aws_bcm.NewCfnScheduledReportPropsMixin(&CfnScheduledReportMixinProps{
//   	DashboardArn: jsii.String("dashboardArn"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ScheduleConfig: &ScheduleConfigProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   		ScheduleExpressionTimeZone: jsii.String("scheduleExpressionTimeZone"),
//   		SchedulePeriod: &SchedulePeriodProperty{
//   			EndTime: jsii.String("endTime"),
//   			StartTime: jsii.String("startTime"),
//   		},
//   		State: jsii.String("state"),
//   	},
//   	ScheduledReportExecutionRoleArn: jsii.String("scheduledReportExecutionRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WidgetDateRangeOverride: &DateTimeRangeProperty{
//   		EndTime: &DateTimeValueProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   		StartTime: &DateTimeValueProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WidgetIds: []*string{
//   		jsii.String("widgetIds"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcm-scheduledreport.html
//
type CfnScheduledReportPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnScheduledReportMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnScheduledReportPropsMixin
type jsiiProxy_CfnScheduledReportPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnScheduledReportPropsMixin) Props() *CfnScheduledReportMixinProps {
	var returns *CfnScheduledReportMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledReportPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BCM::ScheduledReport`.
func NewCfnScheduledReportPropsMixin(props *CfnScheduledReportMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnScheduledReportPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnScheduledReportPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScheduledReportPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnScheduledReportPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BCM::ScheduledReport`.
func NewCfnScheduledReportPropsMixin_Override(c CfnScheduledReportPropsMixin, props *CfnScheduledReportMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnScheduledReportPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnScheduledReportPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScheduledReportPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnScheduledReportPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScheduledReportPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnScheduledReportPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScheduledReportPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnScheduledReportPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

