//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateDirective_CustomParameters(statement *string) error {
	return nil
}

