package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a reference to an evaluator for online evaluation.
//
// Use the static factory methods to create evaluator references:
// - `EvaluatorReference.builtin()` for built-in evaluators
// - `EvaluatorReference.custom()` for custom evaluators
//
// Example:
//   // Using custom evaluators
//   var myCustomEvaluator IEvaluator
//   // Using built-in evaluators
//   helpfulness := agentcore.EvaluatorReference_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS())
//   custom := agentcore.EvaluatorReference_Custom(myCustomEvaluator)
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type EvaluatorReference interface {
	// The evaluator identifier.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	EvaluatorId() *string
	// Binds the evaluator reference to produce the L1 property.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Bind() *EvaluatorReferenceBindResult
}

// The jsii proxy struct for EvaluatorReference
type jsiiProxy_EvaluatorReference struct {
	_ byte // padding
}

func (j *jsiiProxy_EvaluatorReference) EvaluatorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluatorId",
		&returns,
	)
	return returns
}


// Creates a reference to a built-in evaluator.
//
// Built-in evaluators are provided by Amazon Bedrock AgentCore and assess
// different aspects of agent performance at various levels (session, trace, or tool call).
//
// Returns: An EvaluatorReference instance.
//
// Example:
//   helpfulness := agentcore.EvaluatorReference_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS())
//   goalSuccess := agentcore.EvaluatorReference_Builtin(agentcore.BuiltinEvaluator_GOAL_SUCCESS_RATE())
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func EvaluatorReference_Builtin(evaluator BuiltinEvaluator) EvaluatorReference {
	_init_.Initialize()

	if err := validateEvaluatorReference_BuiltinParameters(evaluator); err != nil {
		panic(err)
	}
	var returns EvaluatorReference

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorReference",
		"builtin",
		[]interface{}{evaluator},
		&returns,
	)

	return returns
}

// Creates a reference to a custom evaluator.
//
// Custom evaluators are created using the `Evaluator` construct and can be
// LLM-as-a-Judge or code-based (Lambda) evaluators.
//
// Returns: An EvaluatorReference instance.
//
// Example:
//   var myCustomEvaluator IEvaluator
//
//   ref := agentcore.EvaluatorReference_Custom(myCustomEvaluator)
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func EvaluatorReference_Custom(evaluator IEvaluator) EvaluatorReference {
	_init_.Initialize()

	if err := validateEvaluatorReference_CustomParameters(evaluator); err != nil {
		panic(err)
	}
	var returns EvaluatorReference

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorReference",
		"custom",
		[]interface{}{evaluator},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvaluatorReference) Bind() *EvaluatorReferenceBindResult {
	var returns *EvaluatorReferenceBindResult

	_jsii_.Invoke(
		e,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

