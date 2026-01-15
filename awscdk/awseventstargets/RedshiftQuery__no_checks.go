//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedshiftQuery) validateBindParameters(rule interfacesawsevents.IRuleRef) error {
	return nil
}

func validateNewRedshiftQueryParameters(clusterArn *string, props *RedshiftQueryProps) error {
	return nil
}

