package previewawsbackupmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbackup/previewawsbackupmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a report plan.
//
// A report plan is a document that contains information about the contents of the report and where AWS Backup will deliver it.
//
// If you call `CreateReportPlan` with a plan that already exists, you receive an `AlreadyExistsException` exception.
//
// For a sample CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/assigning-resources.html#assigning-resources-cfn) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var reportDeliveryChannel interface{}
//   var reportSetting interface{}
//
//   cfnReportPlanPropsMixin := awscdkmixinspreview.Mixins.NewCfnReportPlanPropsMixin(&CfnReportPlanMixinProps{
//   	ReportDeliveryChannel: reportDeliveryChannel,
//   	ReportPlanDescription: jsii.String("reportPlanDescription"),
//   	ReportPlanName: jsii.String("reportPlanName"),
//   	ReportPlanTags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ReportSetting: reportSetting,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-reportplan.html
//
type CfnReportPlanPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnReportPlanMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReportPlanPropsMixin
type jsiiProxy_CfnReportPlanPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnReportPlanPropsMixin) Props() *CfnReportPlanMixinProps {
	var returns *CfnReportPlanMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportPlanPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Backup::ReportPlan`.
func NewCfnReportPlanPropsMixin(props *CfnReportPlanMixinProps, options *mixins.CfnPropertyMixinOptions) CfnReportPlanPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReportPlanPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReportPlanPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnReportPlanPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Backup::ReportPlan`.
func NewCfnReportPlanPropsMixin_Override(c CfnReportPlanPropsMixin, props *CfnReportPlanMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnReportPlanPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnReportPlanPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReportPlanPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnReportPlanPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReportPlanPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_backup.mixins.CfnReportPlanPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReportPlanPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnReportPlanPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

