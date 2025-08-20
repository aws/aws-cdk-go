//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func validateLambdaValidator_FromFunctionParameters(func_ awslambda.Function) error {
	return nil
}

