package mixinsawsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsquicksight/mixinsawsquicksight/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a refresh schedule for a dataset in Quick Suite .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRefreshSchedulePropsMixin := awscdkmixinspreview.Mixins.NewCfnRefreshSchedulePropsMixin(&CfnRefreshScheduleMixinProps{
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-refreshschedule.html
//
type CfnRefreshSchedulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRefreshScheduleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRefreshSchedulePropsMixin
type jsiiProxy_CfnRefreshSchedulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnRefreshSchedulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::RefreshSchedule`.
func NewCfnRefreshSchedulePropsMixin(props *CfnRefreshScheduleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRefreshSchedulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRefreshSchedulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRefreshSchedulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnRefreshSchedulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::RefreshSchedule`.
func NewCfnRefreshSchedulePropsMixin_Override(c CfnRefreshSchedulePropsMixin, props *CfnRefreshScheduleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnRefreshSchedulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRefreshSchedulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRefreshSchedulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnRefreshSchedulePropsMixin",
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
		"@aws-cdk/mixins-preview.aws_quicksight.mixins.CfnRefreshSchedulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRefreshSchedulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

