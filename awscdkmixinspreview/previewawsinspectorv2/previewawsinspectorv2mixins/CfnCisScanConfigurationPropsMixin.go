package previewawsinspectorv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsinspectorv2/previewawsinspectorv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The CIS scan configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var oneTime interface{}
//
//   cfnCisScanConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnCisScanConfigurationPropsMixin(&CfnCisScanConfigurationMixinProps{
//   	ScanName: jsii.String("scanName"),
//   	Schedule: &ScheduleProperty{
//   		Daily: &DailyScheduleProperty{
//   			StartTime: &TimeProperty{
//   				TimeOfDay: jsii.String("timeOfDay"),
//   				TimeZone: jsii.String("timeZone"),
//   			},
//   		},
//   		Monthly: &MonthlyScheduleProperty{
//   			Day: jsii.String("day"),
//   			StartTime: &TimeProperty{
//   				TimeOfDay: jsii.String("timeOfDay"),
//   				TimeZone: jsii.String("timeZone"),
//   			},
//   		},
//   		OneTime: oneTime,
//   		Weekly: &WeeklyScheduleProperty{
//   			Days: []*string{
//   				jsii.String("days"),
//   			},
//   			StartTime: &TimeProperty{
//   				TimeOfDay: jsii.String("timeOfDay"),
//   				TimeZone: jsii.String("timeZone"),
//   			},
//   		},
//   	},
//   	SecurityLevel: jsii.String("securityLevel"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Targets: &CisTargetsProperty{
//   		AccountIds: []*string{
//   			jsii.String("accountIds"),
//   		},
//   		TargetResourceTags: map[string][]*string{
//   			"targetResourceTagsKey": []*string{
//   				jsii.String("targetResourceTags"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-cisscanconfiguration.html
//
type CfnCisScanConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCisScanConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCisScanConfigurationPropsMixin
type jsiiProxy_CfnCisScanConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCisScanConfigurationPropsMixin) Props() *CfnCisScanConfigurationMixinProps {
	var returns *CfnCisScanConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCisScanConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::InspectorV2::CisScanConfiguration`.
func NewCfnCisScanConfigurationPropsMixin(props *CfnCisScanConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCisScanConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCisScanConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCisScanConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::InspectorV2::CisScanConfiguration`.
func NewCfnCisScanConfigurationPropsMixin_Override(c CfnCisScanConfigurationPropsMixin, props *CfnCisScanConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCisScanConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCisScanConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCisScanConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_inspectorv2.mixins.CfnCisScanConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCisScanConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCisScanConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

