//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func validateNewPrincipalPolicyFragmentParameters(principalJson *map[string]*[]*string) error {
	return nil
}

