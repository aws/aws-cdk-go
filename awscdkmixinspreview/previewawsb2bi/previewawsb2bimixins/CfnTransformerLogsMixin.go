package previewawsb2bimixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsb2bi/previewawsb2bimixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a transformer. AWS B2B Data Interchange currently supports two scenarios:.
//
// - *Inbound EDI* : the AWS customer receives an EDI file from their trading partner. AWS B2B Data Interchange converts this EDI file into a JSON or XML file with a service-defined structure. A mapping template provided by the customer, in JSONata or XSLT format, is optionally applied to this file to produce a JSON or XML file with the structure the customer requires.
// - *Outbound EDI* : the AWS customer has a JSON or XML file containing data that they wish to use in an EDI file. A mapping template, provided by the customer (in either JSONata or XSLT format) is applied to this file to generate a JSON or XML file in the service-defined structure. This file is then converted to an EDI file.
//
// > The following fields are provided for backwards compatibility only: `fileFormat` , `mappingTemplate` , `ediType` , and `sampleDocument` .
// >
// > - Use the `mapping` data type in place of `mappingTemplate` and `fileFormat`
// > - Use the `sampleDocuments` data type in place of `sampleDocument`
// > - Use either the `inputConversion` or `outputConversion` in place of `ediType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnTransformerLogsMixin := awscdkmixinspreview.Mixins.NewCfnTransformerLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html
//
type CfnTransformerLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransformerLogsMixin
type jsiiProxy_CfnTransformerLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTransformerLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformerLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::B2BI::Transformer`.
func NewCfnTransformerLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnTransformerLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransformerLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransformerLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::B2BI::Transformer`.
func NewCfnTransformerLogsMixin_Override(c CfnTransformerLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTransformerLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformerLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransformerLogsMixin_B2BI_EXECUTION_LOGS() CfnTransformerB2biExecutionLogs {
	_init_.Initialize()
	var returns CfnTransformerB2biExecutionLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerLogsMixin",
		"B2BI_EXECUTION_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransformerLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransformerLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

