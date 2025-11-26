package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a prompt flow that you can use to send an input through various steps to yield an output.
//
// You define a flow by configuring nodes, each of which corresponds to a step of the flow, and creating connections between the nodes to create paths to different outputs. You can define the flow in one of the following ways:
//
// - Define a [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html) in the `Definition` property.
// - Provide the definition in the `DefinitionString` property as a JSON-formatted string matching the [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html) property.
// - Provide an Amazon S3 location in the `DefinitionS3Location` property that matches the [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html) .
//
// If you use the `DefinitionString` or `DefinitionS3Location` property, you can use the `DefinitionSubstitutions` property to define key-value pairs to replace at runtime.
//
// For more information, see [How it works](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-how-it-works.html) and [Create a prompt flow in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-create.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnFlowLogsMixin := awscdkmixinspreview.Mixins.NewCfnFlowLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-flow.html
//
type CfnFlowLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFlowLogsMixin
type jsiiProxy_CfnFlowLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFlowLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::Bedrock::Flow`.
func NewCfnFlowLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnFlowLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnFlowLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlowLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::Bedrock::Flow`.
func NewCfnFlowLogsMixin_Override(c CfnFlowLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFlowLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowLogsMixin_APPLICATION_LOGS() CfnFlowApplicationLogs {
	_init_.Initialize()
	var returns CfnFlowApplicationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowLogsMixin",
		"APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlowLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFlowLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

