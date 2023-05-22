//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateInitGroup_FromNameParameters(groupName *string) error {
	return nil
}

func validateNewInitGroupParameters(groupName *string) error {
	return nil
}

