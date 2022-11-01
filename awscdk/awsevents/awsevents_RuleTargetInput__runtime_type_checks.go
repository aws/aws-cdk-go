//go:build !no_runtime_type_checking

package awsevents

import (
	"fmt"
)

func (r *jsiiProxy_RuleTargetInput) validateBindParameters(rule IRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateRuleTargetInput_FromEventPathParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateRuleTargetInput_FromMultilineTextParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

func validateRuleTargetInput_FromObjectParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func validateRuleTargetInput_FromTextParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

