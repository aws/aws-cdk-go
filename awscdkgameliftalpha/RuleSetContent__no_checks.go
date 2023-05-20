//go:build no_runtime_type_checking

package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuleSetContent) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func validateRuleSetContent_FromInlineParameters(body *string) error {
	return nil
}

func validateRuleSetContent_FromJsonFileParameters(path *string) error {
	return nil
}

func validateNewRuleSetContentParameters(props *RuleSetContentProps) error {
	return nil
}

