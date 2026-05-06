package awsbedrockagentcorealpha


// Inference configuration for a custom LLM-as-a-Judge evaluator.
//
// Controls how the foundation model generates evaluation responses.
//
// Example:
//   // LLM-as-a-Judge with categorical rating scale
//   categoricalEvaluator := agentcore.NewEvaluator(this, jsii.String("CategoricalEvaluator"), &EvaluatorProps{
//   	EvaluatorName: jsii.String("domain_accuracy_evaluator"),
//   	Level: agentcore.EvaluationLevel_SESSION(),
//   	Description: jsii.String("Evaluates domain-specific accuracy of agent responses"),
//   	EvaluatorConfig: agentcore.EvaluatorConfig_LlmAsAJudge(&LlmAsAJudgeOptions{
//   		Instructions: jsii.String("Evaluate whether the agent response is accurate within the healthcare domain."),
//   		ModelId: jsii.String("us.anthropic.claude-sonnet-4-6"),
//   		RatingScale: agentcore.EvaluatorRatingScale_Categorical([]CategoricalRatingOption{
//   			&CategoricalRatingOption{
//   				Label: jsii.String("Accurate"),
//   				Definition: jsii.String("The response contains factually correct healthcare information."),
//   			},
//   			&CategoricalRatingOption{
//   				Label: jsii.String("Inaccurate"),
//   				Definition: jsii.String("The response contains incorrect or misleading healthcare information."),
//   			},
//   		}),
//   	}),
//   })
//
//   // LLM-as-a-Judge with numerical rating scale and inference config
//   numericalEvaluator := agentcore.NewEvaluator(this, jsii.String("NumericalEvaluator"), &EvaluatorProps{
//   	EvaluatorName: jsii.String("response_quality_evaluator"),
//   	Level: agentcore.EvaluationLevel_TRACE(),
//   	EvaluatorConfig: agentcore.EvaluatorConfig_*LlmAsAJudge(&LlmAsAJudgeOptions{
//   		Instructions: jsii.String("Rate the overall quality of the agent response on a scale of 1 to 5."),
//   		ModelId: jsii.String("us.anthropic.claude-sonnet-4-6"),
//   		RatingScale: agentcore.EvaluatorRatingScale_Numerical([]NumericalRatingOption{
//   			&NumericalRatingOption{
//   				Label: jsii.String("Poor"),
//   				Definition: jsii.String("Inadequate response."),
//   				Value: jsii.Number(1),
//   			},
//   			&NumericalRatingOption{
//   				Label: jsii.String("Below Average"),
//   				Definition: jsii.String("Partially addresses the query."),
//   				Value: jsii.Number(2),
//   			},
//   			&NumericalRatingOption{
//   				Label: jsii.String("Average"),
//   				Definition: jsii.String("Adequately addresses the query."),
//   				Value: jsii.Number(3),
//   			},
//   			&NumericalRatingOption{
//   				Label: jsii.String("Good"),
//   				Definition: jsii.String("Well-structured and accurate response."),
//   				Value: jsii.Number(4),
//   			},
//   			&NumericalRatingOption{
//   				Label: jsii.String("Excellent"),
//   				Definition: jsii.String("Outstanding response exceeding expectations."),
//   				Value: jsii.Number(5),
//   			},
//   		}),
//   		InferenceConfig: &EvaluatorInferenceConfig{
//   			MaxTokens: jsii.Number(1024),
//   			Temperature: jsii.Number(0.1),
//   		},
//   	}),
//   })
//
// Experimental.
type EvaluatorInferenceConfig struct {
	// The maximum number of tokens to generate in the model response.
	// Default: - The foundation model's default maximum token limit is used.
	//
	// Experimental.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// The temperature value that controls randomness in the model's responses.
	//
	// Higher values produce more diverse outputs. Range: 0.0 to 1.0.
	// Default: - The foundation model's default temperature is used.
	//
	// Experimental.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// The top-p sampling parameter that controls the diversity of the model's responses.
	//
	// Range: 0.0 to 1.0.
	// Default: - The foundation model's default top-p value is used.
	//
	// Experimental.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

