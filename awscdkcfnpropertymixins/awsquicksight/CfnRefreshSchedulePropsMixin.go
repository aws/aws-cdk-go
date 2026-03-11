package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a refresh schedule for a dataset in Quick .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRefreshSchedulePropsMixin := awscdkcfnpropertymixins.Aws_quicksight.NewCfnRefreshSchedulePropsMixin(&CfnRefreshScheduleMixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DataSetId: jsii.String("dataSetId"),
//   	Schedule: &RefreshScheduleMapProperty{
//   		RefreshType: jsii.String("refreshType"),
//   		ScheduleFrequency: &ScheduleFrequencyProperty{
//   			Interval: jsii.String("interval"),
//   			RefreshOnDay: &RefreshOnDayProperty{
//   				DayOfMonth: jsii.String("dayOfMonth"),
//   				DayOfWeek: jsii.String("dayOfWeek"),
//   			},
//   			TimeOfTheDay: jsii.String("timeOfTheDay"),
//   			TimeZone: jsii.String("timeZone"),
//   		},
//   		ScheduleId: jsii.String("scheduleId"),
//   		StartAfterDateTime: jsii.String("startAfterDateTime"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-refreshschedule.html
//
type CfnRefreshSchedulePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRefreshScheduleMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRefreshSchedulePropsMixin
type jsiiProxy_CfnRefreshSchedulePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRefreshSchedulePropsMixin) Props() *CfnRefreshScheduleMixinProps {
	var returns *CfnRefreshScheduleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRefreshSchedulePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::RefreshSchedule`.
func NewCfnRefreshSchedulePropsMixin(props *CfnRefreshScheduleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRefreshSchedulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRefreshSchedulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRefreshSchedulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnRefreshSchedulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::RefreshSchedule`.
func NewCfnRefreshSchedulePropsMixin_Override(c CfnRefreshSchedulePropsMixin, props *CfnRefreshScheduleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnRefreshSchedulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRefreshSchedulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRefreshSchedulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnRefreshSchedulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRefreshSchedulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnRefreshSchedulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRefreshSchedulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRefreshSchedulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

