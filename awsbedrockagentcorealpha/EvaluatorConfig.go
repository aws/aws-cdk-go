package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Configuration for a custom evaluator.
//
// Defines how an evaluator assesses agent performance. Supports two strategies:
// - **LLM-as-a-Judge**: Uses a foundation model with custom instructions and a rating scale.
// - **Code-based**: Uses a Lambda function for custom evaluation logic.
//
// Example:
//   // Code-based evaluator
//   var myEvalFunction IFunction
//   // LLM-as-a-Judge evaluator
//   llmConfig := agentcore.EvaluatorConfig_LlmAsAJudge(&LlmAsAJudgeOptions{
//   	Instructions: jsii.String("Evaluate whether the agent response is helpful."),
//   	ModelId: jsii.String("us.anthropic.claude-sonnet-4-6"),
//   	RatingScale: agentcore.EvaluatorRatingScale_Categorical([]CategoricalRatingOption{
//   		&CategoricalRatingOption{
//   			Label: jsii.String("Good"),
//   			Definition: jsii.String("The response is helpful."),
//   		},
//   		&CategoricalRatingOption{
//   			Label: jsii.String("Bad"),
//   			Definition: jsii.String("The response is not helpful."),
//   		},
//   	}),
//   })
//   codeConfig := agentcore.EvaluatorConfig_CodeBased(&CodeBasedOptions{
//   	LambdaFunction: myEvalFunction,
//   })
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type EvaluatorConfig interface {
	// The Lambda function used for code-based evaluation, if applicable.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LambdaFunction() awslambda.IFunction
}

// The jsii proxy struct for EvaluatorConfig
type jsiiProxy_EvaluatorConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_EvaluatorConfig) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}


// Creates a code-based evaluator configuration using a Lambda function.
//
// The Lambda function implements custom evaluation logic. The function will
// automatically be granted invoke permissions for the bedrock-agentcore service.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func EvaluatorConfig_CodeBased(options *CodeBasedOptions) EvaluatorConfig {
	_init_.Initialize()

	if err := validateEvaluatorConfig_CodeBasedParameters(options); err != nil {
		panic(err)
	}
	var returns EvaluatorConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorConfig",
		"codeBased",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an LLM-as-a-Judge evaluator configuration.
//
// Uses a foundation model to assess agent performance based on custom
// instructions and a rating scale.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func EvaluatorConfig_LlmAsAJudge(options *LlmAsAJudgeOptions) EvaluatorConfig {
	_init_.Initialize()

	if err := validateEvaluatorConfig_LlmAsAJudgeParameters(options); err != nil {
		panic(err)
	}
	var returns EvaluatorConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorConfig",
		"llmAsAJudge",
		[]interface{}{options},
		&returns,
	)

	return returns
}

