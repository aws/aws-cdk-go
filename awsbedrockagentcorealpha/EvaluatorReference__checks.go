//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"
)

func validateEvaluatorReference_BuiltinParameters(evaluator BuiltinEvaluator) error {
	if evaluator == nil {
		return fmt.Errorf("parameter evaluator is required, but nil was provided")
	}

	return nil
}

func validateEvaluatorReference_CustomParameters(evaluator IEvaluator) error {
	if evaluator == nil {
		return fmt.Errorf("parameter evaluator is required, but nil was provided")
	}

	return nil
}

