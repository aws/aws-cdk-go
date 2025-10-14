//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedInstancesCapacityProvider) validateBindParameters(cluster ICluster) error {
	return nil
}

func validateManagedInstancesCapacityProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewManagedInstancesCapacityProviderParameters(scope constructs.Construct, id *string, props *ManagedInstancesCapacityProviderProps) error {
	return nil
}

