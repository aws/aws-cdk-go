package awsbedrockagentcorealpha


// Properties for creating an Evaluator.
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
//   	Evaluators: []EvaluatorReference{
//   		agentcore.EvaluatorReference_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS()),
//   		agentcore.EvaluatorReference_Custom(evaluator),
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
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type EvaluatorProps struct {
	// The configuration that defines how the evaluator assesses agent performance.
	//
	// Use `EvaluatorConfig.llmAsAJudge()` for model-based evaluation or
	// `EvaluatorConfig.codeBased()` for Lambda-based evaluation.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	EvaluatorConfig EvaluatorConfig `field:"required" json:"evaluatorConfig" yaml:"evaluatorConfig"`
	// The name of the evaluator.
	//
	// Must be unique within your account. Valid characters are a-z, A-Z, 0-9, _ (underscore).
	// Must start with a letter and can be up to 48 characters long.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	EvaluatorName *string `field:"required" json:"evaluatorName" yaml:"evaluatorName"`
	// The level at which the evaluator assesses agent performance.
	//
	// Determines what granularity of data the evaluator operates on:
	// tool call, trace (single request-response), or session (full conversation).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Level EvaluationLevel `field:"required" json:"level" yaml:"level"`
	// The description of the evaluator.
	// Default: - No description.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags for the evaluator.
	//
	// A list of key:value pairs of tags to apply to this Evaluator resource.
	// Default: - No tags.
	//
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

