//go:build no_runtime_type_checking

package awscloudtrail

// Building without runtime type checking enabled, so all the below just return nil

func validateNewInsightTypeParameters(value *string) error {
	return nil
}

