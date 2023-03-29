//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateInitUser_FromNameParameters(userName *string, options *InitUserOptions) error {
	return nil
}

func validateNewInitUserParameters(userName *string, userOptions *InitUserOptions) error {
	return nil
}

