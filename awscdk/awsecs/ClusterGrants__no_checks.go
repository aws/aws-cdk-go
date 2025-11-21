//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterGrants) validateTaskProtectionParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateClusterGrants_FromClusterParameters(resource interfacesawsecs.IClusterRef) error {
	return nil
}

