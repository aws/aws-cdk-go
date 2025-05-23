//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuleTargetInput) validateBindParameters(rule IRule) error {
	return nil
}

func validateRuleTargetInput_FromEventPathParameters(path *string) error {
	return nil
}

func validateRuleTargetInput_FromMultilineTextParameters(text *string) error {
	return nil
}

func validateRuleTargetInput_FromObjectParameters(obj interface{}) error {
	return nil
}

func validateRuleTargetInput_FromTextParameters(text *string) error {
	return nil
}

