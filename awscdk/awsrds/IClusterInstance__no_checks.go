//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IClusterInstance) validateBindParameters(scope constructs.Construct, cluster IDatabaseCluster, options *ClusterInstanceBindOptions) error {
	return nil
}

