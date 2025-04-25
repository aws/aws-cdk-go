//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterInstance) validateBindParameters(scope constructs.Construct, cluster IDatabaseCluster, props *ClusterInstanceBindOptions) error {
	return nil
}

func validateClusterInstance_ProvisionedParameters(id *string, props *ProvisionedClusterInstanceProps) error {
	return nil
}

func validateClusterInstance_ServerlessV2Parameters(id *string, props *ServerlessV2ClusterInstanceProps) error {
	return nil
}

