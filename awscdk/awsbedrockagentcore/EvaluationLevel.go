package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The level at which a custom evaluator assesses agent performance.
//
// Determines what granularity of data the evaluator operates on.
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
type EvaluationLevel interface {
	// The string value of the evaluation level.
	Value() *string
}

// The jsii proxy struct for EvaluationLevel
type jsiiProxy_EvaluationLevel struct {
	_ byte // padding
}

func (j *jsiiProxy_EvaluationLevel) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewEvaluationLevel(value *string) EvaluationLevel {
	_init_.Initialize()

	if err := validateNewEvaluationLevelParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_EvaluationLevel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluationLevel",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewEvaluationLevel_Override(e EvaluationLevel, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluationLevel",
		[]interface{}{value},
		e,
	)
}

func EvaluationLevel_SESSION() EvaluationLevel {
	_init_.Initialize()
	var returns EvaluationLevel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluationLevel",
		"SESSION",
		&returns,
	)
	return returns
}

func EvaluationLevel_TOOL_CALL() EvaluationLevel {
	_init_.Initialize()
	var returns EvaluationLevel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluationLevel",
		"TOOL_CALL",
		&returns,
	)
	return returns
}

func EvaluationLevel_TRACE() EvaluationLevel {
	_init_.Initialize()
	var returns EvaluationLevel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.EvaluationLevel",
		"TRACE",
		&returns,
	)
	return returns
}

