//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validateEvaluatorSelector_BuiltinParameters(evaluator BuiltinEvaluator) error {
	return nil
}

func validateEvaluatorSelector_CustomParameters(evaluator IEvaluator) error {
	return nil
}

