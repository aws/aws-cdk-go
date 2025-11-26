package previewawscurmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscur/previewawscurmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The definition of AWS Cost and Usage Report.
//
// You can specify the report name, time unit, report format, compression format, S3 bucket, additional artifacts, and schema elements in the definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReportDefinitionPropsMixin := awscdkmixinspreview.Mixins.NewCfnReportDefinitionPropsMixin(&CfnReportDefinitionMixinProps{
//   	AdditionalArtifacts: []*string{
//   		jsii.String("additionalArtifacts"),
//   	},
//   	AdditionalSchemaElements: []*string{
//   		jsii.String("additionalSchemaElements"),
//   	},
//   	BillingViewArn: jsii.String("billingViewArn"),
//   	Compression: jsii.String("compression"),
//   	Format: jsii.String("format"),
//   	RefreshClosedReports: jsii.Boolean(false),
//   	ReportName: jsii.String("reportName"),
//   	ReportVersioning: jsii.String("reportVersioning"),
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Prefix: jsii.String("s3Prefix"),
//   	S3Region: jsii.String("s3Region"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeUnit: jsii.String("timeUnit"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html
//
type CfnReportDefinitionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnReportDefinitionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReportDefinitionPropsMixin
type jsiiProxy_CfnReportDefinitionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnReportDefinitionPropsMixin) Props() *CfnReportDefinitionMixinProps {
	var returns *CfnReportDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportDefinitionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CUR::ReportDefinition`.
func NewCfnReportDefinitionPropsMixin(props *CfnReportDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnReportDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReportDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReportDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cur.mixins.CfnReportDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CUR::ReportDefinition`.
func NewCfnReportDefinitionPropsMixin_Override(c CfnReportDefinitionPropsMixin, props *CfnReportDefinitionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cur.mixins.CfnReportDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnReportDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReportDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cur.mixins.CfnReportDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReportDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cur.mixins.CfnReportDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReportDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnReportDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

