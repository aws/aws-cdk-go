//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IInstanceEngine) validateBindToInstanceParameters(scope constructs.Construct, options *InstanceEngineBindOptions) error {
	return nil
}

