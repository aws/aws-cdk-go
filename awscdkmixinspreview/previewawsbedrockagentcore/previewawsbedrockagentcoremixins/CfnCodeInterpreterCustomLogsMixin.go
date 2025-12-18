package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrockagentcore/previewawsbedrockagentcoremixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AgentCore Code Interpreter tool enables agents to securely execute code in isolated sandbox environments.
//
// It offers advanced configuration support and seamless integration with popular frameworks.
//
// For more information about using the custom code interpreter, see [Execute code and analyze data using Amazon Bedrock AgentCore Code Interpreter](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/code-interpreter-tool.html) .
//
// See the *Properties* section below for descriptions of both the required and optional properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnCodeInterpreterCustomLogsMixin := awscdkmixinspreview.Mixins.NewCfnCodeInterpreterCustomLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html
//
type CfnCodeInterpreterCustomLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCodeInterpreterCustomLogsMixin
type jsiiProxy_CfnCodeInterpreterCustomLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCodeInterpreterCustomLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeInterpreterCustomLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::BedrockAgentCore::CodeInterpreterCustom`.
func NewCfnCodeInterpreterCustomLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnCodeInterpreterCustomLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnCodeInterpreterCustomLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCodeInterpreterCustomLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::BedrockAgentCore::CodeInterpreterCustom`.
func NewCfnCodeInterpreterCustomLogsMixin_Override(c CfnCodeInterpreterCustomLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCodeInterpreterCustomLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCodeInterpreterCustomLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCodeInterpreterCustomLogsMixin_APPLICATION_LOGS() CfnCodeInterpreterCustomApplicationLogs {
	_init_.Initialize()
	var returns CfnCodeInterpreterCustomApplicationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomLogsMixin",
		"APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func CfnCodeInterpreterCustomLogsMixin_USAGE_LOGS() CfnCodeInterpreterCustomUsageLogs {
	_init_.Initialize()
	var returns CfnCodeInterpreterCustomUsageLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomLogsMixin",
		"USAGE_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCodeInterpreterCustomLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCodeInterpreterCustomLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

