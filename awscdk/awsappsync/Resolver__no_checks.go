//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func validateResolver_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewResolverParameters(scope constructs.Construct, id *string, props *ResolverProps) error {
	return nil
}

