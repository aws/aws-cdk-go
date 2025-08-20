//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func validateAccessPolicy_FromAccessPolicyNameParameters(policyName *string, options *AccessPolicyNameOptions) error {
	return nil
}

func validateNewAccessPolicyParameters(props *AccessPolicyProps) error {
	return nil
}

