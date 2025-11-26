package previewawskinesisanalyticsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawskinesisanalytics/previewawskinesisanalyticsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds an external destination to your Amazon Kinesis Analytics application.
//
// If you want Amazon Kinesis Analytics to deliver data from an in-application stream within your application to an external destination (such as an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an Amazon Lambda function), you add the relevant configuration to your application using this operation. You can configure one or more outputs for your application. Each output configuration maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your in-application error stream to an external destination so that you can analyze the errors. For more information, see [Understanding Application Output (Destination)](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html) .
//
// Any configuration update, including adding a streaming source using this operation, results in a new version of the application. You can use the `DescribeApplication` operation to find the current application version.
//
// For the limits on the number of application inputs and outputs you can configure, see [Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html) .
//
// This operation requires permissions to perform the `kinesisanalytics:AddApplicationOutput` action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationOutputPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationOutputPropsMixin(&CfnApplicationOutputMixinProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	Output: &OutputProperty{
//   		DestinationSchema: &DestinationSchemaProperty{
//   			RecordFormatType: jsii.String("recordFormatType"),
//   		},
//   		KinesisFirehoseOutput: &KinesisFirehoseOutputProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		KinesisStreamsOutput: &KinesisStreamsOutputProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		LambdaOutput: &LambdaOutputProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		Name: jsii.String("name"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html
//
type CfnApplicationOutputPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationOutputMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationOutputPropsMixin
type jsiiProxy_CfnApplicationOutputPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationOutputPropsMixin) Props() *CfnApplicationOutputMixinProps {
	var returns *CfnApplicationOutputMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationOutputPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::KinesisAnalytics::ApplicationOutput`.
func NewCfnApplicationOutputPropsMixin(props *CfnApplicationOutputMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationOutputPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationOutputPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationOutputPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationOutputPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::KinesisAnalytics::ApplicationOutput`.
func NewCfnApplicationOutputPropsMixin_Override(c CfnApplicationOutputPropsMixin, props *CfnApplicationOutputMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationOutputPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationOutputPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationOutputPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationOutputPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationOutputPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_kinesisanalytics.mixins.CfnApplicationOutputPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationOutputPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationOutputPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

