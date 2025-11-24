package mixinsawskinesisanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awskinesisanalytics/mixinsawskinesisanalytics/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::KinesisAnalytics::Application` resource creates an Amazon Kinesis Data Analytics application.
//
// For more information, see the [Amazon Kinesis Data Analytics Developer Guide](https://docs.aws.amazon.com//kinesisanalytics/latest/dev/what-is.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationPropsMixin(&CfnApplicationMixinProps{
//   	ApplicationCode: jsii.String("applicationCode"),
//   	ApplicationDescription: jsii.String("applicationDescription"),
//   	ApplicationName: jsii.String("applicationName"),
//   	Inputs: []interface{}{
//   		&InputProperty{
//   			InputParallelism: &InputParallelismProperty{
//   				Count: jsii.Number(123),
//   			},
//   			InputProcessingConfiguration: &InputProcessingConfigurationProperty{
//   				InputLambdaProcessor: &InputLambdaProcessorProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   			},
//   			InputSchema: &InputSchemaProperty{
//   				RecordColumns: []interface{}{
//   					&RecordColumnProperty{
//   						Mapping: jsii.String("mapping"),
//   						Name: jsii.String("name"),
//   						SqlType: jsii.String("sqlType"),
//   					},
//   				},
//   				RecordEncoding: jsii.String("recordEncoding"),
//   				RecordFormat: &RecordFormatProperty{
//   					MappingParameters: &MappingParametersProperty{
//   						CsvMappingParameters: &CSVMappingParametersProperty{
//   							RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   							RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   						},
//   						JsonMappingParameters: &JSONMappingParametersProperty{
//   							RecordRowPath: jsii.String("recordRowPath"),
//   						},
//   					},
//   					RecordFormatType: jsii.String("recordFormatType"),
//   				},
//   			},
//   			KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			KinesisStreamsInput: &KinesisStreamsInputProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			NamePrefix: jsii.String("namePrefix"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-application.html
//
type CfnApplicationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationPropsMixin
type jsiiProxy_CfnApplicationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Props() *CfnApplicationMixinProps {
	var returns *CfnApplicationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KinesisAnalytics::Application`.
func NewCfnApplicationPropsMixin(props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KinesisAnalytics::Application`.
func NewCfnApplicationPropsMixin_Override(c CfnApplicationPropsMixin, props *CfnApplicationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

