//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsamplify

// Building without runtime type checking enabled, so all the below just return nil

func validateNewCustomRuleParameters(options *CustomRuleOptions) error {
	return nil
}

