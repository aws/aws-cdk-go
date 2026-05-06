//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateEvaluatorReference_BuiltinParameters(evaluator BuiltinEvaluator) error {
	return nil
}

func validateEvaluatorReference_CustomParameters(evaluator IEvaluator) error {
	return nil
}

