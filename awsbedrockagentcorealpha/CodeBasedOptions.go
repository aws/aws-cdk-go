package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Options for configuring a code-based custom evaluator using a Lambda function.
//
// Uses a Lambda function to implement custom evaluation logic.
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
type CodeBasedOptions struct {
	// The Lambda function used for evaluation.
	//
	// The function will be granted invoke permissions for the
	// `bedrock-agentcore.amazonaws.com` service principal, scoped
	// to this specific evaluator resource.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// The timeout for the Lambda function invocation during evaluation.
	//
	// When not specified, the AgentCore evaluation service uses its default
	// timeout for Lambda-based evaluators.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/custom-evaluators.html
	//
	// Default: - The AgentCore evaluation service's default Lambda timeout is used.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

