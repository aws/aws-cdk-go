package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a reference to an evaluator for online evaluation.
//
// Use the static factory methods to create evaluator references:
// - `EvaluatorSelector.builtin()` for built-in evaluators
// - `EvaluatorSelector.custom()` for custom evaluators
//
// Example:
//   // Using custom evaluators
//   var myCustomEvaluator IEvaluator
//   // Using built-in evaluators
//   helpfulness := agentcore.EvaluatorSelector_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS())
//   custom := agentcore.EvaluatorSelector_Custom(myCustomEvaluator)
//
type EvaluatorSelector interface {
	// The evaluator identifier.
	EvaluatorId() *string
	// Binds the evaluator reference to produce the L1 property.
	Bind() *EvaluatorSelectorBindResult
}

// The jsii proxy struct for EvaluatorSelector
type jsiiProxy_EvaluatorSelector struct {
	_ byte // padding
}

func (j *jsiiProxy_EvaluatorSelector) EvaluatorId() *string {
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
// Returns: An EvaluatorSelector instance.
//
// Example:
//   helpfulness := agentcore.EvaluatorSelector_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS())
//   goalSuccess := agentcore.EvaluatorSelector_Builtin(agentcore.BuiltinEvaluator_GOAL_SUCCESS_RATE())
//
func EvaluatorSelector_Builtin(evaluator BuiltinEvaluator) EvaluatorSelector {
	_init_.Initialize()

	if err := validateEvaluatorSelector_BuiltinParameters(evaluator); err != nil {
		panic(err)
	}
	var returns EvaluatorSelector

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorSelector",
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
// Returns: An EvaluatorSelector instance.
//
// Example:
//   var myCustomEvaluator IEvaluator
//
//   ref := agentcore.EvaluatorSelector_Custom(myCustomEvaluator)
//
func EvaluatorSelector_Custom(evaluator IEvaluator) EvaluatorSelector {
	_init_.Initialize()

	if err := validateEvaluatorSelector_CustomParameters(evaluator); err != nil {
		panic(err)
	}
	var returns EvaluatorSelector

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluatorSelector",
		"custom",
		[]interface{}{evaluator},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvaluatorSelector) Bind() *EvaluatorSelectorBindResult {
	var returns *EvaluatorSelectorBindResult

	_jsii_.Invoke(
		e,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

