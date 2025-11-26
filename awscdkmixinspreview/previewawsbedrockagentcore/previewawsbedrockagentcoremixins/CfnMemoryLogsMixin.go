package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrockagentcore/previewawsbedrockagentcoremixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Memory allows AI agents to maintain both immediate and long-term knowledge, enabling context-aware and personalized interactions.
//
// For more information about using Memory in Amazon Bedrock AgentCore, see [Host agent or tools with Amazon Bedrock AgentCore Memory](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/memory-getting-started.html) .
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
//   cfnMemoryLogsMixin := awscdkmixinspreview.Mixins.NewCfnMemoryLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html
//
type CfnMemoryLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMemoryLogsMixin
type jsiiProxy_CfnMemoryLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMemoryLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMemoryLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::BedrockAgentCore::Memory`.
func NewCfnMemoryLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnMemoryLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnMemoryLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMemoryLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::BedrockAgentCore::Memory`.
func NewCfnMemoryLogsMixin_Override(c CfnMemoryLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMemoryLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMemoryLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMemoryLogsMixin_APPLICATION_LOGS() CfnMemoryApplicationLogs {
	_init_.Initialize()
	var returns CfnMemoryApplicationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryLogsMixin",
		"APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func CfnMemoryLogsMixin_TRACES() CfnMemoryTraces {
	_init_.Initialize()
	var returns CfnMemoryTraces
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryLogsMixin",
		"TRACES",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMemoryLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMemoryLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

