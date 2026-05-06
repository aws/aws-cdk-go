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
//   var evalFunction IFunction
//
//
//   codeEvaluator := agentcore.NewEvaluator(this, jsii.String("CodeEvaluator"), &EvaluatorProps{
//   	EvaluatorName: jsii.String("custom_code_evaluator"),
//   	Level: agentcore.EvaluationLevel_TOOL_CALL(),
//   	Description: jsii.String("Evaluates tool call accuracy using custom logic"),
//   	EvaluatorConfig: agentcore.EvaluatorConfig_CodeBased(&CodeBasedOptions{
//   		LambdaFunction: evalFunction,
//   		Timeout: cdk.Duration_Seconds(jsii.Number(30)),
//   	}),
//   })
//
// Experimental.
type CodeBasedOptions struct {
	// The Lambda function used for evaluation.
	//
	// The function will be granted invoke permissions for the
	// `bedrock-agentcore.amazonaws.com` service principal, scoped
	// to this specific evaluator resource.
	// Experimental.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// The timeout for the Lambda function invocation during evaluation.
	//
	// When not specified, the AgentCore evaluation service uses its default
	// timeout for Lambda-based evaluators.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/custom-evaluators.html
	//
	// Default: - The AgentCore evaluation service's default Lambda timeout is used.
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

