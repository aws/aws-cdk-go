//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"
)

func validateEvaluatorSelector_BuiltinParameters(evaluator BuiltinEvaluator) error {
	if evaluator == nil {
		return fmt.Errorf("parameter evaluator is required, but nil was provided")
	}

	return nil
}

func validateEvaluatorSelector_CustomParameters(evaluator IEvaluator) error {
	if evaluator == nil {
		return fmt.Errorf("parameter evaluator is required, but nil was provided")
	}

	return nil
}

