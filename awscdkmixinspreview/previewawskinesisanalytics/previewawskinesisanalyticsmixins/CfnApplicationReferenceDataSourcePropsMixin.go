package previewawskinesisanalyticsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawskinesisanalytics/previewawskinesisanalyticsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds a reference data source to an existing application.
//
// Amazon Kinesis Analytics reads reference data (that is, an Amazon S3 object) and creates an in-application table within your application. In the request, you provide the source (S3 bucket name and object key name), name of the in-application table to create, and the necessary mapping information that describes how data in Amazon S3 object maps to columns in the resulting in-application table.
//
// For conceptual information, see [Configuring Application Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html) . For the limits on data sources you can add to your application, see [Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html) .
//
// This operation requires permissions to perform the `kinesisanalytics:AddApplicationOutput` action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationReferenceDataSourcePropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationReferenceDataSourcePropsMixin(&CfnApplicationReferenceDataSourceMixinProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	ReferenceDataSource: &ReferenceDataSourceProperty{
//   		ReferenceSchema: &ReferenceSchemaProperty{
//   			RecordColumns: []interface{}{
//   				&RecordColumnProperty{
//   					Mapping: jsii.String("mapping"),
//   					Name: jsii.String("name"),
//   					SqlType: jsii.String("sqlType"),
//   				},
//   			},
//   			RecordEncoding: jsii.String("recordEncoding"),
//   			RecordFormat: &RecordFormatProperty{
//   				MappingParameters: &MappingParametersProperty{
//   					CsvMappingParameters: &CSVMappingParametersProperty{
//   						RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   						RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   					},
//   					JsonMappingParameters: &JSONMappingParametersProperty{
//   						RecordRowPath: jsii.String("recordRowPath"),
//   					},
//   				},
//   				RecordFormatType: jsii.String("recordFormatType"),
//   			},
//   		},
//   		S3ReferenceDataSource: &S3ReferenceDataSourceProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			FileKey: jsii.String("fileKey"),
//   			ReferenceRoleArn: jsii.String("referenceRoleArn"),
//   		},
//   		TableName: jsii.String("tableName"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationreferencedatasource.html
//
type CfnApplicationReferenceDataSourcePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationReferenceDataSourceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationReferenceDataSourcePropsMixin
type jsiiProxy_CfnApplicationReferenceDataSourcePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourcePropsMixin) Props() *CfnApplicationReferenceDataSourceMixinProps {
	var returns *CfnApplicationReferenceDataSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationReferenceDataSourcePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KinesisAnalytics::ApplicationReferenceDataSource`.
func NewCfnApplicationReferenceDataSourcePropsMixin(props *CfnApplicationReferenceDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationReferenceDataSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationReferenceDataSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationReferenceDataSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationReferenceDataSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KinesisAnalytics::ApplicationReferenceDataSource`.
func NewCfnApplicationReferenceDataSourcePropsMixin_Override(c CfnApplicationReferenceDataSourcePropsMixin, props *CfnApplicationReferenceDataSourceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationReferenceDataSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationReferenceDataSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationReferenceDataSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationReferenceDataSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationReferenceDataSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationReferenceDataSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourcePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnApplicationReferenceDataSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

