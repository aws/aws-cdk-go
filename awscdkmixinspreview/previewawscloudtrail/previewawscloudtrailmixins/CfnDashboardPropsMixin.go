package previewawscloudtrailmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudtrail/previewawscloudtrailmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a custom dashboard or the Highlights dashboard.
//
// - *Custom dashboards* - Custom dashboards allow you to query events in any event data store type. You can add up to 10 widgets to a custom dashboard. You can manually refresh a custom dashboard, or you can set a refresh schedule.
// - *Highlights dashboard* - You can create the Highlights dashboard to see a summary of key user activities and API usage across all your event data stores. CloudTrail Lake manages the Highlights dashboard and refreshes the dashboard every 6 hours. To create the Highlights dashboard, you must set and enable a refresh schedule.
//
// CloudTrail runs queries to populate the dashboard's widgets during a manual or scheduled refresh. CloudTrail must be granted permissions to run the `StartQuery` operation on your behalf. To provide permissions, run the `PutResourcePolicy` operation to attach a resource-based policy to each event data store. For more information, see [Example: Allow CloudTrail to run queries to populate a dashboard](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-eds-dashboard) in the *AWS CloudTrail User Guide* .
//
// To set a refresh schedule, CloudTrail must be granted permissions to run the `StartDashboardRefresh` operation to refresh the dashboard on your behalf. To provide permissions, run the `PutResourcePolicy` operation to attach a resource-based policy to the dashboard. For more information, see [Resource-based policy example for a dashboard](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-dashboards) in the *AWS CloudTrail User Guide* .
//
// For more information about dashboards, see [CloudTrail Lake dashboards](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-dashboard.html) in the *AWS CloudTrail User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDashboardPropsMixin := awscdkmixinspreview.Mixins.NewCfnDashboardPropsMixin(&CfnDashboardMixinProps{
//   	Name: jsii.String("name"),
//   	RefreshSchedule: &RefreshScheduleProperty{
//   		Frequency: &FrequencyProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		Status: jsii.String("status"),
//   		TimeOfDay: jsii.String("timeOfDay"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminationProtectionEnabled: jsii.Boolean(false),
//   	Widgets: []interface{}{
//   		&WidgetProperty{
//   			QueryParameters: []*string{
//   				jsii.String("queryParameters"),
//   			},
//   			QueryStatement: jsii.String("queryStatement"),
//   			ViewProperties: map[string]*string{
//   				"viewPropertiesKey": jsii.String("viewProperties"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-dashboard.html
//
type CfnDashboardPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDashboardMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDashboardPropsMixin
type jsiiProxy_CfnDashboardPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDashboardPropsMixin) Props() *CfnDashboardMixinProps {
	var returns *CfnDashboardMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDashboardPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudTrail::Dashboard`.
func NewCfnDashboardPropsMixin(props *CfnDashboardMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDashboardPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDashboardPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDashboardPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnDashboardPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudTrail::Dashboard`.
func NewCfnDashboardPropsMixin_Override(c CfnDashboardPropsMixin, props *CfnDashboardMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnDashboardPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDashboardPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDashboardPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnDashboardPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDashboardPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnDashboardPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDashboardPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDashboardPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

