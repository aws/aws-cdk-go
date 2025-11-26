package previewawsssmcontactsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsssmcontacts/previewawsssmcontactsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a rotation in an on-call schedule.
//
// > *Template example* : We recommend creating all Incident Manager `Contacts` resources using a single AWS CloudFormation template. For a demonstration, see the examples for [AWS::SSMContacts::Contacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRotationPropsMixin := awscdkmixinspreview.Mixins.NewCfnRotationPropsMixin(&CfnRotationMixinProps{
//   	ContactIds: []*string{
//   		jsii.String("contactIds"),
//   	},
//   	Name: jsii.String("name"),
//   	Recurrence: &RecurrenceSettingsProperty{
//   		DailySettings: []*string{
//   			jsii.String("dailySettings"),
//   		},
//   		MonthlySettings: []interface{}{
//   			&MonthlySettingProperty{
//   				DayOfMonth: jsii.Number(123),
//   				HandOffTime: jsii.String("handOffTime"),
//   			},
//   		},
//   		NumberOfOnCalls: jsii.Number(123),
//   		RecurrenceMultiplier: jsii.Number(123),
//   		ShiftCoverages: []interface{}{
//   			&ShiftCoverageProperty{
//   				CoverageTimes: []interface{}{
//   					&CoverageTimeProperty{
//   						EndTime: jsii.String("endTime"),
//   						StartTime: jsii.String("startTime"),
//   					},
//   				},
//   				DayOfWeek: jsii.String("dayOfWeek"),
//   			},
//   		},
//   		WeeklySettings: []interface{}{
//   			&WeeklySettingProperty{
//   				DayOfWeek: jsii.String("dayOfWeek"),
//   				HandOffTime: jsii.String("handOffTime"),
//   			},
//   		},
//   	},
//   	StartTime: jsii.String("startTime"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeZoneId: jsii.String("timeZoneId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-rotation.html
//
type CfnRotationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRotationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRotationPropsMixin
type jsiiProxy_CfnRotationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRotationPropsMixin) Props() *CfnRotationMixinProps {
	var returns *CfnRotationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRotationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSMContacts::Rotation`.
func NewCfnRotationPropsMixin(props *CfnRotationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRotationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRotationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRotationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmcontacts.mixins.CfnRotationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSMContacts::Rotation`.
func NewCfnRotationPropsMixin_Override(c CfnRotationPropsMixin, props *CfnRotationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmcontacts.mixins.CfnRotationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRotationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRotationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssmcontacts.mixins.CfnRotationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRotationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ssmcontacts.mixins.CfnRotationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRotationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRotationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

