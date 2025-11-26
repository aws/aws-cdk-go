package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbedrock/previewawsbedrockmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an agent as a resource in a top-level template. Minimally, you must specify the following properties:.
//
// - AgentName – Specify a name for the agent.
// - AgentResourceRoleArn – Specify the Amazon Resource Name (ARN) of the service role with permissions to invoke API operations on the agent. For more information, see [Create a service role for Agents for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-permissions.html) .
// - FoundationModel – Specify the model ID of a foundation model to use when invoking the agent. For more information, see [Supported regions and models for Agents for Amazon Bedrock](https://docs.aws.amazon.com//bedrock/latest/userguide/agents-supported.html) .
//
// For more information about using agents in Amazon Bedrock , see [Agents for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/agents.html) .
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
//   cfnAgentLogsMixin := awscdkmixinspreview.Mixins.NewCfnAgentLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-agent.html
//
type CfnAgentLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAgentLogsMixin
type jsiiProxy_CfnAgentLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAgentLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAgentLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::Bedrock::Agent`.
func NewCfnAgentLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnAgentLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnAgentLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAgentLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::Bedrock::Agent`.
func NewCfnAgentLogsMixin_Override(c CfnAgentLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAgentLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAgentLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAgentLogsMixin_APPLICATION_LOGS() CfnAgentApplicationLogs {
	_init_.Initialize()
	var returns CfnAgentApplicationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentLogsMixin",
		"APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAgentLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAgentLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

