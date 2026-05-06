//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateEvaluatorConfig_CodeBasedParameters(options *CodeBasedOptions) error {
	return nil
}

func validateEvaluatorConfig_LlmAsAJudgeParameters(options *LlmAsAJudgeOptions) error {
	return nil
}

