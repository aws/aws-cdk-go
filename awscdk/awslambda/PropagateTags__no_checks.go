//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validatePropagateTags_ExplicitParameters(tags *map[string]*string) error {
	return nil
}

