package awsbedrockagentcore


// Options for configuring an LLM-as-a-Judge custom evaluator.
//
// Uses a foundation model to assess agent performance based on
// custom instructions and a rating scale.
//
// Example:
//   // Create a custom LLM-as-a-Judge evaluator
//   evaluator := agentcore.NewEvaluator(this, jsii.String("MyEvaluator"), &EvaluatorProps{
//   	EvaluatorName: jsii.String("my_custom_evaluator"),
//   	Level: agentcore.EvaluationLevel_SESSION(),
//   	EvaluatorConfig: agentcore.EvaluatorConfig_LlmAsAJudge(&LlmAsAJudgeOptions{
//   		Instructions: jsii.String("Evaluate whether the agent response is helpful and accurate."),
//   		ModelId: jsii.String("us.anthropic.claude-sonnet-4-6"),
//   		RatingScale: agentcore.EvaluatorRatingScale_Categorical([]CategoricalRatingOption{
//   			&CategoricalRatingOption{
//   				Label: jsii.String("Good"),
//   				Definition: jsii.String("The response is helpful and accurate."),
//   			},
//   			&CategoricalRatingOption{
//   				Label: jsii.String("Bad"),
//   				Definition: jsii.String("The response is not helpful or contains errors."),
//   			},
//   		}),
//   	}),
//   })
//
//   // Use the custom evaluator in an online evaluation configuration
//   // Use the custom evaluator in an online evaluation configuration
//   agentcore.NewOnlineEvaluationConfig(this, jsii.String("MyEvaluation"), &OnlineEvaluationConfigProps{
//   	OnlineEvaluationConfigName: jsii.String("my_evaluation"),
//   	Evaluators: []EvaluatorSelector{
//   		agentcore.EvaluatorSelector_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS()),
//   		agentcore.EvaluatorSelector_Custom(evaluator),
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
type LlmAsAJudgeOptions struct {
	// The evaluation instructions that guide the language model in assessing agent performance.
	//
	// These instructions define the evaluation criteria, context, and expected behavior.
	// Instructions must contain placeholders appropriate for the evaluation level
	// (e.g., `{context}`, `{available_tools}` for SESSION level).
	//
	// Note: Evaluators using reference-input placeholders (e.g., `{expected_tool_trajectory}`,
	// `{assertions}`, `{expected_response}`) are only compatible with on-demand evaluation,
	// not online evaluation.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/custom-evaluators.html
	//
	Instructions *string `field:"required" json:"instructions" yaml:"instructions"`
	// The identifier of the Amazon Bedrock model to use for evaluation.
	//
	// Accepts standard model IDs (e.g., `'anthropic.claude-sonnet-4-6'`)
	// and cross-region inference profile IDs with region prefixes
	// (e.g., `'us.anthropic.claude-sonnet-4-6'`, `'eu.anthropic.claude-sonnet-4-6'`).
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// The rating scale that defines how the evaluator should score agent performance.
	RatingScale EvaluatorRatingScale `field:"required" json:"ratingScale" yaml:"ratingScale"`
	// Additional model-specific request fields.
	// Default: - No additional fields.
	//
	AdditionalModelRequestFields *map[string]interface{} `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// Optional inference configuration parameters that control model behavior during evaluation.
	//
	// When not specified, the foundation model uses its own default values for
	// maxTokens, temperature, and topP.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/custom-evaluators.html
	//
	// Default: - The foundation model's default inference parameters are used.
	//
	InferenceConfig *EvaluatorInferenceConfig `field:"optional" json:"inferenceConfig" yaml:"inferenceConfig"`
}

