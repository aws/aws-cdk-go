package previewawsarczonalshiftmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsarczonalshift/previewawsarczonalshiftmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The zonal autoshift configuration for a resource includes the practice run configuration and the status for running autoshifts, zonal autoshift status.
//
// When a resource has a practice run configuation, ARC starts weekly zonal shifts for the resource, to shift traffic away from an Availability Zone. Weekly practice runs help you to make sure that your application can continue to operate normally with the loss of one Availability Zone.
//
// You can update the zonal autoshift autoshift status to enable or disable zonal autoshift. When zonal autoshift is `ENABLED` , you authorize AWS to shift away resource traffic for an application from an Availability Zone during events, on your behalf, to help reduce time to recovery. Traffic is also shifted away for the required weekly practice runs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnZonalAutoshiftConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnZonalAutoshiftConfigurationPropsMixin(&CfnZonalAutoshiftConfigurationMixinProps{
//   	PracticeRunConfiguration: &PracticeRunConfigurationProperty{
//   		BlockedDates: []*string{
//   			jsii.String("blockedDates"),
//   		},
//   		BlockedWindows: []*string{
//   			jsii.String("blockedWindows"),
//   		},
//   		BlockingAlarms: []interface{}{
//   			&ControlConditionProperty{
//   				AlarmIdentifier: jsii.String("alarmIdentifier"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		OutcomeAlarms: []interface{}{
//   			&ControlConditionProperty{
//   				AlarmIdentifier: jsii.String("alarmIdentifier"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	ZonalAutoshiftStatus: jsii.String("zonalAutoshiftStatus"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arczonalshift-zonalautoshiftconfiguration.html
//
type CfnZonalAutoshiftConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnZonalAutoshiftConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnZonalAutoshiftConfigurationPropsMixin
type jsiiProxy_CfnZonalAutoshiftConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnZonalAutoshiftConfigurationPropsMixin) Props() *CfnZonalAutoshiftConfigurationMixinProps {
	var returns *CfnZonalAutoshiftConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnZonalAutoshiftConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ARCZonalShift::ZonalAutoshiftConfiguration`.
func NewCfnZonalAutoshiftConfigurationPropsMixin(props *CfnZonalAutoshiftConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnZonalAutoshiftConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnZonalAutoshiftConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnZonalAutoshiftConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnZonalAutoshiftConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ARCZonalShift::ZonalAutoshiftConfiguration`.
func NewCfnZonalAutoshiftConfigurationPropsMixin_Override(c CfnZonalAutoshiftConfigurationPropsMixin, props *CfnZonalAutoshiftConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnZonalAutoshiftConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnZonalAutoshiftConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnZonalAutoshiftConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnZonalAutoshiftConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnZonalAutoshiftConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_arczonalshift.mixins.CfnZonalAutoshiftConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnZonalAutoshiftConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnZonalAutoshiftConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

