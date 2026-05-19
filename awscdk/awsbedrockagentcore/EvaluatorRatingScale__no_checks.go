//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validateEvaluatorRatingScale_CategoricalParameters(options *[]*CategoricalRatingOption) error {
	return nil
}

func validateEvaluatorRatingScale_NumericalParameters(options *[]*NumericalRatingOption) error {
	return nil
}

