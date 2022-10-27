//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func validateNewQueryStringParameters(props *QueryStringProps) error {
	return nil
}

