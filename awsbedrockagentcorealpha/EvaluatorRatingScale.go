package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a rating scale for custom LLM-as-a-Judge evaluators.
//
// Rating scales define how the evaluator scores agent performance.
// Use either categorical (discrete labels) or numerical (labeled numeric values) scales.
//
// Example:
//   // Categorical rating scale
//   categorical := agentcore.EvaluatorRatingScale_Categorical([]CategoricalRatingOption{
//   	&CategoricalRatingOption{
//   		Label: jsii.String("Good"),
//   		Definition: jsii.String("The response fully addresses the query."),
//   	},
//   	&CategoricalRatingOption{
//   		Label: jsii.String("Bad"),
//   		Definition: jsii.String("The response fails to address the query."),
//   	},
//   })
//
//   // Numerical rating scale
//   numerical := agentcore.EvaluatorRatingScale_Numerical([]NumericalRatingOption{
//   	&NumericalRatingOption{
//   		Label: jsii.String("Poor"),
//   		Definition: jsii.String("Inadequate response."),
//   		Value: jsii.Number(1),
//   	},
//   	&NumericalRatingOption{
//   		Label: jsii.String("Good"),
//   		Definition: jsii.String("Adequate response."),
//   		Value: jsii.Number(3),
//   	},
//   	&NumericalRatingOption{
//   		Label: jsii.String("Excellent"),
//   		Definition: jsii.String("Outstanding response."),
//   		Value: jsii.Number(5),
//   	},
//   })
//
// Experimental.
type EvaluatorRatingScale interface {
}

// The jsii proxy struct for EvaluatorRatingScale
type jsiiProxy_EvaluatorRatingScale struct {
	_ byte // padding
}

// Creates a categorical rating scale.
//
// Categorical scales define discrete labels for scoring, such as "Good" / "Bad"
// or "Pass" / "Fail".
// Experimental.
func EvaluatorRatingScale_Categorical(options *[]*CategoricalRatingOption) EvaluatorRatingScale {
	_init_.Initialize()

	if err := validateEvaluatorRatingScale_CategoricalParameters(options); err != nil {
		panic(err)
	}
	var returns EvaluatorRatingScale

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorRatingScale",
		"categorical",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a numerical rating scale.
//
// Numerical scales define labeled numeric values for scoring, such as
// 1 (Poor) through 5 (Excellent).
// Experimental.
func EvaluatorRatingScale_Numerical(options *[]*NumericalRatingOption) EvaluatorRatingScale {
	_init_.Initialize()

	if err := validateEvaluatorRatingScale_NumericalParameters(options); err != nil {
		panic(err)
	}
	var returns EvaluatorRatingScale

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.EvaluatorRatingScale",
		"numerical",
		[]interface{}{options},
		&returns,
	)

	return returns
}

