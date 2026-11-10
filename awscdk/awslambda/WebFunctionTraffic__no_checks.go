//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateWebFunctionTraffic_SingleParameters(revisionId *string) error {
	return nil
}

