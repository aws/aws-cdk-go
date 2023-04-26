//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func validateFargateProfile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewFargateProfileParameters(scope constructs.Construct, id *string, props *FargateProfileProps) error {
	return nil
}

