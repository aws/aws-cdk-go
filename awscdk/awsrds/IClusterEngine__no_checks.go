//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IClusterEngine) validateBindToClusterParameters(scope constructs.Construct, options *ClusterEngineBindOptions) error {
	return nil
}

