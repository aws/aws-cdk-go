package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Built-in evaluators provided by Amazon Bedrock AgentCore.
//
// These evaluators assess different aspects of agent performance
// at various levels (session, trace, or tool call).
//
// Example:
//   var customEvaluator Evaluator
//
//
//   evaluation := agentcore.NewOnlineEvaluationConfig(this, jsii.String("MixedEvaluation"), &OnlineEvaluationConfigProps{
//   	OnlineEvaluationConfigName: jsii.String("mixed_evaluation"),
//   	Evaluators: []EvaluatorReference{
//   		agentcore.EvaluatorReference_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS()),
//   		agentcore.EvaluatorReference_*Builtin(agentcore.BuiltinEvaluator_CORRECTNESS()),
//   		agentcore.EvaluatorReference_Custom(customEvaluator),
//   	},
//   	DataSource: agentcore.DataSourceConfig_FromCloudWatchLogs(&CloudWatchLogsDataSourceConfig{
//   		LogGroupNames: []*string{
//   			jsii.String("/aws/bedrock-agentcore/my-agent"),
//   		},
//   		ServiceNames: []*string{
//   			jsii.String("my-agent.default"),
//   		},
//   	}),
//   })
//
// Experimental.
type BuiltinEvaluator interface {
	// The string value of the built-in evaluator.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for BuiltinEvaluator
type jsiiProxy_BuiltinEvaluator struct {
	_ byte // padding
}

func (j *jsiiProxy_BuiltinEvaluator) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewBuiltinEvaluator(value *string) BuiltinEvaluator {
	_init_.Initialize()

	if err := validateNewBuiltinEvaluatorParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_BuiltinEvaluator{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewBuiltinEvaluator_Override(b BuiltinEvaluator, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		[]interface{}{value},
		b,
	)
}

func BuiltinEvaluator_COHERENCE() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"COHERENCE",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_CONCISENESS() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"CONCISENESS",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_CORRECTNESS() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"CORRECTNESS",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_FAITHFULNESS() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"FAITHFULNESS",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_GOAL_SUCCESS_RATE() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"GOAL_SUCCESS_RATE",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_HARMFULNESS() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"HARMFULNESS",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_HELPFULNESS() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"HELPFULNESS",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_INSTRUCTION_FOLLOWING() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"INSTRUCTION_FOLLOWING",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_REFUSAL() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"REFUSAL",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_RESPONSE_RELEVANCE() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"RESPONSE_RELEVANCE",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_STEREOTYPING() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"STEREOTYPING",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_TOOL_PARAMETER_ACCURACY() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"TOOL_PARAMETER_ACCURACY",
		&returns,
	)
	return returns
}

func BuiltinEvaluator_TOOL_SELECTION_ACCURACY() BuiltinEvaluator {
	_init_.Initialize()
	var returns BuiltinEvaluator
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BuiltinEvaluator",
		"TOOL_SELECTION_ACCURACY",
		&returns,
	)
	return returns
}

