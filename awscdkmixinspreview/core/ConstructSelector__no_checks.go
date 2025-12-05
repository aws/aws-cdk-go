//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func validateConstructSelector_ByIdParameters(pattern *string) error {
	return nil
}

func validateConstructSelector_ByPathParameters(pattern *string) error {
	return nil
}

